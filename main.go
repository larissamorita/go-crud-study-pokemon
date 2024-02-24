package main

import (
	"encoding/json"
	"fmt"
	"github.com/larissamorita/go-crud-study-pokemon/model"
	"io"
	"log"
	"net/http"
)

func main() {
	city := "Denver"
	estimatedCost := "160"
	sum := GetData(city, estimatedCost)
	fmt.Println(sum)
}

func GetData(city, estimatedCost string) int {
	resp, err := http.Get(fmt.Sprintf("https://jsonmock.hackerrank.com/api/food_outlets?city=%s&estimated_cost=%v", city, estimatedCost))
	if err != nil {
		log.Fatalln(err)
	}

	bytesBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	responseStruct := model.Response{}

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
