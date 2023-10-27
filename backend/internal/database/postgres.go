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
