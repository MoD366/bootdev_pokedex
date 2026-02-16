module github.com/MoD366/bootdev_pokedex

go 1.25.0

replace github.com/MoD366/pokeapi v0.0.0 => ../internal/pokeapi
replace github.com/MoD366/pokecache v0.0.0 => ../internal/pokecache

require github.com/MoD366/pokeapi v0.0.0
require github.com/MoD366/pokecache v0.0.0
