-- name: ListUsers :many
SELECT * FROM users ORDER BY id;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = ? LIMIT 1;

-- name: CreateUser :one
INSERT INTO users ( id ) VALUES ( ? ) RETURNING *;

-- name: UpdateUser :exec
-- UPDATE users SET * where id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;
