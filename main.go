package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Error scanning input:", err)
			}
			break
		}
		input := cleanInput(scanner.Text())

		fmt.Println("Your command was:", input)
	}
}

func cleanInput(test string) string {
	return strings.ToLower(strings.Split(strings.TrimSpace(test), " ")[0])
}
