package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"encoding/json"
)

var client = &http.Client{Timeout: 10 * time.Second}

type planet struct {
	Name    string
	Climate string
	Terrain string
	Films   []string
}

type planets struct {
	Count   uint
	Results []planet
}

func getJSON(url string, target interface{}) {
	r, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&target)
}

func getPlanetMovieCount(name string) uint {
	var p planets

	getJSON(fmt.Sprintf("https://swapi.dev/api/planets?search=%s", name), &p)

	return p.Count
}
