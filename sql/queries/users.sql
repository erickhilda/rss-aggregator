-- name: CreateUser :execresult
INSERT INTO
  users (id, email, created_at, updated_at, api_key)
VALUES
  (?, ?, ?, ?, ?);

-- name: GetUser :one
SELECT
  id,
  email,
  created_at,
  updated_at
FROM
  users
WHERE
  id = ?;

-- name: GetUserByApiKey :one
SELECT
  id,
  email,
  created_at,
  updated_at,
  api_key
FROM
  users
WHERE
  api_key = ?;

-- name: UpdateUser :execresult
UPDATE users
SET
  email = ?,
  updated_at = ?
WHERE
  id = ?;

-- name: DeleteUser :execresult
DELETE FROM users
WHERE
  id = ?;

-- name: ListUsers :many
SELECT
  id,
  email,
  created_at,
  updated_at
FROM
  users
ORDER BY
  id
LIMIT
  ?;