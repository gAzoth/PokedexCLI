package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2"

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, nil
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, nil
	}
	locationResp := RespShallowLocations{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespShallowLocations{}, nil
	}

	return locationResp, nil
}
