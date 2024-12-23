-- name: GetAllProduct :many
SELECT *
FROM product
ORDER BY id DESC;

-- name: GetProductByID :one
SELECT *
FROM product
WHERE id = $1;

-- name: CreateProduct :one
INSERT INTO product (name, price)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateProduct :execrows
UPDATE product
SET name = $1,
	price = $2
WHERE id = $3;

-- name: DeleteProduct :execrows
DELETE FROM product
WHERE id = $1;