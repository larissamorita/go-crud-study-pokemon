package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Response struct {
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Data       []Data `json:"data"`
}

type Data struct {
	City          string     `json:"city"`
	Name          string     `json:"name"`
	EstimatedCost int        `json:"estimated_cost"`
	UserRating    UserRating `json:"user_rating"`
	Id            int        `json:"id"`
}

type UserRating struct {
	AverageRating float64 `json:"average_rating"`
	Votes         int     `json:"votes"`
}

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

	responseStruct := Response{}

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
