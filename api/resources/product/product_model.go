package product

import "gopkg.in/mgo.v2/bson"

// Product represents information about product
type Product struct {
	ID   bson.ObjectId `json:"id,omitempty"`
	Name string        `json:"name,omitempty"`
}

// Products represents information about product list
type Products []Product
