package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serve).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
	greeter()
}

func greeter() {
	fmt.Println("Hello, World!")
}

func serve(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello, World!</h1>"))
}
