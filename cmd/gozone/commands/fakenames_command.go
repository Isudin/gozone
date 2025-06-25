package commands

import (
	"fmt"

	"github.com/isudin/gozone/internal/usecases/stalkers"
)

func newFakeNameCommand() Command {
	return Command{
		Name: "fakename",
		Desc: "Generates a new stalker name",
		Exec: fakeNames,
	}
}

func fakeNames(args []string) error {
	names := stalkers.GenerateName(2, false)
	for _, name := range names {
		fmt.Println(name)
	}

	return nil
}
