-- name: CreatePost :execresult
INSERT INTO
  posts (
    created_at,
    updated_at,
    feed_id,
    title,
    url,
    description,
    published_at
  )
VALUES
  (?, ?, ?, ?, ?, ?, ?);

-- name: GetPostsForUser :many
SELECT
  posts.*
FROM
  posts
  JOIN feed_follows ON feed_follows.feed_id = posts.feed_id
WHERE
  feed_follows.user_id = ?
ORDER BY
  posts.published_at DESC
LIMIT
  ?;