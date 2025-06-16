package cmd

func AllCommands() []Command {
	return []Command{
		&LsCommand{},
		&ExitCommand{},
		&CdCommand{},
		&CatCommand{},
		&MvCommand{},
		&RmCommand{},
		&RnCommand{},
	}
}
