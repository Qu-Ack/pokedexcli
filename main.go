package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Qu-Ack/pokedexcli/clicommands"
)



func main() {
		commandMap := clicommands.InitMap()
		scanner:= bufio.NewScanner(os.Stdin)
		fmt.Print("pokecli > ")
		MainLoop:
		for scanner.Scan() {

				text:= scanner.Text()
				parts := strings.SplitN(text, " " , 2)
				command := parts[0]
				argument := ""
				if len(parts) > 1 {
						argument = parts[1]
				}

				switch command {
				case "help":
						commandMap["help"].Callback()
				case "exit":
						break MainLoop
				case "map":
						err :=	commandMap["map"].Callback()
						if err != nil {
								fmt.Println(err)
						}
				case "mapb":
						err := commandMap["mapb"].Callback()
						if err != nil {
								fmt.Println(err)
						}
				case "explore": 
						err := commandMap["explore"].Callback(argument)
						if err != nil {
								fmt.Println(err)
						}
				case "catch": 
						err := commandMap["catch"].Callback(argument)
						if err != nil {
								fmt.Println(err)
						}
				case "inspect":
						err := commandMap["inspect"].Callback(argument)
						if err != nil {
								fmt.Println(err)
						}

				case "pokedex":
						err := commandMap["pokedex"].Callback()
						if err != nil {
								fmt.Println(err)
						}
				default:
						fmt.Println("Invalid Command")
				} 
				fmt.Print("pokecli > ")
		}

}

