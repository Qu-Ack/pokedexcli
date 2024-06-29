package clicommands

import (
	"errors"
	"fmt"

	"github.com/Qu-Ack/pokedexcli/pokeapirequest"
)


type clicommand struct {
		name string
		description string
		Callback func(args ...string) error
}

func InitMap() map[string]clicommand{
		commands := make(map[string]clicommand)
		helpCommand := clicommand{
				name: "help",
				description: "help: display help for pokedex",
				Callback: func(args ...string) error {
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
		exploreCommand := clicommand {
				name: "explore",
				description: "explore: takes the city as the argument and returns the pokemons in that area",
				Callback:commandExplore, 
				
		}
		commands["help"] = helpCommand 
		commands["exit"] = exitCommand
		commands["map"] = mapCommand
		commands["mapb"] = mapBCommand
		commands["explore"] = exploreCommand

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

func CommandExit(args ...string) error {
		return nil
}

func commandMap(args ...string) error {
		err := pokeapirequest.PokeLocationGet()		
		if err != nil {
				return errors.New(fmt.Sprintf("%v", err))
		}
		return nil



} 

func commandBMap(args ...string) error {
		err := pokeapirequest.PokePrevLocationGet()
		if err != nil {
				return errors.New(fmt.Sprintf("%v", err))
				
		}
		return nil

}


func commandExplore(args ...string) error {
		if len(args) < 1 {
				return errors.New("Not Enough arguments")
		}
		fmt.Println("exploring pokemon ...")
		cityname := args[0]
		err := pokeapirequest.PokePokemonGet(cityname)
		if err != nil {
				return err
		}
		return nil
}
