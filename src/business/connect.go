package business

import (
	"fmt"

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
	var connIn types.User

	err := json.NewDecoder(r.Body).Decode(&connIn)
	if err != nil {
		logger.Debugln(err)
		w.WriteHeader(400)
		fmt.Fprintf(w, "could not decode ConnIn")
	}

	err = db.PutConn(types.User{
		ConnectionID: connIn.ConnectionID,
		Username:     connIn.Username,
	})
	if err != nil {
		logger.Debugln(err)
		fmt.Fprintf(w, "could not initiate connection: %s", err)
		return
	}

	_, err = fmt.Fprintf(w, "connected username: %s with connection id: %s", connIn.Username, connIn.ConnectionID)
	if err != nil {
		logger.Errorln(err)
	}
}
