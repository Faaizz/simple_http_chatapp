package business

import (
	"fmt"

	"encoding/json"
	"net/http"

	"github.com/Faaizz/simple_http_chatapp/msg"
	"github.com/Faaizz/simple_http_chatapp/types"
)

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	var inMsg types.Message

	err := json.NewDecoder(r.Body).Decode(&inMsg)
	if err != nil {
		logger.Debugln(err)
		w.WriteHeader(400)
		fmt.Fprint(w, "invalid message")
		return
	}
	if inMsg.ConnectionID == "" || inMsg.FromUsername == "" {
		fmt.Fprint(w, "could not initiate connection. 'connectionId' and 'from_username' required")
		return
	}

	err = msg.Message(inMsg.ConnectionID, inMsg.Message, inMsg.FromUsername)
	if err != nil {
		logger.Debugln(err)
		w.WriteHeader(400)
		fmt.Fprint(w, "could not send message")
		return
	}

	fmt.Fprint(w, "message sent")
}
