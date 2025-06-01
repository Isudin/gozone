package commands

func init() {
	Commands["init"] = Command{
		Name: "init",
		Desc: "Wipe, initialize and populate database",
		Exec: initialize,
	}
}

func initialize(args []string) error {
	return nil
}
