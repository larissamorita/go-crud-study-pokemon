package service

import (
	"encoding/json"
	"fmt"
	"github.com/larissamorita/go-crud-study-pokemon/resource"
	"io"
	"log"
	"net/http"
)

func searchpkmR() {

	resp, err := http.Get("https://pokeapi.co/api/v2/pokedex/kanto")
	if err != nil {
		log.Fatalln(err)
	}

	bytesBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	responseStruct2 := resource.ResponsePKM{}

	err = json.Unmarshal(bytesBody, &responseStruct2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print(searchpkmR)
}
