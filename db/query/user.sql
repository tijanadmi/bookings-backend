-- name: CreateUser :one
INSERT INTO users (
	first_name,
	last_name,
	email,
	phone,
	password,
	access_level
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
  password = COALESCE(sqlc.narg(password), password),
  updated_at = COALESCE(sqlc.narg(updated_at), updated_at),
  first_name = COALESCE(sqlc.narg(first_name), first_name),
  last_name = COALESCE(sqlc.narg(last_name), last_name),
  phone = COALESCE(sqlc.narg(phone), phone),
  access_level = COALESCE(sqlc.narg(access_level), access_level)
WHERE
  email = sqlc.arg(email)
RETURNING *;