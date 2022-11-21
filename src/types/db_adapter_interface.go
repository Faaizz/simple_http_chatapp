package types

import (
	"context"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

// A DBAdapter provides a layer of abstraction for interaction with the underlying database
type DBAdapter interface {
	SetTableName(string)
	CheckExists(context.Context) error
	PutConn(context.Context, Connection) error
	ConnectionID(context.Context, string) (string, error)
	SetUsername(context.Context, User) error
	Username(context.Context, string) (string, error)
	AvailableUsers(context.Context, User) ([]User, error)
	Disconnect(context.Context, User) error
}
