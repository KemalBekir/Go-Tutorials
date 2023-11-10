package models

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Image struct {
	ID  primitive.ObjectID `json:"id" bson:"_id"`
	URL string             `json:"url" bson:"url"`
}

//TODO add required

type Cake struct {
	ID          primitive.ObjectID   `json:"id" bson:"_id"`
	CakeName    string               `json:"cakeName" bson:"cakeName"`
	Description string               `json:"description" bson:"description"`
	Price       float64              `json:"price" bson:"price"`
	Type        string               `json:"type" bson:"type"`
	Images      []Image              `json:"images" bson:"images"`
	Likes       []primitive.ObjectID `json:"likes" bson:"likes"`
	Owner       primitive.ObjectID   `json:"owner" bson:"owner"`
	OnOffer     bool                 `json:"onOffer" bson:"onOffer"`
	Discount    float64              `json:"discount" bson:"discount"`
	CreatedAt   time.Time            `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at" bson:"updated_at"`
}

func (cake *Cake) MarshalJSON() ([]byte, error) {
	type Alias Cake
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(cake),
	})
}
