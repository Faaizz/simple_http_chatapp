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
	PutConn(context.Context, User) error
	AvailableUsers(ctx context.Context) ([]User, error)
}
