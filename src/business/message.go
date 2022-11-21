package business

import (
	"bytes"
	"fmt"
	"io"

	"encoding/json"
	"net/http"

	"github.com/Faaizz/simple_http_chatapp/db"
	"github.com/Faaizz/simple_http_chatapp/msg"
	"github.com/Faaizz/simple_http_chatapp/types"
)

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	logger.Debug("sending message...")

	rBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Errorln(err)
		msg := "could not decode input"
		logger.Errorln(msg)
		fmt.Fprint(w, msg)
		return
	}
	logger.Debugf("request body: \n%v", string(rBytes))

	var inMsg types.Message

	err = json.NewDecoder(bytes.NewReader(rBytes)).Decode(&inMsg)
	if err != nil {
		logger.Debugln(err)
		msg := "could not decode input"
		logger.Errorln(msg)
		w.WriteHeader(400)
		fmt.Fprint(w, msg)
		return
	}
	if inMsg.ConnectionID == "" || inMsg.Username == "" || inMsg.URL == "" {
		msg := "could not initiate connection. 'connectionId', 'username', and 'url' required"
		logger.Errorln(msg)
		w.WriteHeader(400)
		fmt.Fprint(w, msg)
		return
	}

	// get ConnectionID of the target User
	targetConnectionID, err := db.ConnectionID(inMsg.Username)
	if err != nil {
		logger.Errorln(err)
		w.WriteHeader(400)
		fmt.Fprintf(w, "could not get target connectionId")
		return
	}

	// get username of source User
	sourceUsername, err := db.Username(inMsg.ConnectionID)
	if err != nil {
		logger.Errorln(err)
		w.WriteHeader(400)
		fmt.Fprintf(w, "could not get source username")
		return
	}

	err = msg.Message(targetConnectionID, inMsg.Message, sourceUsername, inMsg.URL)
	if err != nil {
		logger.Errorln(err)
		msg := "could not send message"
		logger.Errorln(msg)
		w.WriteHeader(400)
		fmt.Fprint(w, msg)
		return
	}

	msg := fmt.Sprintf("sent message: %s\nto: %s\n", inMsg.Message, inMsg.Username)
	logger.Debugln(msg)
	_, err = fmt.Fprintln(w, msg)
	if err != nil {
		logger.Errorln(err)
	}
}
