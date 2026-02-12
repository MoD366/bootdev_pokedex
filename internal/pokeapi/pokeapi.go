package pokeapi

import (
	"net/http"
	"time"
	"io"
	"encoding/json"
)

type LocationArea struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Prev *string `json:"previous"`
	Results []struct{
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"results"`
}

func CallPokeapi(url string) (LocationArea,error) {
	
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return LocationArea{}, err
	}
	
	resp, err := client.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locations := LocationArea{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationArea{}, err
	}

	return locations, nil
}