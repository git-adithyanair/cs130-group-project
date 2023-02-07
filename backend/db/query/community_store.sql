-- name: CreateCommunityStore :one
INSERT INTO community_stores (
  community_id,
  store_id
) VALUES (
  $1,
  $2
) RETURNING *;

-- name: GetCommunityStore :one
SELECT * FROM community_stores WHERE
    community_id = $1 AND
    store_id = $2;

-- name: ListCommunityStoresByCommunity :many
SELECT * FROM community_stores WHERE community_id = $1;

-- name: ListCommunityStoresByStores :many
SELECT * FROM community_stores WHERE store_id = $1;

-- name: DeleteCommunityStore :exec
DELETE FROM community_stores WHERE
    community_id = $1 AND
    store_id = $2;

-- name: DeleteCommunityStoresByCommunity :exec
DELETE FROM community_stores WHERE
    community_id = $1;

-- name: DeleteCommunityStoresByStore :exec
DELETE FROM community_stores WHERE
    store_id = $1;
