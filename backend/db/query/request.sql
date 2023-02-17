-- name: CreateRequest :one
INSERT INTO requests (
    user_id, 
    community_id, 
    store_id
) VALUES (
    $1, $2, $3
) RETURNING *; 

-- name: GetRequest :one
SELECT * FROM requests WHERE id = $1; 

-- name: GetRequestsByUserId :many
SELECT * FROM requests 
WHERE user_id = $1
LIMIT $2
OFFSET $3; 

-- name: GetRequestsByStoreId :many
SELECT * FROM requests WHERE store_id = $1; 

-- name: GetRequestsByCommunityId :many
SELECT * FROM requests 
WHERE community_id = $1
LIMIT $2
OFFSET $3; 

-- name: GetPendingRequestsByCommunityId :many
SELECT * FROM requests 
WHERE community_id = $1 AND status = 'pending'; 

-- name: GetPendingRequestsByStoreId :many
SELECT * FROM requests 
WHERE store_id = $1 AND status = 'pending'; 

-- name: GetRequestsByErrandId :many
SELECT * FROM requests WHERE errand_id = $1; 

-- name: GetRequestsForUserByStatus :many
SELECT * FROM requests
WHERE user_id = $1 and status = $2; 

-- name: ListRequests :many
SELECT * FROM requests
LIMIT $1
OFFSET $2; 

-- name: UpdateRequest :one
UPDATE requests SET 
    user_id = $2, 
    community_id = $3, 
    status = $4, 
    errand_id = $5, 
    store_id = $6
WHERE id = $1
RETURNING *; 

-- name: UpdateRequestStatus :one
UPDATE requests SET status = $2 
WHERE id = $1
RETURNING *; 

-- name: UpdateRequestErrandAndStatus :exec
UPDATE requests SET 
    errand_id = $2,
    status = $3
WHERE id = $1; 

-- name: DeleteRequest :exec
DELETE FROM requests WHERE id = $1; 

-- name: DeleteRequestsByStore :exec
DELETE FROM requests WHERE store_id = $1; 

-- name: DeleteRequestsByUser :exec
DELETE FROM requests WHERE user_id = $1; 

-- name: DeleteRequestsByErrand :exec
DELETE FROM requests WHERE errand_id = $1; 
