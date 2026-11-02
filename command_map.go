package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type locationArea struct {
	Next    *string `json:"next"`
	Prev    *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func commandMap(cfg *config) error {
	var area locationArea
	if cfg.next == nil {
		return errors.New("Error retrieving next url")
	}
	if value, ok := cfg.cache.Get(*cfg.next); ok {
		err := json.Unmarshal(value, &area)
		if err != nil {
			return errors.New("error unmarshalling data")
		}
	} else {
		res, err := http.Get(*cfg.next)
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return errors.New("issue reading response body")
		}
		if err != nil {
			return errors.New("Issue getting maps")
		}
		err = json.Unmarshal(data, &area)
		if err != nil {
			return errors.New("Issue decoding json")
		}

		cfg.cache.Add(*cfg.next, data)
	}
	for _, v := range area.Results {
		fmt.Println(v.Name)
	}
	cfg.prev = area.Prev
	cfg.next = area.Next
	return nil
}

func commandMapb(cfg *config) error {
	var area locationArea
	if cfg.prev == nil {
		return errors.New("previous URL nil")
	}
	if cached, ok := cfg.cache.Get(*cfg.prev); ok {
		err := json.Unmarshal(cached, &area)
		if err != nil {
			return errors.New("error unmarshalling data")
		}
	} else {
		res, err := http.Get(*cfg.prev)
		if err != nil {
			return errors.New("issue with getting response from previous URL")
		}
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return errors.New("Issue reading response body")
		}
		err = json.Unmarshal(data, &area)
		if err != nil {
			return errors.New("issue decoding response body")
		}

		cfg.cache.Add(*cfg.prev, data)
	}

	for _, v := range area.Results {
		fmt.Println(v.Name)
	}
	cfg.next = area.Next
	cfg.prev = area.Prev
	return nil
}
