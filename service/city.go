package service

import (
	"encoding/json"
	"fmt"
	"github.com/larissamorita/go-crud-study-pokemon/resource"
	"io"
	"log"
	"net/http"
)

func GetData(city, estimatedCost string) int {
	resp, err := http.Get(fmt.Sprintf("https://jsonmock.hackerrank.com/api/food_outlets?city=%s&estimated_cost=%v", city, estimatedCost))
	if err != nil {
		log.Fatalln(err)
	}

	bytesBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	responseStruct := resource.Response{}

	err = json.Unmarshal(bytesBody, &responseStruct)
	if err != nil {
		log.Fatalln(err)
	}

	var sumVotes int
	for _, data := range responseStruct.Data {
		if data.UserRating.AverageRating > 3.5 {
			sumVotes += data.UserRating.Votes
		}
	}

	return sumVotes
}
