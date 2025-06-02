package commands

import (
	"fmt"

	"github.com/isudin/gozone/internal/repository/sqlc"
	"github.com/isudin/gozone/internal/usecases"
)

func newInitCommand(queries *sqlc.Queries) Command {
	return Command{
		Name: "init",
		Desc: "Wipe and populate database",
		Exec: func(args []string) error {
			err := usecases.InitDatabase(queries)
			if err != nil {
				fmt.Println(err)
				return err
			}

			return nil
		},
	}
}
