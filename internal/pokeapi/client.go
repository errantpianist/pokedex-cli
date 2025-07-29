package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationAreasResp struct {
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (LocationAreasResp, error) {
	resp, err := http.Get(url)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	var locations LocationAreasResp
	err = json.Unmarshal(body, &locations)
		if err != nil {
			return LocationAreasResp{}, err
		}


	return locations, nil
}
