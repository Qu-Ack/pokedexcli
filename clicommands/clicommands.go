package clicommands

import (
	"errors"
	"fmt"

	"github.com/Qu-Ack/pokedexcli/pokeapirequest"
	"github.com/Qu-Ack/pokedexcli/pokecache"
)


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
		mapCommand := clicommand{
				name: "map",
				description: "map: returns the first 20 locations in the Pokemon world. Each subsequent call will display next 20 locations",
				Callback: commandMap,
				
		}
		mapBCommand := clicommand{
				name: "mapb",
				description: "mapb: returns the previous 20 locations which map returned",
				Callback: commandBMap,
		}
		commands["help"] = helpCommand 
		commands["exit"] = exitCommand
		commands["map"] = mapCommand
		commands["mapb"] = mapBCommand

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
		
		pokecache.ReapLoop(10)
		return nil
}

func CommandExit() error {
		return nil
}

func commandMap() error {
		err := pokeapirequest.PokeLocationGet()		
		if err != nil {
				return errors.New(fmt.Sprintf("%v", err))
		}
		return nil



} 

func commandBMap() error {
		err := pokeapirequest.PokePrevLocationGet()
		if err != nil {
				return errors.New(fmt.Sprintf("%v", err))
				
		}
		return nil
}
