package business

import (
	"fmt"

	"net/http"

	"github.com/Faaizz/simple_http_chatapp/db"
)

func DisconnectHandler(w http.ResponseWriter, r *http.Request) {
	err := db.Disconnect()
	if err != nil {
		logger.Debugln(err)
		w.WriteHeader(400)
		fmt.Fprintf(w, "could not disconnect")
	}

	fmt.Fprintln(w, "disconnected")
}
