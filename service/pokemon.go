package service

import (
	"encoding/json"
	"fmt"
	"github.com/larissamorita/go-crud-study-pokemon/resource"
	"io"
	"log"
	"net/http"
)

func SearchpkmR() {

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

	var sumPKMr []string
	output := resource.Output{}

	for _, batatas := range responseStruct2.PkmEntries {
		if string(batatas.PkmSpecies.PkmName[0]) == "r" {
			sumPKMr = append(sumPKMr, batatas.PkmSpecies.PkmName)

			//adicao do novo GET do conteudo de color
			resp, err := http.Get(batatas.PkmSpecies.PkmURL)
			if err != nil {
				log.Fatalln(err)
			}

			bytesBody2, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}

			responseStruct3 := resource.URLcolor{}

			err = json.Unmarshal(bytesBody2, &responseStruct3)
			if err != nil {
				log.Fatalln(err)
			}
			sumPKMr = append(sumPKMr, responseStruct3.Color.ColorName)

			mandioca := resource.DataPkm2{
				Name:  batatas.PkmSpecies.PkmName,
				Color: responseStruct3.Color.ColorName,
			}
			output.Data = append(output.Data, mandioca)
		}
	}
	responseBody4, err := json.Marshal(output)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print(string(responseBody4))

	// Resultado do primeiro exercicio fmt.Print(sumPKMr)
}
