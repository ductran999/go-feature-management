-- name: GetByID :one
SELECT * FROM todos WHERE id = $1 LIMIT 1;

-- name: List :many
SELECT * FROM todos ORDER BY created_at;
