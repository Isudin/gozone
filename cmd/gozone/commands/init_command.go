package commands

import (
	"context"

	"github.com/isudin/gozone/internal/repository/sqlc"
)

func newInitCommand(queries *sqlc.Queries) Command {
	return Command{
		Name: "init",
		Desc: "Wipe, initialize and populate database",
		Exec: func(args []string) error {
			queries.CreateFaction(context.Background(), "Test")
			return nil
		},
	}
}
