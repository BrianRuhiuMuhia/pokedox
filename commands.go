package main

import (
	"fmt"
	"os"
)

var userCommands = map[string]Command{
	"help": {name: "help", instruction: "The help command shows all the commands and how to use them ",
		callback: helpCmd,
	},
	"map":  {name: "map", instruction: "The map command fetches data from the pokemon server", callback: mapWorld},
	"exit": {name: "exit", instruction: "The exit command exits the program", callback: exitCmd},
}

func helpCmd() {
	fmt.Println("Hello World")
}
func mapWorld() {
	url := "https://pokeapi.co/api/v2/location"
	data, err := fetchData(url)
	if err != nil {
		fmt.Printf("Something Went Wrong")
		return
	}
	for _, value := range data.Results {
		fmt.Println(value.Name)
	}
}
func exitCmd() {
	fmt.Println("Exiting")
	os.Exit(0)
}
