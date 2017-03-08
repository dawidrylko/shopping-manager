package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"../../pkg/mux"
	"../../server"
)

var products Products

// Resource method show all endpoint in product package
func Resource() {
	products = productsDeserializer()

	server.Router.HandleFunc("/product", getProducts).Methods("GET")
	server.Router.HandleFunc("/product/{id}", getProduct).Methods("GET")
}

func productsDeserializer() Products {
	raw, error := ioutil.ReadFile("./api/resources/product/product.json")

	if error != nil {
		fmt.Println(error.Error())
		os.Exit(1)
	}

	var products Products
	json.Unmarshal(raw, &products)

	return products
}

func getProducts(responseWriter http.ResponseWriter, request *http.Request) {
	json.NewEncoder(responseWriter).Encode(products)
}

func getProduct(responseWriter http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	for _, product := range products {
		if strconv.Itoa(product.ID) == params["id"] {
			json.NewEncoder(responseWriter).Encode(product)

			return
		}
	}

	json.NewEncoder(responseWriter).Encode(&Product{})
}
