## ponysd-users – User & Auth Service

This service manages **users**, **authentication** and **token balances**.
It exposes a JSON HTTP API built on **Echo v5** and uses PostgreSQL (via `sqlc`) to persist data.

### Features

- **User identities**
  - Username + password (hashed with Argon2)
  - OAuth login via **Google** and **Discord**
  - **Telegram WebApp**-based login
- **Token balance**
  - Each user has a `tokens` balance
  - Every balance change is written to `user_transaction` (positive and negative amounts)
- **JWT auth**
  - Short-lived access token + long-lived refresh token
  - Stateless validation on each request

### Tech stack

- Go 1.22
- Echo v5
- PostgreSQL + `sqlc`
- Redis & MinIO are wired into the shared `ServicesContext` (currently only PostgreSQL is required for core user flows)

### Database schema

Main tables (see `sqlc/migrations/002_database.sql`):

- `user_identity`
  - `id UUID PRIMARY KEY`
  - `username TEXT UNIQUE`
  - `password TEXT NULL` (Argon2 hash)
  - `tokens BIGINT NOT NULL DEFAULT 0`
  - `telegram_id BIGINT NULL`
  - `discord_id TEXT NULL`
  - `google_id TEXT NULL`
  - `created_at TIMESTAMP NOT NULL DEFAULT now()`
- `user_transaction`
  - `id UUID PRIMARY KEY`
  - `user_id UUID REFERENCES user_identity(id)`
  - `amount BIGINT NOT NULL` (can be positive or negative)
  - `meta JSONB NOT NULL DEFAULT '{}'::jsonb`
  - `created_at TIMESTAMP NOT NULL DEFAULT now()`

Migrations are under `sqlc/migrations` and are compatible with `goose`.

### Running locally

#### 1. Start infrastructure (Postgres, Redis, MinIO)

Use the provided `docker-compose.yml`:

```bash
docker compose up -d
```

Configure the following environment variables (for example in `.env` and load them with your shell, or use a process manager that reads them):

- Database:
  - `DATABASE_DSN` (e.g. `postgres://postgres:password@localhost:5432/general`)
- Redis:
  - `REDIS_ADDR` (e.g. `localhost:6379`)
  - `REDIS_PWD`
  - `REDIS_DB`
- MinIO:
  - `MINIO_ADDR` (e.g. `127.0.0.1:9000`)
  - `MINIO_KEY`
  - `MINIO_SECRET`

For the auth flows:

- JWT & sessions:
  - `JWT_SECRET` – secret for signing access/refresh tokens
  - `SESSION_SECRET` – cookie store secret for OAuth sessions
- Google OAuth:
  - `GOOGLE_CLIENT_ID`
  - `GOOGLE_CLIENT_SECRET`
  - `GOOGLE_REDIRECT_URL` (default: `http://localhost:8080/api/auth/google/callback`)
- Discord OAuth:
  - `DISCORD_CLIENT_ID`
  - `DISCORD_CLIENT_SECRET`
  - `DISCORD_REDIRECT_URL` (default: `http://localhost:8080/api/auth/discord/callback`)
- Telegram WebApp:
  - `TG_TOKEN` – Telegram bot token used to validate WebApp `init_data`

#### 2. Run migrations

You can run the migrations using your preferred tool.
The repo is set up for `goose`, so inside a container (see `docker-compose.yml`) it runs:

```bash
goose -dir ./sqlc/migrations up
```

Ensure `user_identity` and `user_transaction` are created.

#### 3. Run the server

From the project root:

```bash
go run ./cmd/server
```

The server listens on `HTTP_ADDR` (default `:8080`).

### HTTP API overview

Base path: `/api`

#### Auth

- **POST** `/api/auth/register`
  - Body: `{ "username": "user", "password": "pass" }`
  - Creates a user and returns `{ "access_token", "refresh_token" }`.

- **POST** `/api/auth/login`
  - Body: `{ "username": "user", "password": "pass" }`
  - Returns a JWT token pair.

- **POST** `/api/auth/refresh`
  - Body: `{ "access_token": "...", "refresh_token": "..." }`
  - Validates the refresh token and returns a new pair.

- **GET** `/api/auth/provider/google`
- **GET** `/api/auth/provider/google/callback`
- **GET** `/api/auth/provider/discord`
- **GET** `/api/auth/provider/discord/callback`
  - Standard OAuth 2.0 flow handled by Goth (`application/gothic.go` and `infra/controllers/auth`).
  - On success, the service resolves/creates a `user_identity` record and returns a JWT token pair.

#### Users (CRUD, token-protected)

All `/api/users` endpoints require a valid **Bearer** access token:

Header:

```http
Authorization: Bearer <access_token>
```

- **GET** `/api/users`
  - Lists users.
- **GET** `/api/users/:id`
  - Returns a single user.
- **POST** `/api/users`
  - Body: `{ "username": "user", "password": "pass" }`
  - Admin-style creation; hashes password and stores the user.
- **PUT** `/api/users/:id`
  - Body (any subset):
    - `username`, `password`, `telegram_id`, `discord_id`, `google_id`
  - Updates mutable fields.
- **DELETE** `/api/users/:id`
  - Deletes the user.

- **POST** `/api/users/:id/tokens`
  - Body:
    ```json
    {
      "amount": 10,
      "meta": {
        "reason": "bonus"
      }
    }
    ```
  - Changes the user’s `tokens` by `amount` (can be negative).
  - Writes a row into `user_transaction` with the same amount and `meta`.

#### Telegram WebApp auth

Use the `TelegramWebAppAuth` middleware from `infra/middleware` when you want to protect routes with Telegram WebApp `init_data`:

- Middleware: `middleware.TelegramWebAppAuth(cfg, servicesCtx)`
- Expects `init_data` query parameter with Telegram WebApp init data.
- Validates it using `TG_TOKEN` and the official algorithm (`init-data-golang`).
- Ensures a `user_identity` row via `EnsureTelegramUser` and exposes `user_id`/`username` in the Echo context.

### Where things live

- `cmd/server/main.go` – entrypoint, wires config and `Application`.
- `application/` – Echo setup, routing and service wiring.
- `infra/controllers/auth` – registration, login, refresh, Google/Discord OAuth.
- `infra/controllers/users` – user CRUD and token balance change endpoints.
- `infra/middleware` – token auth and Telegram WebApp auth, validation-aware error handler.
- `services/auth` – password hashing, JWT access + refresh token pair generation and validation.
- `services/user` – user CRUD, identity lookups (username, Google, Discord, Telegram), token balance and transactions.
- `sqlc/query` + `sqlc/migrations` – SQL schema and queries used by `sqlc` to generate `internal/repos`.

### Notes

- Every token balance change is applied via `services/user.ChangeTokens`, which:
  - Updates `user_identity.tokens`.
  - Inserts a corresponding `user_transaction` record.
- Unused legacy pieces (e.g. GitHub identity) have been removed from the schema and services to keep this service focused on users, auth and tokens.

