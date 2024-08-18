package repository

import (
	"context"
	"errors"
	"time"

	"github.com/ShoreLab/shorelab-backend/lib/dto"
	"github.com/ShoreLab/shorelab-backend/lib/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) GetProjects() ([]*dto.Project, error) {
	var p []*dto.Project
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	col := r.db.MongoDBDatabase.Collection("projects")

	cursor, err := col.Find(ctx, bson.M{}, options.Find().SetSort(bson.D{{Key: "status", Value: 1}}))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var proj model.Project
		if err := cursor.Decode(&proj); err != nil {
			return nil, err
		}
		p = append(p, &dto.Project{
			ID:       proj.ID,
			Title:    proj.Title,
			Location: proj.Location,
			Status:   proj.Status,
			Type:     proj.Type,
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return p, nil
}

var ErrInvalidID = errors.New("invalid ID")

func (r *Repository) GetProjectByName(projectID string) (*dto.ProjectDetail, error) {
	var p model.Project
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	col := r.db.MongoDBDatabase.Collection("projects")
	_id, err := primitive.ObjectIDFromHex(projectID)
	if err != nil {
		return nil, ErrInvalidID
	}
	err = col.FindOne(ctx, bson.M{"_id": _id}).Decode(&p)
	if err != nil {
		return nil, err
	}
	return &dto.ProjectDetail{
		Title:       p.Title,
		Location:    p.Location,
		Status:      p.Status,
		Type:        p.Type,
		Description: p.Description,
		Date:        p.Date,
		Price:       p.Price,
		Image:       p.Image,
	}, nil
}
