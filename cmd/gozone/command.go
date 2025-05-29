package main

import "fmt"

type command struct {
	Name string
	Desc string
	Exec func(args []string) error
}

var Commands = map[string]command{
	"init": {
		Name: "init",
		Desc: "Initialize and populate database",
		Exec: func(args []string) error {
			fmt.Println("Init command executed")
			for _, arg := range args {
				fmt.Println("Arg: " + arg)
			}

			return nil
		},
	},
	"start": {
		Name: "start",
		Desc: "Start the game",
		Exec: func(args []string) error {
			fmt.Println("Started new game")
			return nil
		},
	},
}
