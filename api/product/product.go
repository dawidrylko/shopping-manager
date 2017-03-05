package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"../pkg/mux"
)

// Declare unexported products array
var products Products

// Resource method show all endpoint in product package
func Resource() {
	router := mux.NewRouter()

	products = productsDeserializer()

	router.HandleFunc("/product", getProducts).Methods("GET")
	router.HandleFunc("/product/{id}", getProduct).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func productsDeserializer() Products {
	raw, error := ioutil.ReadFile("./api/product/product.json")

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
