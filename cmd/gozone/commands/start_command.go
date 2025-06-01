package commands

func init() {
	Commands["start"] = Command{
		Name: "start",
		Desc: "Start the game",
		Exec: start,
	}
}

func start(args []string) error {
	// TODO: Add checking if database is initialized; if not, prompt to run init command
	return nil
}
