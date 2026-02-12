package main

import(
	"strings"
	"fmt"
	"bufio"
	"os"
	"internal/pokeapi"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

type config struct {
	next string
	prev string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays this help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Show the next 20 locations",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Show the previous 20 locations",
			callback: commandMapb,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},

	}
}

func startRepl() {

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
				err := comm.callback()
				if err != nil {
					fmt.Println("Error running %s: %s", comm.name, err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")

	for _, val := range getCommands() {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}
	return nil
}

func commandMap(conf *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	
	if *conf.next != nil {
		url = conf.next 
	}

	err = callPokeapi(url)
	
	if err != nil {
		return err
	}


}


func cleanInput(text string) []string {
	
	if text == "" {
		return []string{}
	}

	result := strings.Fields(strings.ToLower(text))
	
	return result
}