package repository

import (
	"context"
	"time"

	"github.com/ShoreLab/shorelab-backend/lib/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) GetUser(userEmail string) (*model.User, error) {
	var u model.User
	ctx, cancel := context.WithTimeout(r.ctx, 5*time.Second)
	defer cancel()

	col := r.db.MongoDBDatabase.Collection("users")

	err := col.FindOne(ctx, bson.M{"email": userEmail}).Decode(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
