-- name: GetAllUser :many
SELECT *
FROM users
ORDER BY id DESC;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (name, email)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateUser :execrows
UPDATE users
SET name = $1,
	email = $2
WHERE id = $3;

-- name: DeleteUser :execrows
DELETE FROM users
WHERE id = $1;