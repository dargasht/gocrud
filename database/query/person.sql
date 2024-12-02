-- name: GetAllPerson :many
SELECT *
FROM person
ORDER BY id DESC;

-- name: GetPersonByID :one
SELECT *
FROM person
WHERE id = $1;

-- name: CreatePerson :one
INSERT INTO person (name, email)
VALUES ($1, $2)
RETURNING *;

-- name: UpdatePerson :execrows
UPDATE person
SET name = $1,
	email = $2
WHERE id = $3;

-- name: DeletePerson :execrows
DELETE FROM person
WHERE id = $1;