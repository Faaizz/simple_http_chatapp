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
	var connIn types.ConnInput

	err := json.NewDecoder(r.Body).Decode(&connIn)
	if err != nil {
		logger.Errorln(err)
		logger.Fatalf("could not decode ConnIn")
	}

	cID := "01234567890"
	err = db.PutConn(types.PutConnInput{
		ConnectionID: cID,
		Username:     connIn.Username,
	})
	if err != nil {
		logger.Errorln(err)
		logger.Errorln("could not initiate connection")
		fmt.Fprintf(w, "could not connect. %s", err)
		return
	}

	fmt.Fprintf(w, "connected username: %s with connection id: %s", connIn.Username, cID)
}
