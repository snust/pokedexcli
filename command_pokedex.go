package main

import (
	"fmt"
)

func commandPokedex(cfg *config, params ...string) error {

	fmt.Println("Your Pokedex:")
	for key, _ := range cfg.pokePC {
		fmt.Printf("  - %s\n", key)
	}

	return nil
}
