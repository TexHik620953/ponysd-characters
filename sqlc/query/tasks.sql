-- name: CreateTask :one
INSERT INTO tasks (
    status, 
    positive, 
    negative, 
    seed,
    character_id,
    user_id
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id;

-- name: GetTask :one
SELECT * FROM tasks
WHERE id = $1 limit 1;

-- name: GetTaskWithStatus :one
SELECT * FROM tasks
WHERE status = $1 limit 1;


-- name: ListTasksWithStatus :many
SELECT * FROM tasks
where status = $1
ORDER BY created_at;

-- name: UpdateTask :one
UPDATE tasks
SET 
    status = COALESCE($2, status),
    positive = COALESCE($3, positive),
    negative = COALESCE($4, negative),
    seed = COALESCE($5, seed)
WHERE id = $1
RETURNING *;