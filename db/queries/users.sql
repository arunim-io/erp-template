-- name: CreateUser :one
INSERT INTO users ( username, password_hash)
VALUES ( $1, $2 )
RETURNING *;

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;
