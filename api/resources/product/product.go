package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"../../server"
)

type ProductController struct {
	session *mgo.Session
}

func NewProductController(session *mgo.Session) *ProductController {
	return &ProductController{session}
}

var products Products

// Resource method show all endpoint in product package
func Resource() {
	products = productsDeserializer()

	pc := NewProductController(server.GetSession())

	server.Router.HandleFunc("/product", getProducts).Methods("GET")
	server.Router.HandleFunc("/product/{id}", getProduct).Methods("GET")
	server.Router.HandleFunc("/product", pc.createProduct).Methods("POST")
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
	// params := mux.Vars(request)

	// for _, product := range products {
	// 	if product.ID == params["id"] {
	// 		json.NewEncoder(responseWriter).Encode(product)

	// 		return
	// 	}
	// }

	// json.NewEncoder(responseWriter).Encode(&Product{})
}

func (productController ProductController) createProduct(responseWriter http.ResponseWriter, request *http.Request) {
	// Stub an product to be populated from the body
	product := Product{}

	// Populate the product data
	json.NewDecoder(request.Body).Decode(&product)

	// Add an ID
	product.ID = bson.NewObjectId()

	// Write the user to mongo
	productController.session.DB("shopping-manager").C("products").Insert(product)

	// Marshal provided interface into JSON structure
	productJSON, _ := json.Marshal(product)

	// Write content-type, statuscode, payload
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(201)
	fmt.Fprintf(responseWriter, "%s", productJSON)
}
