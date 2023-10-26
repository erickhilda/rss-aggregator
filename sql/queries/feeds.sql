-- name: CreateFeed :execresult
INSERT INTO
  feeds (created_at, updated_at, name, url, user_id)
VALUES
  (?, ?, ?, ?, ?);

-- name: GetFeed :many
SELECT
  *
FROM
  feeds;

-- name: GetFeedByID :one
SELECT
  *
FROM
  feeds
WHERE
  id = ?;

-- name: GetNextFeedsToFetch :many
SELECT
  *
FROM
  feeds
ORDER BY
  last_fetched_at IS NULL DESC,
  last_fetched_at DESC
LIMIT
  ?;

-- name: MarkFeedAsFetched :execresult
UPDATE feeds
SET
  last_fetched_at = NOW(),
  updated_at = NOW()
WHERE
  id = ?;