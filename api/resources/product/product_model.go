package product

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Product represents information about product
type Product struct {
	ID            bson.ObjectId `json:"id" bson:"_id"`
	Name          string
	Manufacturer  string
	Ean           string
	Energy        string
	NetWeight     string
	UnitOfMeasure string

	Price Prices
}

// Products represents information about product list
type Products []Product

type Price struct {
	Date  time.Time `json:"date"`
	Value float32
}

type Prices []Price
