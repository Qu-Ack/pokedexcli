package clicommands

type clicommand struct {
		name string
		description string
		callback func() error
}

func clicommands() {
		commands := make(map[string]clicommand)
		helpCommand := clicommand{
				name: "help",
				description: "display help for pokedex",
				callback: commandHelp,
		}
		commands["help"] = helpCommand 
}

func commandHelp() error {
	return nil	
}

