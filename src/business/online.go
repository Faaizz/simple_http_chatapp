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

func OnlineHandler(w http.ResponseWriter, r *http.Request) {
	logger.Debugln("connecting user...")

	rBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Errorln(err)
		msg := "could not decode request into bytes"
		logger.Errorln(msg)
		w.WriteHeader(400)
		fmt.Fprint(w, msg)
		return
	}
	logger.Debugf("request body: \n%v", string(rBytes))

	var u types.User
	var inMsg types.Message

	err = json.NewDecoder(bytes.NewReader(rBytes)).Decode(&u)
	if err != nil {
		logger.Errorln(err)
		msg := "could not decode input"
		logger.Errorln(msg)
		w.WriteHeader(400)
		fmt.Fprint(w, msg)
		return
	}

	err = json.NewDecoder(bytes.NewReader(rBytes)).Decode(&inMsg)
	if err != nil {
		logger.Errorln(err)
		msg := "could not decode input"
		logger.Errorln(msg)
		w.WriteHeader(400)
		fmt.Fprint(w, msg)
		return
	}

	// get username of User
	un, err := db.Username(u.ConnectionID)
	if err != nil {
		logger.Errorln(err)
		msg := "could not obtain username"
		logger.Errorln(msg)
		w.WriteHeader(400)
		fmt.Fprint(w, msg)
		return
	}

	u.Username = un
	users, err := db.AvailableUsers(u)
	if err != nil {
		logger.Errorln(err)
		msg := "could not find users"
		logger.Errorln(msg)
		w.WriteHeader(400)
		fmt.Fprint(w, msg)
		return
	}
	logger.Debugf("response body: \n%v", users)

	// push message to client via WebSocket
	uStr, err := json.Marshal(users)
	if err != nil {
		logger.Errorln(err)
		msg := "could not push message via WebSocket. unable to marshal response"
		logger.Errorln(msg)
		w.WriteHeader(400)
		fmt.Fprint(w, msg)
		return
	}
	err = msg.Message(inMsg.ConnectionID, string(uStr), inMsg.Username, inMsg.URL)
	if err != nil {
		logger.Errorln(err)
		msg := "could not push message via WebSocket"
		logger.Errorln(msg)
		w.WriteHeader(400)
		fmt.Fprint(w, msg)
		return
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		logger.Errorln(err)
		msg := "could not decode users"
		logger.Errorln(msg)
		w.WriteHeader(400)
		fmt.Fprint(w, msg)
	}

}
