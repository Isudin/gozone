package usecases

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/isudin/gozone/internal/repository/sqlc"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Join(filepath.Dir(b), "..", "..")

	envPath := filepath.Join(basepath, ".env.test")
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("failed to load env file: %v", err)
	}
}

var db *sql.DB

func TestMain(m *testing.M) {
	var err error
	db, err = sql.Open("postgres", os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}

	err = goose.Up(db, os.Getenv("GOOSE_MIGRATION_DIR"))
	if err != nil {
		log.Fatalf("could not migrate the database: %v", err)
	}
	code := m.Run()
	db.Close()
	os.Exit(code)
}

func withTx(t *testing.T, fn func(*sqlc.Queries)) {
	tx, err := db.Begin()
	require.NoError(t, err)
	defer tx.Rollback()

	q := sqlc.New(tx)
	fn(q)
}

func TestInitFactions(t *testing.T) {
	withTx(t, func(q *sqlc.Queries) {
		err := InitDatabase(q)
		require.NoError(t, err)

		factions, err := q.GetAllFactions(context.Background())
		require.NoError(t, err)
		require.Equal(t, len(factionNames), len(factions))

		npcCount, err := q.CountNpc(context.Background())
		require.NoError(t, err)
		require.Equal(t, true, npcCount > 0)
	})
}
