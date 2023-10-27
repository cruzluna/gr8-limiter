package database

import (
	"context"
	"embed"
	"fmt"
	"net"
	"net/url"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/ory/dockertest/v3"
)

//go:embed schema.sql
var fs embed.FS

func TestSetUpDb(t *testing.T) {
	// tb.Helper()

	dsn := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("username", "password"),
		Path:   "neondb",
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")
	dsn.RawQuery = q.Encode()

	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatalf("Could not connect to Docker: %s", err)
	}

	// docker + pool is configured at this point

	pw, _ := dsn.User.Password()
	env := []string{
		fmt.Sprintf("POSTGRES_USER=%s", dsn.User.Username()),
		fmt.Sprintf("POSTGRES_PASSWORD=%s", pw),
		fmt.Sprintf("POSTGRES_DB=%s", dsn.Path),
	}

	resource, err := pool.Run("postgres", "13-alpine", env)
	if err != nil {
		t.Fatalf("Could not start postgres container: %v", err)
	}
	t.Cleanup(func() {
		err = pool.Purge(resource)
		if err != nil {
			t.Fatalf("Could not purge container: %v", err)
		}
	})
	// ----------------

	_ = resource.Expire(60)

	// dsn.Host = resource.Container.NetworkSettings.IPAddress
	// dsn.Host = fmt.Sprintf("%s:5432", resource.Container.NetworkSettings.IPAddress)
	//
	// // MacOS specific
	// if runtime.GOOS == "darwin" {
	// 	dsn.Host = net.JoinHostPort(resource.GetBoundIP("5432/tcp"), resource.GetPort("5432/tcp"))
	// 	t.Log("runtime is darwin")
	// }

	dsn.Host = net.JoinHostPort(resource.GetBoundIP("5432/tcp"), resource.GetPort("5432/tcp"))
	ctx := context.Background()

	// var conn *pgx.Conn

	// database := &Database{
	// 	DB: conn,
	// }

	// err = Init(ctx, dsn.String())
	var db *pgx.Conn
	// gets around DB reconnect fail
	// https://github.com/lib/pq/issues/835
	for i := 0; i < 20; i++ {
		db, err = pgx.Connect(ctx, dsn.String())
		if err == nil {
			break
		}

		time.Sleep(500 * time.Millisecond)
	}

	if db == nil {
		t.Fatalf("Couldn't connect to database: %s", err)
	}

	// database := Database{
	// 	DB: db,
	// }

	defer db.Close(ctx)

	if err = pool.Retry(func() (err error) {
		return db.Ping(ctx)
	}); err != nil {
		t.Fatalf("Couldn't ping DB: %s", err)
	}

	// migrate table in docker
	m, err := migrate.New(
		"file://migrations",
		dsn.String())
	if err != nil {
		t.Fatal(err)
	}
	if err := m.Up(); err != nil {
		t.Fatal(err)
	}

	params := InsertApiKeyParams{
		Apikey: "5fa877b4-88ba-47ec-ad8d-c62e2f46598f",
		Userid: "dummyUserId",
	}

	// insert api key
	qu := New(db)
	err = qu.InsertApiKey(ctx, params)
	if err != nil {
		t.Fatal("unable to insert: ", err)
	}

	quTest := Database{
		DB: db,
	}

	if !quTest.IsApiKeyInTable(ctx, params.Apikey) {
		t.Fatal("unable to find: ", err)
	}

	err = qu.DeleteByApiKey(ctx, params.Apikey)

	if err != nil {
		t.Fatal("unable to delete: ", err)
	}
}
