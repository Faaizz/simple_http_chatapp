package business

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/Faaizz/simple_http_chatapp/db"
	"github.com/Faaizz/simple_http_chatapp/types"
)

func DisconnectHandler(w http.ResponseWriter, r *http.Request) {
	var u types.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		logger.Debugln(err)
		w.WriteHeader(400)
		fmt.Fprintf(w, "could not decode input")
		return
	}

	err = db.Disconnect(u)
	if err != nil {
		logger.Debugln(err)
		w.WriteHeader(400)
		fmt.Fprintf(w, "could not disconnect")
		return
	}

	fmt.Fprintln(w, "disconnected")
}
