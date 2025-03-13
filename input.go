package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome To Pokedox\ncommands for getting around the pokemon map ")
	for _, value := range userCommands {
		fmt.Println(value.name)
		fmt.Println(value.instruction)
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
		}

	}

}
func checkInput(str string) {

}
