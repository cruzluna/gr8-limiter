package database

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"runtime"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
)

func StartTestDatabase(tb testing.TB) {
	tb.Helper()

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
		tb.Fatalf("Could not connect to Docker: %s", err)
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
		tb.Fatalf("Could not start postgres container: %v", err)
	}
	tb.Cleanup(func() {
		err = pool.Purge(resource)
		if err != nil {
			tb.Fatalf("Could not purge container: %v", err)
		}
	})
	// ----------------

	_ = resource.Expire(60)

	// MacOS specific
	if runtime.GOOS == "darwin" {
		dsn.Host = net.JoinHostPort(resource.GetBoundIP("5432/tcp"), resource.GetPort("5432/tcp"))
	}

	ctx := context.Background()

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
		tb.Fatalf("Couldn't connect to database: %s", err)
	}

	defer db.Close(ctx)

	if err = pool.Retry(func() (err error) {
		return db.Ping(ctx)
	}); err != nil {
		tb.Fatalf("Couldn't ping DB: %s", err)
	}

	// migrate table in docker
	m, err := migrate.New(
		"file://migrations",
		dsn.String())
	if err != nil {
		tb.Fatal(err)
	}
	if err := m.Up(); err != nil {
		tb.Fatal(err)
	}

	err = Init(ctx, dsn.String())

	if err != nil {
		tb.Fatal("Unable to initialize singleton db connection: ", err)
	}
}

func TestDatabase(t *testing.T) {
	StartTestDatabase(t)

	ctx := context.Background()

	testRecord := ApiTableRecord{
		ApiKey: "5fa877b4-88ba-47ec-ad8d-c62e2f46598f",
		UserId: "dummyUserId",
	}

	err := Conn.Insert(ctx, testRecord)
	assert.Nil(t, err, "Insert Api Key")

	assert.True(t, func() bool {
		result, _ := Conn.IsApiKeyInTable(ctx, testRecord.ApiKey)
		return result
	}(), "API key should be found in the table")
	// assert.Equal(t, true, Conn.IsApiKeyInTable(ctx, testRecord.ApiKey), "Should find Api Key")

	err = Conn.DeleteByApiKey(ctx, "5fa877b4-88ba-47ec-ad8d-c62e2f46598f")
	assert.Nil(t, err, "Delete Api Key")

	assert.False(t, func() bool {
		result, _ := Conn.IsApiKeyInTable(ctx, testRecord.ApiKey)
		return result
	}(), "Should NOT find Api Key")

	err = Conn.DeleteByUser(ctx, testRecord.UserId)
	assert.Nil(t, err, "Attempting to delete a deleted record. Delete should always give nil.")

	err = Conn.Insert(ctx, testRecord)
	if err != nil {
		t.Fatal("Unable to insert into database.")
	}

	assert.True(t, func() bool {
		result, _ := Conn.IsApiKeyInTable(ctx, testRecord.ApiKey)
		return result
	}(), "API key should be found in the table")

	err = Conn.DeleteByUserAndApiKey(ctx, testRecord.ApiKey, testRecord.UserId)
	assert.Nil(t, err, "Delete User and Api Key")
}
