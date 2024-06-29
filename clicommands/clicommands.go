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
		catchCommand := clicommand{
				name:"catch", 
				description: "catch: takes the pokemon as the argument and based on a chance catches it",
				Callback: commandCatch,
		}
		inspectCommand := clicommand{
				name:"inspect",
				description: "inspect: takes the pokemon as the argument and show it's stats",
				Callback: commandInspect,
		}
		pokedexCommand := clicommand{
				name:"pokedex",
				description: "pokedex: displays all the pokemons you have caught yet",
				Callback: commandPokedex,
		}
		
		commands["help"] = helpCommand 
		commands["exit"] = exitCommand
		commands["map"] = mapCommand
		commands["mapb"] = mapBCommand
		commands["explore"] = exploreCommand
		commands["catch"] = catchCommand
		commands["inspect"] = inspectCommand
		commands["pokedex"] = pokedexCommand
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


func commandCatch(args ...string) error {
		if len(args) < 1 {
				return errors.New("Not Enough arguments")
		}
		pokename := args[0]
		fmt.Printf("throwing a pokeball at %v", pokename)
		fmt.Println("")
		err := pokeapirequest.PokeCatch(pokename)
		if err != nil {
				return err
		}
		return nil
}

func commandInspect(args ...string ) error {
		if len(args)  < 1 {
				return errors.New("Not enought arguments")
		}
		pokename := args[0]
		pokeStats := pokeapirequest.GetPokeStats(pokename)	
		fmt.Println(fmt.Sprintf("Name : %v", pokeStats.Name))
		fmt.Println(fmt.Sprintf("Base Exp: %v", pokeStats.BaseExperience))
		fmt.Println(fmt.Sprintf("Height : %v" , pokeStats.Height))
		fmt.Println(fmt.Sprintf("Weight : %v" , pokeStats.Weight))

		return nil
}

func commandPokedex(args ...string) error {
		fmt.Println("Your Pokedex: ")
		for name , _ := range pokeapirequest.Pokedex {
				fmt.Println(fmt.Sprintf(" -%v", name))
		}
		return nil
}
