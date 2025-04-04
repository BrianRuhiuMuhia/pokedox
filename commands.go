package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var userCommands = map[string]Command{
	"help": {name: "help", instruction: "The help command shows all the commands and how to use them\nThe map command fetches data from the pokemon server\nThe exit command exits the program ",
		callback: helpCmd,
	},
	"map":  {name: "map", instruction: "The map command fetches data from the pokemon server", callback: mapWorld},
	"exit": {name: "exit", instruction: "The exit command exits the program\n", callback: exitCmd},
}

func helpCmd() {
	color.Blue("The help command shows all the commands and how to use them\nThe map command fetches data from the pokemon server\nThe exit command exits the program ")
}
func mapWorld() {
	url := "https://pokeapi.co/api/v2/location"
	data, err := fetchData(url)
	if err != nil {
		color.Red("Something Went Wrong: %v\n", err)
		return
	}
	if value, exists := my_cache.cache[url]; exists {
		fmt.Println("getting data from cache")
		for _, obj := range value {
			color.Green("%s\n", obj.Name)
		}
		return
	} else {
		fmt.Println("getting data from server storing in cache")
		for _, value := range data.Results {
			color.Green(value.Name)
		}
		resultChan := make(chan CacheResult)
		my_cache.addData(url, data.Results, resultChan)
		my_cache.deleteData()

	}

}

func exitCmd() {
	fmt.Println("Exiting")
	os.Exit(0)
}
