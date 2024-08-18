package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

type ProjectDetail struct {
	Title       string    `json:"title"`
	Location    string    `json:"location"`
	Status      uint8     `json:"status"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Price       int       `json:"price"`
	Image       string    `json:"image"`
}

type ProjectDetailResponse struct {
	Data *ProjectDetail `json:"data"`
}
