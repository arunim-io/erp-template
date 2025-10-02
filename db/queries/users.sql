-- noqa: disable=AM04

-- name: ListUsers :many
SELECT * FROM user_view ORDER BY id LIMIT ? OFFSET ? ;

-- name: GetUserByID :one
SELECT * FROM user_view WHERE id = ? LIMIT 1 ;

-- name: GetUserByUsername :one
SELECT * FROM user_view WHERE username = ? LIMIT 1 ;

-- name: GetUserByEmail :one
SELECT * FROM user_view WHERE email = ? LIMIT 1 ;

-- name: CreateUser :one
INSERT INTO users (
username,
email,
password,
first_name,
last_name
) VALUES (?, ?, ?, ?, ?) RETURNING * ;

-- name: UpdateUser :exec
UPDATE users SET
username = ?,
email = ?,
first_name = ?,
last_name = ?
WHERE id = ? ;

-- name: UpdateUserPassword :exec
UPDATE users SET password = ? WHERE id = ? ;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ? ;
