-- name: CreateItem :one
INSERT INTO items (
    name,
    requested_by,
    request_id,
    quantity_type,
    quantity,
    preferred_brand,
    image,
    extra_notes
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetItem :one
SELECT * FROM items WHERE id = $1;

-- name: GetItemsByUser :many
SELECT * FROM items WHERE requested_by = $1;

-- name: GetItemsByRequest :many
SELECT * FROM items WHERE request_id = $1;

-- name: ListItems :many
SELECT * FROM items
LIMIT $1
OFFSET $2;

-- name: UpdateItem :one
UPDATE items SET
    name = $2,
    requested_by = $3,
    request_id = $4,
    quantity_type = $5,
    quantity = $6,
    preferred_brand = $7,
    image = $8,
    extra_notes = $9
WHERE id = $1
RETURNING *;

-- name: UpdateItemName :one
UPDATE items SET
    name = $2
WHERE id = $1
RETURNING *;

-- name: UpdateItemQuantity :one
UPDATE items SET
    quantity = $2,
    quantity_type = $3
WHERE id = $1
RETURNING *;

-- name: UpdateItemPreferredBrand :one
UPDATE items SET
    preferred_brand = $2
WHERE id = $1
RETURNING *;

-- name: UpdateItemImage :one
UPDATE items SET
    image = $2
WHERE id = $1
RETURNING *;

-- name: UpdateItemExtraNotes :one
UPDATE items SET
    extra_notes = $2
WHERE id = $1
RETURNING *;

-- name: UpdateItemFound :one
UPDATE items SET
    found = $2
WHERE id = $1
RETURNING *;

-- name: DeleteItem :exec
DELETE FROM items WHERE id = $1;

-- name: DeleteItemsByUser :exec
DELETE FROM items WHERE requested_by = $1;

-- name: DeleteItemsByRequest :exec
DELETE FROM items WHERE request_id = $1;