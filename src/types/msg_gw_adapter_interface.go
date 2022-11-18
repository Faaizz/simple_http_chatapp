package types

import (
	"context"
)

type MsgGwAdapter interface {
	Message(context.Context, string, string, string, string) error
}
