package repository

import (
	"context"
	"time"

	"github.com/ShoreLab/shorelab-backend/lib/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) GetProjects() ([]*model.Project, error) {
	p := []*model.Project{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	col := r.db.MongoDBDatabase.Collection("projects")

	cursor, err := col.Find(ctx, bson.M{}, options.Find().SetSort(bson.D{{Key: "status", Value: 1}}))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var project model.Project
		if err := cursor.Decode(&project); err != nil {
			return nil, err
		}
		p = append(p, &project)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return p, nil
}
