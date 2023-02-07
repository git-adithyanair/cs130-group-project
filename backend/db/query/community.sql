-- name: CreateCommunity :one
INSERT INTO communities (
  name,
  admin,
  center_x_coord,
  center_y_coord,
  range,
  place_id,
  address
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7
) RETURNING *;

-- name: GetCommunity :one
SELECT * FROM communities WHERE id = $1;

-- name: GetCommunitiesByAdmin :many
SELECT * FROM communities WHERE admin = $1;

-- name: ListCommunities :many
SELECT * FROM communities
LIMIT $1
OFFSET $2;

-- name: UpdateCommunity :one
UPDATE communities SET
  name = $2,
  admin = $3,
  center_x_coord = $4,
  center_y_coord = $5,
  range = $6,
  place_id = $7,
  address = $8
WHERE id = $1
RETURNING *;

-- name: UpdateCommunityAdmin :one
UPDATE communities SET
  admin = $2
WHERE id = $1
RETURNING *;

-- name: DeleteCommunity :exec
DELETE FROM communities WHERE id = $1;
