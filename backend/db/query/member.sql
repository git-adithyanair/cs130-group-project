-- name: CreateMember :one
INSERT INTO members (
  user_id,
  community_id
) VALUES (
  $1,
  $2
) RETURNING *;

-- name: GetMember :one
SELECT * FROM members WHERE
    user_id = $1 AND
    community_id = $2;

-- name: ListMembersByCommunity :many
SELECT * FROM members WHERE community_id = $1;

-- name: ListMembersByUser :many
SELECT * FROM members WHERE user_id = $1;

-- name: DeleteMember :exec
DELETE FROM members WHERE
    user_id = $1 AND
    community_id = $2;

-- name: DeleteMembersByCommunity :exec
DELETE FROM members WHERE
    community_id = $1;

-- name: DeleteMembersByUser :exec
DELETE FROM members WHERE
    user_id = $1;
