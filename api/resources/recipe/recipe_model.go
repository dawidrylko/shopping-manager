package recipe

import (
	"gopkg.in/mgo.v2/bson"

	"../ingredient"
)

// Recipe represent information about recipe
type Recipe struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name"`
	Description string        `json:"name"`

	Ingredients ingredient.Ingredients `json:"ingredients"`
}

// Recipes Recipes information about recipe list
type Recipes []Recipe
