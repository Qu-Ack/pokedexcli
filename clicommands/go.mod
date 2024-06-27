module github.com/Qu-Ack/pokedexcli/clicommands

replace (
		github.com/Qu-Ack/pokedexcli/pokeapirequest v0.0.0 => ../pokeapirequest
)

require(
		github.com/Qu-Ack/pokedexcli/pokeapirequest v0.0.0
)


go 1.22.4
