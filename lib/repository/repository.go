package repository

import (
	"context"

	"github.com/ShoreLab/shorelab-backend/lib/db"
)

type Repository struct {
	db  *db.DBConnections
	ctx context.Context
}

func NewRepository(db *db.DBConnections, ctx context.Context) *Repository {
	return &Repository{
		db,
		ctx,
	}
}
