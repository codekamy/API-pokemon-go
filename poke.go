// Usaremos a pokeapi (lá tem tudo sobre pokemon)
// para consultar: http://pokeapi.co/api/v2/pokedex/kanto/
// usaremos o verbo GET

// OUTROS VERBOS HTTP:
// GET: irá listar todos os registros/pokemons que se encontram no http, do arquivo JSON
// POST: Adiciona um novo registro
// DELETE: Remove um registro
// PUT e PATCH: Editar um registro
// uma vai ter o papel RESPONSE
// uma POKEMON
// uma POKEMONSPECIES

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Response struct {
	name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	numero  int            `json:"entry_number"`
	especie PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	name string `json:"name"`
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := io.ReadAll(response.Body) // dados --> bytes
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

	//DESEMPACOTANDO

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.name)
	fmt.Println(responseObject.Pokemon)

	for _, pokemon := range responseObject.Pokemon {
		fmt.Println(pokemon.especie.name)
	}

}
