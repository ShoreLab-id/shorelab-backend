package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Title    string             `json:"title" bson:"title"`
	Location string             `json:"location" bson:"location"`
	Status   uint8              `json:"status" bson:"status"`
	Type     string             `json:"type" bson:"type"`
}
