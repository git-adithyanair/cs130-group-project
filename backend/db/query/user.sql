-- name: CreateUser :one
INSERT INTO users (
    email,
    hashed_password,
    full_name,
    phone_number,
    place_id,
    address,
    x_coord,
    y_coord
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByPhoneNumber :one
SELECT * FROM users WHERE phone_number = $1;

-- name: ListUsers :many
SELECT * FROM users
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users SET
    email = $2,
    hashed_password = $3,
    full_name = $4,
    phone_number = $5,
    place_id = $6,
    address = $7,
    x_coord = $8,
    y_coord = $9
WHERE id = $1
RETURNING *;

-- name: UpdateUserLocation :exec
UPDATE users SET
    place_id = $2,
    address = $3,
    x_coord = $4,
    y_coord = $5
WHERE id = $1;

-- name: UpdateUserProfilePicture :one
UPDATE users SET profile_picture = $2 
WHERE id = $1
RETURNING *; 

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
