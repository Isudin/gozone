package commands

func init() {
	Commands["help"] = Command{
		Name: "help",
		Desc: "List all commands",
		Exec: func(args []string) error {
			for _, Command := range Commands {
				Command.ShowHelpInfo()
			}

			return nil
		},
	}
}
