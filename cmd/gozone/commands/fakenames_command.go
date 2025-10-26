package commands

import (
	"flag"
	"fmt"
	"os"

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
	f := flag.NewFlagSet("fakename", flag.ExitOnError)
	number := f.Int("n", 1, "Number of names to generate")
	faction := f.String("faction", "loners", "Declares what faction will the stalker's name belong to")
	f.StringVar(faction, "f", "loners", "Declares what faction will the stalker's name belong to (shorthand)")
	f.Parse(os.Args[2:])

	names := stalkers.GenerateName(*number, *faction)
	for _, name := range names {
		fmt.Println(name)
	}

	return nil
}
