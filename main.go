package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Qu-Ack/pokedexcli/clicommands"
	"github.com/Qu-Ack/pokedexcli/pokeapirequest"
)



func main() {
		commandMap := clicommands.InitMap()
		scanner:= bufio.NewScanner(os.Stdin)
		fmt.Print("pokecli > ")
		MainLoop:
		for scanner.Scan() {

				text:= scanner.Text()
				switch text {
				case "help":
						commandMap["help"].Callback()
						pokeapirequest.ApiRequest("https://pokeapi.co/api/v2/location/1")	
				case "exit":
						break MainLoop
				default:
						fmt.Println("Invalid Command")
				} 
				fmt.Print("pokecli > ")
		}

}

