package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonByName(name string) (Pokemon, error) {
	if name == "" {
		return Pokemon{}, fmt.Errorf("no name provied for fetch")
	}

	url := fmt.Sprintf("%s/pokemon/%s", baseURL, name)
	var data []byte

	// check the cache
	data, ok := c.cache.Get(url)

	// not found call the API
	if !ok {
		// cache entry does not exist; fetch data
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error GET %s : %w", url, err)
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			if resp.StatusCode == 404 {
				return Pokemon{}, fmt.Errorf("Pokemon not found, check your spelling")
			}

			return Pokemon{}, fmt.Errorf("Issue fetching pokemon")

		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon{}, err
		}

		c.cache.Add(url, data)
	}

	pokeResp := Pokemon{}

	err := json.Unmarshal(data, &pokeResp)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error unmarshaling: %w", err)
	}

	return pokeResp, nil

}
