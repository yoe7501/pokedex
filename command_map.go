package main

import (
	"encoding/json"
	"errors"
	"fmt"
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
	res, err := http.Get(*cfg.next)
	if err != nil {
		return errors.New("Issue getting maps")
	}
	err = json.NewDecoder(res.Body).Decode(&area)
	if err != nil {
		return errors.New("issue deocding json")
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
	res, err := http.Get(*cfg.prev)
	if err != nil {
		return errors.New("issue with getting response from previous URL")
	}
	err = json.NewDecoder(res.Body).Decode(&area)
	if err != nil {
		return errors.New("issue decoding response body")
	}
	for _, v := range area.Results {
		fmt.Println(v.Name)
	}
	cfg.next = area.Next
	cfg.prev = area.Prev
	return nil
}
