package usecases

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/isudin/gozone/internal/repository/sqlc"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

var queries *sqlc.Queries

func TestMain(m *testing.M) {
	err := godotenv.Load(".env.test")
	if err != nil {
		log.Fatalf("could not load `.env.test` file: %v", err)
	}

	db, err := sql.Open("postgres", os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}

	err = goose.Up(db, os.Getenv("GOOSE_MIGRATION_DIR"))
	if err != nil {
		log.Fatalf("could not migrate the database: %v", err)
	}
	queries = sqlc.New(db)
	code := m.Run()
	db.Close()
	os.Exit(code)
}
