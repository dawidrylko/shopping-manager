package recipe

import "gopkg.in/mgo.v2/bson"

// Recipe represent information about recipe
type Recipe struct {
	ID bson.ObjectId `json:"id" bson:"_id"`
}

// Recipes Recipes information about recipe list
type Recipes []Recipe
