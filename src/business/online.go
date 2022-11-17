package business

import (
	"bytes"
	"fmt"
	"io"

	"encoding/json"
	"net/http"

	"github.com/Faaizz/simple_http_chatapp/db"
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

	err = json.NewDecoder(bytes.NewReader(rBytes)).Decode(&u)
	if err != nil {
		logger.Errorln(err)
		msg := "could not decode input"
		logger.Errorln(msg)
		w.WriteHeader(400)
		fmt.Fprint(w, msg)
		return
	}

	users, err := db.AvailableUsers(u)
	if err != nil {
		logger.Errorln(err)
		msg := "could not find users"
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
		fmt.Fprintf(w, msg)
	}

}
