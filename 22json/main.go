package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name  string `json:"name"`
	Price int    `json:"-"`
}

func main() {
	fmt.Println("welcome to json")
	EncodeJson()
}

func EncodeJson() {

	lamoCourse := []course{
		{
			"Go Programming",
			100,
		},
		{
			"Golang Web Development",
			200,
		},
	}

	// package this data

	finalJosn, _ := json.Marshal(lamoCourse)

	fmt.Println(string(finalJosn))
}
