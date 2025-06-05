-- name: CreateOrder :one
INSERT INTO "order" (name, amount)
VALUES ($1, $2)
RETURNING id, name, amount;

-- name: GetOrder :one
SELECT id, name, amount
FROM "order"
WHERE id = $1;

-- name: GetAllOrders :many
SELECT id, name, amount
FROM "order"
ORDER BY id;

-- name: UpdateOrder :one
UPDATE "order"
SET name = $2, amount = $3
WHERE id = $1
RETURNING id, name, amount;

-- name: DeleteOrder :exec
DELETE FROM "order"
WHERE id = $1;