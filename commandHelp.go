package main

import "fmt"

func callbackHelp() {
	fmt.Println("Welcome to the pokedox help menu!")
	fmt.Println("Here are the available commands")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
}
