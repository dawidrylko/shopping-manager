package menu

import "gopkg.in/mgo.v2/bson"
import "time"

type Menu struct {
	ID      bson.ObjectId  `json:"id" bson:"_id"`
	date    time.Time      `json:"date"`
	recipes recipe.Recipes `json:"recipe.Recipes"`
}
