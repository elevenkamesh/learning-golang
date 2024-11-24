package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	performFormdata()
}
func performFormdata() {

	val := url.Values{}
	val.Add("name", "kamesh")
	fmt.Println(val)
	res, err := http.PostForm("/post", val)

	if err != nil {
		fmt.Println("Error:", err)
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
}
