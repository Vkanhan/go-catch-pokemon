package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)

		fmt.Println("echoing: ", cleaned)
	}

}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
