-- name: FindUserByUserId :one
SELECT *
FROM d_users
WHERE user_id = $1
LIMIT 1;
-- name: SaveUser :one
INSERT INTO d_users (user_id, username, created_by, updated_by)
VALUES ($1, $2, $3, $4)
RETURNING id;