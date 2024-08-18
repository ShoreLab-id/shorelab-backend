package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Location    string             `json:"location" bson:"location"`
	Status      uint8              `json:"status" bson:"status"`
	Type        string             `json:"type" bson:"type"`
	Description string             `json:"description" bson:"description"`
	Date        time.Time          `json:"date" bson:"date"`
	Price       int                `json:"price" bson:"price"`
	Image       string             `json:"image" bson:"image"`
}
