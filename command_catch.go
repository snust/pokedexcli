package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, params ...string) error {
	if len(params) != 1 {
		fmt.Println("Invalid pokemon")
		return nil
	}

	pokemon := params[0]

	if pokemon == "" {
		fmt.Println("No pokemon to catch!")
		return nil
	}

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...", pokemon)
	fmt.Println()

	magicNumber := rand.Intn(pokemonResp.BaseExperience)

	if magicNumber < 100 {
		fmt.Printf("%s was caught!", pokemon)
		fmt.Println()

		caught := &Pokemon{
			name:           pokemonResp.Name,
			baseExperience: pokemonResp.BaseExperience,
			height:         pokemonResp.Height,
			weight:         pokemonResp.Weight,
			stats:          make(map[string]int),
			types:          []string{},
		}

		for _, stat := range pokemonResp.Stats {
			caught.stats[stat.Stat.Name] = stat.BaseStat
		}

		for _, t := range pokemonResp.Types {
			caught.types = append(caught.types, t.Type.Name)
		}

		cfg.pokePC[pokemon] = *caught

	} else {
		fmt.Printf("%s escaped!", pokemon)
		fmt.Println()
	}

	return nil
}
