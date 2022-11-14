package types

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// A DynamoDBAdapter provides a layer of abstraction for interaction an underlying AWS DynamoDB database
type MongoDBAdapter struct {
	Client    *mongo.Client
	DBName    string
	TableName string // collection name
}

func (dba *MongoDBAdapter) SetTableName(tn string) {
	dba.TableName = tn
}

// CheckExists would always return a nil because MongoDB creates a new database and/or collection if it doesn't exist
func (dba *MongoDBAdapter) CheckExists(ctx context.Context) error {
	return nil
}

// PutConn inserts a username and connectionId in the underlying DynamoDB table
func (dba *MongoDBAdapter) PutConn(ctx context.Context, pcIn User) error {
	err := dba.CheckUsername(ctx, pcIn.Username)
	if err != nil {
		return err
	}

	_, err = dba.Client.Database(dba.DBName).Collection(dba.TableName).InsertOne(
		ctx,
		bson.D{
			{Key: "username", Value: pcIn.Username},
			{Key: "connectionId", Value: pcIn.ConnectionID},
		},
	)
	return err
}

// CheckUsername checks if username already exists on DynamoDB table
func (dba *MongoDBAdapter) CheckUsername(ctx context.Context, username string) error {
	var res User

	col := dba.Client.Database(dba.DBName).Collection(dba.TableName)
	err := col.FindOne(
		ctx,
		bson.D{{Key: "username", Value: username}},
	).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return err
	}

	return fmt.Errorf("username '%s' already exists", username)
}

// AvailableUsers lists available users and their connection IDs
func (dba *MongoDBAdapter) AvailableUsers(ctx context.Context) ([]User, error) {
	return []User{}, nil
}
