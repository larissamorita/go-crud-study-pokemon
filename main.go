package main

import (
	"fmt"
	"github.com/larissamorita/go-crud-study-pokemon/service"
)

func main() {
	city := "Denver"
	estimatedCost := "160"
	sum := service.GetData(city, estimatedCost)
	fmt.Println(sum)
}
