-- db/query/users.sql

-- name: GetUserByPID :one
SELECT *
FROM "user"
WHERE pid = $1;

-- name: GetUserByPhone :one
SELECT *
FROM "user"
WHERE phone = $1;

-- name: ListUsers :many
SELECT *
FROM "user"
LIMIT 5;

-- name: CreateUser :one
INSERT INTO "user" (pid, phone, country_code, password_hash)
VALUES ($1, $2, $3, $4)
RETURNING pid, phone, country_code;

-- name: UpdateUser :exec
UPDATE "user"
SET pid = $2,
    phone = $3,
    updated_at = NOW()
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;