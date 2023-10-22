-- name: CreateUser :execresult
INSERT INTO
  users (id, email, created_at, updated_at)
VALUES
  (?, ?, ?, ?);

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

-- name: GetUserByEmail :one
SELECT
  id,
  email,
  created_at,
  updated_at
FROM
  users
WHERE
  email = ?;

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