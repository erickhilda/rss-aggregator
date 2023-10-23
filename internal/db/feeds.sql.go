// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: feeds.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createFeed = `-- name: CreateFeed :execresult
INSERT INTO
  feeds (created_at, updated_at, name, url, user_id)
VALUES
  (?, ?, ?, ?, ?)
`

type CreateFeedParams struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
	UserID    string
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createFeed,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Url,
		arg.UserID,
	)
}

const getFeed = `-- name: GetFeed :many
SELECT
  id, created_at, updated_at, name, url, user_id
FROM
  feeds
`

func (q *Queries) GetFeed(ctx context.Context) ([]Feed, error) {
	rows, err := q.db.QueryContext(ctx, getFeed)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Url,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFeedByID = `-- name: GetFeedByID :one
SELECT
  id, created_at, updated_at, name, url, user_id
FROM
  feeds
WHERE
  id = ?
`

func (q *Queries) GetFeedByID(ctx context.Context, id uint64) (Feed, error) {
	row := q.db.QueryRowContext(ctx, getFeedByID, id)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
	)
	return i, err
}
