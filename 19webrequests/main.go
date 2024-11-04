package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {

	fmt.Println("Web Request")

	res, err := http.Get(url)
	if err != nil {
		panic(err)

	}

	fmt.Println("response", res)

	defer res.Body.Close()

	databyte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databyte))
}
