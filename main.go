package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Command struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	os.Exit(0)
	return nil
}
func commandHelp() error {
	fmt.Println("help commands helps with the commands in the pokedox\n")

	return nil
}
func commandPokemon() error {
	resp, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		fmt.Println("There was an error try again")
		os.Exit(0)
	}
	defer resp.Body.Close()
	data, parsing_err := io.ReadAll(resp.Body)
	if parsing_err != nil {
		fmt.Println("There was an error try again")
		os.Exit(0)
	}
	dataStr := string(data)
	fmt.Println(dataStr)

	return nil
}
func main() {
	command_map := map[string]Command{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "contains instructions on commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "displays the names of 20 location areas in the Pokemon world",
			callback:    commandPokemon,
		},
	}
	fmt.Println("Welcome To Pokedox\nUsage:\nhelp: Displays a help message\nexit: Exit the Pokedex\nmap:displays the map of pokemon")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := cleanInput(scanner.Text())
		for key := range command_map {
			if key == input[0] {
				fmt.Println(command_map[key].description)
				command_map[key].callback()
			}
		}

	}

}
