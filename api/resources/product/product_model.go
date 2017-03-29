package product

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Product represents information about product
type Product struct {
	ID            bson.ObjectId `json:"id" bson:"_id"`
	Name          string        `json:"name"`
	Manufacturer  string        `json:"manufacturer"`
	Ean           string        `json:"ean"`
	Energy        string        `json:"energy"`
	NetWeight     string        `json:"netWeight"`
	UnitOfMeasure string        `json:"unitOfMeasure"`
	Price         Prices        `json:"price"`
}

// Products represents information about product list
type Products []Product

type Price struct {
	Date  time.Time `json:"date"`
	Value float32   `json:"value"`
}

type Prices []Price
