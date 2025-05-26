package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(area string) (RespShallowPokemons, error) {
	url := baseURL + "/location-area/" + area

	if val, ok := c.cache.Get(url); ok {
		pokemonsResp := RespShallowPokemons{}
		err := json.Unmarshal(val, &pokemonsResp)
		if err != nil {
			return RespShallowPokemons{}, err
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemons{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemons{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowPokemons{}, err
	}

	pokemonsResp := RespShallowPokemons{}
	err = json.Unmarshal(dat, &pokemonsResp)
	if err != nil {
		return RespShallowPokemons{}, err
	}

	c.cache.Add(url, dat)
	return pokemonsResp, nil
}
