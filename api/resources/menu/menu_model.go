package menu

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"../recipe"
)

type Menu struct {
	ID      bson.ObjectId  `json:"id" bson:"_id"`
	Date    time.Time      `json:"date"`
	Recipes recipe.Recipes `json:"recipes"`
}
