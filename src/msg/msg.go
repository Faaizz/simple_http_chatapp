package msg

import (
	"context"

	"github.com/Faaizz/simple_http_chatapp/types"
)

var mga types.MsgGwAdapter

func SetMsgGwAdapter(mgaInit types.MsgGwAdapter) {
	mga = mgaInit
}

func Message(cID, msg, fromUsername, url string) error {
	return mga.Message(context.TODO(), cID, msg, fromUsername, url)
}
