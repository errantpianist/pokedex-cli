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

type LocationArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	ID               int    `json:"id"`
    Name             string `json:"name"`
    Names            []struct {
        Name     string `json:"name"`
        Language struct {
            Name string `json:"name"`
            URL  string `json:"url"`
        } `json:"language"`
    } `json:"names"`
    PokemonEncounters []struct {
        Pokemon struct {
            Name string `json:"name"`
            URL  string `json:"url"`
        } `json:"pokemon"`
        VersionDetails []struct {
            MaxChance int `json:"max_chance"`
            Version   struct {
                Name string `json:"name"`
                URL  string `json:"url"`
            } `json:"version"`
        } `json:"version_details"`
    } `json:"pokemon_encounters"`
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

func GetLocationArea(areaName string, cache *pokecache.Cache) (LocationArea, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", areaName)

	if cachedData, ok := cache.Get(url); ok {
		fmt.Println("Using cached data for", url)
		var area LocationArea
		err := json.Unmarshal(cachedData, &area)
		return area, err
	}

	fmt.Println("Fetching data from", url)

	resp, err := http.Get(url)
	if err != nil {
		return LocationArea{}, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	cache.Add(url, body)

	var area LocationArea
	err = json.Unmarshal(body, &area)
	if err != nil {
		return LocationArea{}, err
	}

	return area, nil
}

type PokemonResp struct {
	ID int `json:"id"`
	Name string `json:"name"`
	BaseExperience int `json:"base_experience"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func GetPokemon(pokemonName string, cache *pokecache.Cache) (PokemonResp, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)

	if cachedData, ok := cache.Get(url); ok {
		fmt.Println("Using cached data for", url)
		var pokemon PokemonResp
		err := json.Unmarshal(cachedData, &pokemon)
		return pokemon, err
	}

	fmt.Println("Fetching data from", url)

	resp, err := http.Get(url)
	if err != nil {
		return PokemonResp{}, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResp{}, err
	}

	cache.Add(url, body)

	var pokemon PokemonResp
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return PokemonResp{}, err
	}

	return pokemon, nil
}
