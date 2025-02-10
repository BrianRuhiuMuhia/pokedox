package main

import "strings"

func cleanInput(str string) []string {
	newStr := strings.ToLower(str)
	words := strings.Fields(newStr)
	return words
}
