package application

import (
	"context"
	"fmt"

	"ponysd-characters/appconfig"
	"ponysd-characters/services"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/redis/go-redis/v9"
)

type Application struct {
	cfg *appconfig.AppConfig
	ctx context.Context

	pgPool      *pgxpool.Pool
	redisClient *redis.Client
	minioClient *minio.Client

	services services.ServicesContext

	srv *echo.Echo
}

func New(ctx context.Context, cfg *appconfig.AppConfig) (*Application, error) {
	app := &Application{
		cfg: cfg,
		ctx: ctx,
		srv: echo.New(),
	}

	var err error
	// Init database
	app.pgPool, err = pgxpool.New(ctx, cfg.DatabaseDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to init database: %w", err)
	}
	//  Init minio
	app.minioClient, err = minio.New(cfg.MinioAddr, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioKey, cfg.MinioSecret, ""),
		Secure: false,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to minio: %w", err)
	}
	app.services = services.NewServicesContext(app.pgPool, app.minioClient, app.cfg)

	// Configure echo
	app.configureRoutes()

	return app, nil
}

func (app *Application) Run() error {
	return app.srv.Start(app.cfg.HttpAddr)
}

func (app *Application) Stop() {

}
