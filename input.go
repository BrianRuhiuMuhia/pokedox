package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput() {
	fmt.Print("\033[H\033[2J")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome To Pokedox\ncommands for getting around the pokemon map\n ")
	for _, value := range userCommands {
		fmt.Println(value.name)
	}
	for {
		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd = strings.TrimSpace(cmd)
		if value, ok := userCommands[cmd]; ok {
			value.callback()
		} else {
			fmt.Println("Invalid command")
			userCommands["help"].callback()
		}

	}

}
