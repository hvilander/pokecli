package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationByName(name string) (RespDeepLocations, error) {
	if name == "" {
		return RespDeepLocations{}, nil
	}

	url := fmt.Sprintf("%s/location-area/%s", baseURL, name)

	var data []byte
	// check the cache
	data, ok := c.cache.Get(url)

	if !ok {
		// cache entry does not exist; fetch data
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespDeepLocations{}, fmt.Errorf("error GET %s : %w", url, err)
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespDeepLocations{}, err
		}

		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespDeepLocations{}, err
		}

		c.cache.Add(url, data)
	}

	locationsResp := RespDeepLocations{}

	err := json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespDeepLocations{}, err
	}

	return locationsResp, nil
}
