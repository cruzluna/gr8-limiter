// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package database

import (
	"context"
)

const deleteByApiKey = `-- name: DeleteByApiKey :exec
DELETE FROM api_keys WHERE api_key = $1
`

func (q *Queries) DeleteByApiKey(ctx context.Context, apikey string) error {
	_, err := q.db.Exec(ctx, deleteByApiKey, apikey)
	return err
}

const deleteByUser = `-- name: DeleteByUser :exec
DELETE FROM api_keys WHERE user_id = $1
`

func (q *Queries) DeleteByUser(ctx context.Context, userid string) error {
	_, err := q.db.Exec(ctx, deleteByUser, userid)
	return err
}

const deleteByUserAndApiKey = `-- name: DeleteByUserAndApiKey :exec
DELETE FROM api_keys WHERE user_id = $1
`

func (q *Queries) DeleteByUserAndApiKey(ctx context.Context, userid string) error {
	_, err := q.db.Exec(ctx, deleteByUserAndApiKey, userid)
	return err
}

const getApiKeys = `-- name: GetApiKeys :many
SELECT id, api_key, user_id, created_at FROM api_keys
`

func (q *Queries) GetApiKeys(ctx context.Context) ([]ApiKey, error) {
	rows, err := q.db.Query(ctx, getApiKeys)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ApiKey{}
	for rows.Next() {
		var i ApiKey
		if err := rows.Scan(
			&i.ID,
			&i.ApiKey,
			&i.UserID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertApiKey = `-- name: InsertApiKey :exec
INSERT INTO api_keys (api_key, user_id) VALUES ($1, $2)
`

type InsertApiKeyParams struct {
	Apikey string
	Userid string
}

func (q *Queries) InsertApiKey(ctx context.Context, arg InsertApiKeyParams) error {
	_, err := q.db.Exec(ctx, insertApiKey, arg.Apikey, arg.Userid)
	return err
}