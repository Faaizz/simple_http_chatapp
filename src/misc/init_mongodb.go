package misc

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitMongoDB() (context.Context, context.CancelFunc, *mongo.Client, string) {
	logger = Logger()

	// connect to MongoDB
	mu := os.Getenv("MONGODB_USER")
	if mu == "" {
		logger.Fatalf("please set MONGODB_USER env")
	}

	mp := os.Getenv("MONGODB_PASSWORD")
	if mp == "" {
		logger.Fatalf("please set MONGODB_PASSWORD env")
	}

	mHost := os.Getenv("MONGODB_HOST")
	if mHost == "" {
		logger.Fatalf("please set MONGODB_HOST env")
	}

	mPort := os.Getenv("MONGODB_PORT")
	if mPort == "" {
		logger.Fatalf("please set MONGODB_PORT env")
	}

	dbName := os.Getenv("MONGODB_DATABASE")
	if mPort == "" {
		logger.Fatalf("please set MONGODB_DATABASE env")
	}

	mUrl := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		mu,
		mp,
		mHost,
		mPort,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mUrl),
	)
	if err != nil {
		logger.Fatal(err)
	}

	// verify connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logger.Fatal(err)
	}

	return ctx, cancel, client, dbName
}
