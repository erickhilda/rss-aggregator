-- +goose Up
-- +goose StatementBegin
CREATE TABLE feed_follows (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  user_id VARCHAR(36) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  feed_id INTEGER NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
  UNIQUE (user_id, feed_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE feed_follows;
-- +goose StatementEnd
