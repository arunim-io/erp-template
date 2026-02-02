-- name: CreateUser :one
INSERT INTO users ( password_hash, username, email, first_name, last_name, date_joined )
VALUES ( $1, $2, $3, $4, $5, $6 )
RETURNING ( username, email, first_name, last_name );

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;
