package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func getInput() {
	fmt.Print("\033[H\033[2J")
	color.New(color.BgRed).Println("Welcome To Pokedox\ncommands for getting around the pokemon map")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println()
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
