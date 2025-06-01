package main

import "fmt"

type command struct {
	Name string
	Desc string
	Exec func(args []string) error
}

func (c *command) showHelpInfo() {
	fmt.Println(c.Name + ": " + c.Desc)
}

var Commands = map[string]command{
	"init": {
		Name: "init",
		Desc: "Wipe, initialize and populate database",
		Exec: initialize,
	},
	"start": {
		Name: "start",
		Desc: "Start the game",
		Exec: start,
	},
}

func initialize(args []string) error {
	return nil
}

func start(args []string) error {
	return nil
}
