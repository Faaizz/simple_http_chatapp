package business

import (
	"fmt"

	"encoding/json"
	"net/http"

	"github.com/Faaizz/simple_http_chatapp/db"
)

func OnlineHandler(w http.ResponseWriter, r *http.Request) {
	users, err := db.AvailableUsers()
	if err != nil {
		logger.Debugln(err)
		w.WriteHeader(400)
		fmt.Fprintf(w, "could not find users")
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		logger.Debugln(err)
		w.WriteHeader(400)
		fmt.Fprintf(w, "could not decode users")
	}

}
