package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Qu-Ack/pokedexcli/clicommands"
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
				default:
						fmt.Println("Invalid Command")
				} 
				fmt.Print("pokecli > ")
		}

}

