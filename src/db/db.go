package db

import (
	"context"

	"github.com/Faaizz/simple_http_chatapp/types"
)

var dba types.DBAdapter

// set database adapter
func SetDatabaseAdapter(dbaInit types.DBAdapter) {
	dba = dbaInit
}

// CheckExists sets TableName on DatabaseAdapter and checks if table tableName exists
func CheckExists(tableName string) error {
	dba.SetTableName(tableName)
	return dba.CheckExists(context.TODO())
}

// PutConn adds an entry into the table with a connectionId
func PutConn(pcIn types.Connection) error {
	return dba.PutConn(context.TODO(), pcIn)
}

// ConnectionID gets the connection ID associated with the specified username
func ConnectionID(un string) (string, error) {
	return dba.ConnectionID(context.TODO(), un)
}

// SetUsername adds a username entry into the table for the corresponding connectionId
func SetUsername(pcIn types.User) error {
	return dba.SetUsername(context.TODO(), pcIn)
}

// Username gets the username associated with connID
func Username(connID string) (string, error) {
	return dba.Username(context.TODO(), connID)
}

// AvailableUsers lists available users and their connection IDs
func AvailableUsers(u types.User) ([]types.User, error) {
	return dba.AvailableUsers(context.TODO(), u)
}

func Delete(data map[string]string) error {
	return nil
}

func GetUserConnId(username string) (string, error) {
	return "connectionId", nil
}

func Disconnect(u types.User) error {
	return dba.Disconnect(context.TODO(), u)
}
