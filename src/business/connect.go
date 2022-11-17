package business

import (
	"bytes"
	"fmt"
	"io"

	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/Faaizz/simple_http_chatapp/db"
	"github.com/Faaizz/simple_http_chatapp/misc"
	"github.com/Faaizz/simple_http_chatapp/types"
)

var logger *zap.SugaredLogger

func init() {
	logger = misc.Logger()
}

func ConnectHandler(w http.ResponseWriter, r *http.Request) {
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

	err = db.PutConn(types.User{
		ConnectionID: connIn.ConnectionID,
		Username:     connIn.Username,
	})
	if err != nil {
		logger.Errorln(err)
		msg := fmt.Sprintf("could not initiate connection: %s", err)
		logger.Errorln(msg)
		fmt.Fprint(w, msg)
		return
	}
	if connIn.ConnectionID == "" || connIn.Username == "" {
		fmt.Fprint(w, "could not initiate connection. 'connectionId' and 'username' required")
		return
	}

	msg := fmt.Sprintf("connected username: %s with connection id: %s", connIn.Username, connIn.ConnectionID)
	logger.Debugln(msg)
	_, err = fmt.Fprintln(w, msg)
	if err != nil {
		logger.Errorln(err)
	}
}
