-- name: CreateFeedFollow :execresult
INSERT INTO
  feed_follows (created_at, updated_at, feed_id, user_id)
VALUES
  (?, ?, ?, ?);

-- name: GetFeedFollowsByUserID :many
SELECT
  *
FROM
  feed_follows
WHERE
  user_id = ?;

-- name: GetFeedFollowByID :one
SELECT
  *
FROM
  feed_follows
WHERE
  id = ?;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE
  id = ?
  AND user_id = ?;