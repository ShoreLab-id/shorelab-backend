package service

import (
	"context"

	"github.com/ShoreLab/shorelab-backend/lib/db"
	"github.com/ShoreLab/shorelab-backend/lib/repository"
)

type Service struct {
	repository *repository.Repository
}

func NewService(db *db.DBConnections, ctx context.Context) *Service {
	r := repository.NewRepository(db, ctx)
	return &Service{
		r,
	}
}
