package commands

import "fmt"

type Command struct {
	Name string
	Desc string
	Exec func(args []string) error
}

func (c *Command) ShowHelpInfo() {
	fmt.Println(c.Name + ": " + c.Desc)
}

var Commands = map[string]Command{}
