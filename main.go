package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Command struct {
	name        string
	description string
	callback    func() error
}
type Config struct {
}

func commandExit() error {
	os.Exit(0)
	return nil
}
func commandHelp() error {
	fmt.Println("help commands helps with the commands in the pokedox")

	return nil
}
func commandPokemon() error {
	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/location-area/", nil)
	if err != nil {
		log.Fatal("there was an error fetching data")
		return nil
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("there was as error fetching data")
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("there was an error fetching data")
		return nil
	}
	str := string(data)
	fmt.Println(str)
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
		"mapb": {
			name:        "mapb",
			description: "displays the names of the previous 20 location areas in the Pokemon world",
			callback:    nil,
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
