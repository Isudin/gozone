package commands

import (
	"fmt"

	"github.com/isudin/gozone/internal/repository/sqlc"
)

type Command struct {
	Name string
	Desc string
	Exec func(args []string) error
}

func (c *Command) ShowHelpInfo() {
	fmt.Println(c.Name + ": " + c.Desc)
}

func InitCommands(queries *sqlc.Queries) map[string]Command {
	cmds := map[string]Command{
		"init":  newInitCommand(queries),
		"start": newStartCommand(),
	}

	cmds["help"] = newHelpCommand(cmds)
	return cmds
}
