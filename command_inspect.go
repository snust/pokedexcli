package main

import (
	"fmt"
)

func commandInspect(cfg *config, params ...string) error {
	if len(params) != 1 {
		fmt.Println("Invalid inspect")
		return nil
	}

	name := params[0]

	if name == "" {
		fmt.Println("No pokemon to inspect!")
		return nil
	}

	pokemon, ok := cfg.pokePC[name]

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	} else {
		fmt.Printf("Name: %v \n", pokemon.name)
		fmt.Printf("Height: %v \n", pokemon.height)
		fmt.Printf("Weight: %v \n", pokemon.weight)
		fmt.Println("Stats:")

		for stat, val := range pokemon.stats {
			fmt.Printf("  -%s: %v\n", stat, val)
		}

		fmt.Println("Types:")

		for _, val := range pokemon.types {
			fmt.Printf("  -%s\n", val)
		}
	}

	return nil
}
