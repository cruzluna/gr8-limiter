package database

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
)

// singleton db connection
var (
	Conn *Database
	once sync.Once
)

type Database struct {
	DB *pgx.Conn
}

/*
TABLE api_keys (
    id SERIAL PRIMARY KEY,
    api_key uuid NOT NULL,
    user_id text NOT NULL, // clerk user id
    created_at timestamp DEFAULT current_timestamp,
)
*/

type ApiTableRecord struct {
	ApiKey string
	UserId string
}

func Init(ctx context.Context, dataSource string) error {
	var err error
	once.Do(func() {
		Conn, err = StartDatabase(ctx, dataSource)
	})
	return err
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

	err = database.DB.Ping(ctx)
	if err != nil {
		fmt.Println("Couldn't ping db.\n", err)
		return nil, err
	}

	return database, nil
}

func (conn *Database) CloseConn(ctx context.Context) error {
	return conn.DB.Close(ctx)
}

func (conn *Database) Insert(ctx context.Context, record ApiTableRecord) error {
	_, err := conn.DB.Exec(
		ctx,
		"INSERT INTO api_keys (api_key, user_id) VALUES ($1,$2)",
		record.ApiKey,
		record.UserId,
	)
	return err
}

// delete by api key for api_keys table
func (conn *Database) DeleteByApiKey(ctx context.Context, apiKey string) error {
	_, err := conn.DB.Exec(
		ctx,
		"DELETE FROM api_keys WHERE api_key = $1",
		apiKey,
	)
	return err
}

// delete by user id for api_keys table
func (conn *Database) DeleteByUser(ctx context.Context, userId string) error {
	_, err := conn.DB.Exec(
		ctx,
		"DELETE FROM api_keys WHERE user_id = $1",
		userId,
	)
	return err
}

// delete by user id & api key for api_keys table
func (conn *Database) DeleteByUserAndApiKey(
	ctx context.Context,
	apiKey string,
	userId string,
) error {
	_, err := conn.DB.Exec(
		ctx,
		"DELETE FROM api_keys WHERE api_key = $1 AND user_id = $2",
		apiKey,
		userId,
	)
	return err
}

// determine if api key in api_keys table
func (conn *Database) IsApiKeyInTable(ctx context.Context, apiKey string) bool {
	var inTable bool
	conn.DB.QueryRow(
		ctx,
		"SELECT EXISTS(SELECT api_key FROM api_keys WHERE api_key= $1)",
		apiKey,
	).Scan(&inTable)

	return inTable
}
