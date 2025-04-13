package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		strPageURL := *pageURL
		if strPageURL != "" {
			url = strPageURL
		}

	}

	var data []byte
	// check the cache
	data, ok := c.cache.Get(url)

	if !ok {
		fmt.Println("fetching fresh data")
		// cache entry does not exist; fetch data
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespShallowLocations{}, fmt.Errorf("error GET %s : %w", url, err)
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespShallowLocations{}, err
		}

		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespShallowLocations{}, err
		}

		c.cache.Add(url, data)
	}

	locationsResp := RespShallowLocations{}

	err := json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
