package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	DB *pgx.Conn
}

// TABLE api_keys (
//     id SERIAL PRIMARY KEY,
//     api_key text NOT NULL UNIQUE,
//     user_id integer NOT NULL,
//     created_at timestamp DEFAULT current_timestamp,
// )

type ApiTableRecord struct {
	ApiKey string
	UserId int32
}

func StartDatabase(ctx context.Context, dataSource string) (*Database, error) {
	conn, err := pgx.Connect(ctx, dataSource)
	if err != nil {
		fmt.Println("Couldn't start db.\n", err)
		return nil, err
	}

	database := &Database{
		DB: conn,
	}

	return database, nil
}

func (conn Database) Insert(ctx context.Context, record ApiTableRecord) error {
	_, err := conn.DB.Exec(
		ctx,
		"INSERT INTO api_keys (api_key, user_id) VALUES ($1,$2)",
		record.ApiKey,
		record.UserId,
	)
	return err
}
