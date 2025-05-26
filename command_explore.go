package main

import "fmt"

func commandExplore(cfg *config, params ...string) error {
	if len(params) != 1 {
		fmt.Println("Invalid area")
		return nil
	}

	area := params[0]

	if area == "" {
		fmt.Println("No area to explore!")
		return nil
	}

	pokemonsResp, err := cfg.pokeapiClient.ListPokemons(area)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s ...", area)
	fmt.Println()
	fmt.Println("Found Pokemon:")

	for _, encounter := range pokemonsResp.PokemonEncounters {
		fmt.Printf("- %s", encounter.Pokemon.Name)
		fmt.Println()
	}

	return nil
}
