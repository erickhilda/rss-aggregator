-- name: CreateFeed :execresult
INSERT INTO
  feeds (created_at, updated_at, name, url, user_id)
VALUES
  (?, ?, ?, ?, ?);

-- name: GetFeed :many
SELECT
  *
FROM
  feeds