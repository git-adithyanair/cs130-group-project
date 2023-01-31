-- name: CreateUser :one
INSERT INTO users (
    email,
    hashed_password,
    full_name,
    phone_number,
    address_line_1,
    address_line_2,
    zip_code,
    city,
    state
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
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
    address_line_1 = $6,
    address_line_2 = $7,
    zip_code = $8,
    city = $9,
    state = $10
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
