package main

import "github.com/Vkanhan/go-pokedox/internal/pokeapi"

type config struct {
	pokeapiClient pokeapi.Client 
	nextLocationAreaURL *string
	previousLocationAreaURL *string 
}

func main() {

	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startRepl(&cfg)

}
