-- name: CreateErrand :one
INSERT INTO errands (
    user_id, 
    community_id
) VALUES (
    $1, $2
) RETURNING *; 

-- name: GetErrand :one
SELECT * FROM errands WHERE id = $1; 

-- name: GetErrandsByUserId :many
SELECT * FROM errands WHERE user_id = $1; 

-- name: GetErrandsByCommunityId :many
SELECT * FROM errands 
WHERE user_id = $1
LIMIT $2
OFFSET $3; 

-- name: ListErrands :many
SELECT * FROM errands 
LIMIT $1
OFFSET $2; 

-- name: UpdateErrand :one
UPDATE errands SET
    user_id = $2, 
    community_id = $3, 
    is_complete = $4
WHERE id = $1
RETURNING *; 

-- name: DeleteErrand :exec
DELETE FROM errands WHERE id = $1; 