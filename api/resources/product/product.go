package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"../../pkg/mux"
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
	pc := NewProductController(server.GetSession())

	server.Router.HandleFunc("/product", pc.getAll).Methods("GET")
	server.Router.HandleFunc("/product/{id}", pc.get).Methods("GET")
	server.Router.HandleFunc("/product", pc.create).Methods("POST")
}

func (productController ProductController) getAll(responseWriter http.ResponseWriter, request *http.Request) {
	products := Products{}

	// Fetch user
	if err := productController.session.DB("shopping-manager").C("products").Find(nil).All(&products); err != nil {
		responseWriter.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	productsJSON, _ := json.Marshal(products)

	// Write content-type, statuscode, payload
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(200)
	fmt.Fprintf(responseWriter, "%s", productsJSON)
}

func (productController ProductController) get(responseWriter http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	oid := bson.ObjectIdHex(id)

	product := Product{}

	// Fetch user
	if err := productController.session.DB("shopping-manager").C("products").FindId(oid).One(&product); err != nil {
		responseWriter.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	productJSON, _ := json.Marshal(product)

	// Write content-type, statuscode, payload
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(200)
	fmt.Fprintf(responseWriter, "%s", productJSON)
}

func (productController ProductController) create(responseWriter http.ResponseWriter, request *http.Request) {
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
	fmt.Fprintf(responseWriter, "%s", productJSON, "lalala")
}
