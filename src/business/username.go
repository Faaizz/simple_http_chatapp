package business

import (
	"bytes"
	"fmt"
	"io"

	"encoding/json"
	"net/http"

	"github.com/Faaizz/simple_http_chatapp/db"
	"github.com/Faaizz/simple_http_chatapp/misc"
	"github.com/Faaizz/simple_http_chatapp/types"
)

func init() {
	logger = misc.Logger()
}

func UsernameHandler(w http.ResponseWriter, r *http.Request) {
	logger.Debugln("assigning username...")

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

	var connIn types.User

	err = json.NewDecoder(bytes.NewReader(rBytes)).Decode(&connIn)
	if err != nil {
		logger.Errorln(err)
		msg := "could not decode input"
		logger.Errorln(msg)
		w.WriteHeader(400)
		fmt.Fprint(w, msg)
		return
	}

	err = db.SetUsername(types.User{
		ConnectionID: connIn.ConnectionID,
		Username:     connIn.Username,
	})
	if err != nil {
		logger.Errorln(err)
		msg := "could not set username"
		logger.Errorln(msg)
		w.WriteHeader(400)
		fmt.Fprint(w, msg)
		return
	}
	if connIn.ConnectionID == "" || connIn.Username == "" {
		msg := "could not set username. 'connectionId' and 'username' required"
		logger.Errorln(msg)
		w.WriteHeader(400)
		fmt.Fprint(w, msg)
		return
	}

	msg := fmt.Sprintf("assigned username: %s to connection id: %s", connIn.Username, connIn.ConnectionID)
	logger.Debugln(msg)
	_, err = fmt.Fprintln(w, msg)
	if err != nil {
		logger.Errorln(err)
	}
}
