package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	ID       primitive.ObjectID `json:"_id"`
	Title    string             `json:"title"`
	Location string             `json:"location"`
	Status   uint8              `json:"status"`
	Type     string             `json:"type"`
}

type ProjectListResponse struct {
	Data []*Project `json:"data"`
}
