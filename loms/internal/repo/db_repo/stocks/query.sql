-- name: CreateStock :exec
INSERT INTO stocks."stock" (sku, total_count, reserved)
VALUES ($1, $2, $3);

-- name: GetStockBySKU :one
SELECT 
    s.id, 
    s.sku,
    s.total_count,
    s.reserved,
    s.deleted_at,
    s.created_at
FROM
    stocks."stock" s
WHERE
    s.sku = $1 AND s.deleted_at IS NULL;

-- name: UpdateStock :exec
UPDATE stocks."stock"
SET
    sku = COALESCE($2, sku),
    total_count = COALESCE($3, total_count),
    reserved = COALESCE($4, reserved)
WHERE
    id = $1;