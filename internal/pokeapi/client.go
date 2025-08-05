package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/errantpianist/pokedexcli/internal/pokecache"
)

type LocationAreasResp struct {
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string, cache *pokecache.Cache) (LocationAreasResp, error) {
	if cachedData, ok := cache.Get(url); ok {
		fmt.Println("Using cached data for", url)
		var locations LocationAreasResp
		err := json.Unmarshal(cachedData, &locations)
		return locations, err
	}

	fmt.Println("Fetching data from", url)

	resp, err := http.Get(url)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	cache.Add(url, body)

	var locations LocationAreasResp
	err = json.Unmarshal(body, &locations)
		if err != nil {
			return LocationAreasResp{}, err
		}


	return locations, nil
}
