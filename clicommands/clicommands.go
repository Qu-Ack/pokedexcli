package clicommands


import "fmt"


type clicommand struct {
		name string
		description string
		Callback func() error
}

func InitMap() map[string]clicommand{
		commands := make(map[string]clicommand)
		helpCommand := clicommand{
				name: "help",
				description: "help: display help for pokedex",
				Callback: func() error {
						return commandHelp(commands)
				},

		}
		exitCommand := clicommand{
				name: "exit",
				description: "exit: exit the pokedex",
				Callback: CommandExit,
		}
		commands["help"] = helpCommand 
		commands["exit"] = exitCommand


		return commands
}

func commandHelp(commands map[string]clicommand) error {
		fmt.Println("")
		fmt.Println("")
		fmt.Println("Welcome to Pokedex!")
		fmt.Println("Usage: ")
		fmt.Println("")
		fmt.Println("")
		for _, command := range commands {
				fmt.Println(command.description)
		} 
		fmt.Println("")
		fmt.Println("")
		return nil
}

func CommandExit() error {
		return nil
}

