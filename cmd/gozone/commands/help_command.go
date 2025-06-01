package commands

func newHelpCommand(commands map[string]Command) Command {
	return Command{
		Name: "help",
		Desc: "List all commands",
		Exec: func(args []string) error {
			for _, Command := range commands {
				Command.ShowHelpInfo()
			}

			return nil
		},
	}
}
