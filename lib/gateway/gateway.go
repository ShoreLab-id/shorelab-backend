package gateway

import (
	"context"

	"github.com/ShoreLab/shorelab-backend/lib/db"
	"github.com/ShoreLab/shorelab-backend/lib/service"
)

type Gateway struct {
	Service *service.Service
}

func NewGateway() (*Gateway, error) {
	ctx := context.Background()
	client, err := db.NewDBConnections(ctx)
	if err != nil {
		return nil, err
	}
	defer client.CloudStorageClient.Close()

	srv := service.NewService(client, ctx)

	return &Gateway{
		Service: srv,
	}, nil
}
