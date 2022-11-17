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
	logger.Debugf("request body: \n%v", string(rBytes))

	var connIn types.Connection

	err = json.NewDecoder(bytes.NewReader(rBytes)).Decode(&connIn)
	if err != nil {
		logger.Errorln(err)
		msg := "could not decode input"
		logger.Errorln(msg)
		w.WriteHeader(400)
		fmt.Fprint(w, msg)
		return
	}

	err = db.PutConn(types.Connection{
		ConnectionID: connIn.ConnectionID,
	})
	if err != nil {
		logger.Errorln(err)
		msg := "could not initiate connection"
		logger.Errorln(msg)
		fmt.Fprint(w, msg)
		return
	}
	if connIn.ConnectionID == "" {
		msg := "could not initiate connection. 'connectionId' required"
		logger.Errorln(msg)
		w.WriteHeader(400)
		fmt.Fprint(w, msg)
		return
	}

	msg := fmt.Sprintf("connected user with connection id: %s", connIn.ConnectionID)
	logger.Debugln(msg)
	_, err = fmt.Fprintln(w, msg)
	if err != nil {
		logger.Errorln(err)
	}
}
