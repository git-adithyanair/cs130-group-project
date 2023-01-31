-- name: CreateStore :one
INSERT INTO stores (
    name, 
    address_line_1, 
    address_line_2, 
    zip_code, 
    city, 
    state, 
    x_coord, 
    y_coord, 
    place_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: GetStore :one
SELECT * FROM stores WHERE id = $1; 

-- name: GetStoreByPlaceId :one
SELECT * FROM stores WHERE place_id = $1; 

-- name: ListStores :many
SELECT * FROM stores
LIMIT $1
OFFSET $2; 

-- name: UpdateStore :one
UPDATE stores SET
    name = $2, 
    address_line_1 = $3, 
    address_line_2 = $4,
    zip_code = $5, 
    city = $6, 
    state = $7, 
    x_coord = $8, 
    y_coord = $9, 
    place_id = $10
WHERE id = $1
RETURNING *;

-- name: DeleteStore :exec
DELETE FROM stores WHERE id = $1; 