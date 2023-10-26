-- name: CreatePost :execresult
INSERT INTO
  posts (created_at, updated_at, feed_id, title, url, description, published_at)
VALUES
  (?, ?, ?, ?, ?, ?, ?);