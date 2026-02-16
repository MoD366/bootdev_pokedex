package main

import (
	"time"

	"github.com/MoD366/pokeapi"
	"github.com/MoD366/pokecache"
)

func main() {

	conf := pokeapi.Config{
		Next:  "",
		Prev:  "",
		Cache: pokecache.NewCache(5 * time.Minute),
	}

	startRepl(&conf)
}
