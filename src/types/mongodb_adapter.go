package types

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// A MongoDBAdapter provides a layer of abstraction for interaction an underlying AWS MongoDB database
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

func (dba *MongoDBAdapter) PutConn(ctx context.Context, pcIn Connection) error {
	_, err := dba.Client.Database(dba.DBName).Collection(dba.TableName).InsertOne(
		ctx,
		bson.D{
			{Key: "connectionId", Value: pcIn.ConnectionID},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

// ConnectionID gets the connection ID associated with the specified username
func (dba *MongoDBAdapter) ConnectionID(ctx context.Context, un string) (string, error) {
	return "", nil
}

// PutConn inserts a username and connectionId in the underlying MongoDB table
func (dba *MongoDBAdapter) SetUsername(ctx context.Context, pcIn User) error {
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
	if err != nil {
		return err
	}

	return nil
}

// Username gets the username associated with connID
func (dba *MongoDBAdapter) Username(ctx context.Context, connID string) (string, error) {
	return "", nil
}

// CheckUsername checks if username already exists on MongoDB table
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
func (dba *MongoDBAdapter) AvailableUsers(ctx context.Context, u User) ([]User, error) {
	return []User{}, nil
}

// Disconnect disconnects current User by deleting the user from DB
func (dba *MongoDBAdapter) Disconnect(ctx context.Context, u User) error {
	return nil
}
