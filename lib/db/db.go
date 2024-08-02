package db

import (
	"context"
	"encoding/base64"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/api/option"
)

type DBConnections struct {
	CloudStorageClient *storage.Client
	StorageBucket      *storage.BucketHandle
	MongoDBClient      *mongo.Client
	MongoDBDatabase    *mongo.Database
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

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION")).SetServerAPIOptions(serverAPI)

	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	mc, err := mongo.Connect(timeoutCtx, opts)
	if err != nil {
		log.Default().Println(err.Error())
		return nil, err
	}

	mdb := mc.Database("shorelab-dev")

	return &DBConnections{
		CloudStorageClient: s,
		StorageBucket:      bkt,
		MongoDBClient:      mc,
		MongoDBDatabase:    mdb,
	}, nil
}
