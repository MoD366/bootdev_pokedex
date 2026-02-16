package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	//"time"

	"github.com/MoD366/pokeapi"
	//"github.com/MoD366/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays this help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Show the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Show the Pokemon available at specific location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a Pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func startRepl(conf *pokeapi.Config) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		if scanner.Scan() {
			userInput := scanner.Text()
			cleanedInput := cleanInput(userInput)
			if len(cleanedInput) == 0 {
				continue
			}

			commands := getCommands()

			if comm, ok := commands[cleanedInput[0]]; ok == true {
				arg := ""
				if len(cleanedInput) > 1 {
					arg = cleanedInput[1]
				}
				err := comm.callback(conf, arg)
				if err != nil {
					fmt.Printf("Error running %s: %s\n", comm.name, err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func commandExit(conf *pokeapi.Config, arg string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *pokeapi.Config, arg string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")

	for _, val := range getCommands() {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}
	return nil
}

func commandMap(conf *pokeapi.Config, arg string) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if conf.Next != "" {
		url = conf.Next
	}

	response, err := pokeapi.CallPokeapiLocation(url, conf)

	if err != nil {
		return err
	}

	for _, loc := range response.Results {
		fmt.Println(loc.Name)
	}

	conf.Next = *response.Next
	if response.Prev != nil {
		conf.Prev = *response.Prev
	}

	return nil
}

func commandMapb(conf *pokeapi.Config, arg string) error {
	if conf.Prev == "" {
		return errors.New("you're on the first page")
	}

	url := conf.Prev

	response, err := pokeapi.CallPokeapiLocation(url, conf)

	if err != nil {
		return err
	}

	for _, loc := range response.Results {
		fmt.Println(loc.Name)
	}
	conf.Next = *response.Next
	if response.Prev != nil {
		conf.Prev = *response.Prev
	} else {
		conf.Prev = ""
	}

	return nil
}

func commandExplore(conf *pokeapi.Config, arg string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + arg

	response, err := pokeapi.CallPokeapiEncounters(url, conf)

	if err != nil {
		return err
	}

	for _, mon := range response.PokemonEncounters {
		fmt.Println(mon.Pokemon.Name)
	}

	return nil
}

func commandCatch(conf *pokeapi.Config, arg string) error {
	url := "https://pokeapi.co/api/v2/pokemon/" + arg

	response, err := pokeapi.CallPokeapiPokemon(url, conf)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...", arg)

	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	randomNumber := randomizer.Intn(500)

	fmt.Println("Random number generated was", randomNumber, "has to be at least", response.BaseExperience, "to catch successfully.")

	if randomNumber >= response.BaseExperience {
		fmt.Println(arg, "was caught!")
	} else {
		fmt.Println(arg, "escaped!")
	}

	return nil
}

func cleanInput(text string) []string {

	if text == "" {
		return []string{}
	}

	result := strings.Fields(strings.ToLower(text))

	return result
}
