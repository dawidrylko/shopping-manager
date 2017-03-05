package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("shopping-manager api start")

	serverRun()
	fmt.Println("server start")
}

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Products []Product

func getProducts(responseWriter http.ResponseWriter, request *http.Request) {
	products := Products{
		Product{ID: 1, Name: "ogórek zielony"},
		Product{ID: 2, Name: "papryka czerwona"},
		Product{ID: 3, Name: "Sałata lodowa"},
	}

	json.NewEncoder(responseWriter).Encode(products)
}

func serverRun() {
	http.HandleFunc("/api/product", getProducts)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
