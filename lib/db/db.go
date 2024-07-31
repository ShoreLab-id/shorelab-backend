package db

import (
	"context"
	"encoding/base64"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type DBConnections struct {
	CloudStorageClient *storage.Client
	StorageBucket      *storage.BucketHandle
}

func NewDBConnections(ctx context.Context) (*DBConnections, error) {
	b64, err := base64.StdEncoding.DecodeString(os.Getenv("GOOGLE_CREDENTIALS_BASE64"))
	if err != nil {
		log.Default().Println(err.Error())
		return nil, err
	}

	opt := option.WithCredentialsJSON(b64)
	s, err := storage.NewClient(ctx, opt)
	if err != nil {
		log.Default().Println(err.Error())
		return nil, err
	}

	bkt := s.Bucket(os.Getenv("GCLOUD_BUCKET"))
	return &DBConnections{
		CloudStorageClient: s,
		StorageBucket:      bkt,
	}, nil
}
