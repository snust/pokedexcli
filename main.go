package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {

	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show all commands",
			callback:    commandHelp,
		},
	}

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		reader.Scan()

		command := reader.Text()
		switch command {
		case "exit":
			commands["exit"].callback()
		case "help":
			commands["help"].callback()
		default:
			fmt.Println("Unknown command. Type 'help' to see available commands.")
		}

	}
}

func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:\n")

	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandExit() error {
	fmt.Println("\nClosing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cleanInput(test string) string {
	return strings.ToLower(strings.Split(strings.TrimSpace(test), " ")[0])
}
