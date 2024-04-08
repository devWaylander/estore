-- name: CreateOrder :one
INSERT INTO orders."order" (user_id)
VALUES ($1)
RETURNING id;

-- name: GetOrder :one
SELECT
    o.id,
    o.user_id,
    o.status
FROM
    orders."order" o
WHERE
    id = $1 AND o.deleted_at IS NULL;

-- name: UpdateOrder :exec
UPDATE orders."order"
SET
    status = $2
WHERE
    id = $1;

-- name: CreateItem :exec
INSERT INTO orders."item" (order_id, sku, count)
VALUES ($1, $2, $3);

-- name: GetItemsByOrderID :many
SELECT
    i.id,
    i.order_id,
    i.sku,
    i.count
FROM
    orders."item" i
WHERE
    order_id = $1 AND i.deleted_at IS NULL;