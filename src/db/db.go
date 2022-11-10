package db

import (
	"github.com/Faaizz/simple_http_chatapp/types"
)

var dba types.DBAdapter

func init() {
	dba = &types.DynamoDBAdapter{}
}

// CheckExists checks if table tableName exists
func CheckExists(tableName string) error {
	dba.SetTableName(tableName)
	return dba.CheckExists()
}

func PutConn(pcIn types.PutConnInput) error {
	return dba.PutConn(pcIn)
}

func Delete(data map[string]string) error {
	return nil
}

func GetUserConnId(username string) (string, error) {
	return "connectionId", nil
}
