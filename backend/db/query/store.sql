-- name: CreateStore :one
INSERT INTO stores (
    name, 
    x_coord, 
    y_coord, 
    place_id,
    address
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetStore :one
SELECT * FROM stores WHERE id = $1; 

-- name: GetStoreByPlaceId :one
SELECT * FROM stores WHERE place_id = $1; 

-- name: ListStores :many
SELECT * FROM stores
LIMIT $1
OFFSET $2; 

-- name: GetStoresByCommunity :many
SELECT stores.* 
FROM stores 
LEFT JOIN community_stores ON community_stores.store_id = stores.id
WHERE community_stores.community_id = $1; 

-- name: UpdateStore :one
UPDATE stores SET
    name = $2, 
    x_coord = $3, 
    y_coord = $4, 
    place_id = $5,
    address = $6
WHERE id = $1
RETURNING *;

-- name: DeleteStore :exec
DELETE FROM stores WHERE id = $1; 