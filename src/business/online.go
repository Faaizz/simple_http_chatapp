package business

import (
	"fmt"

	"encoding/json"
	"net/http"

	"github.com/Faaizz/simple_http_chatapp/db"
	"github.com/Faaizz/simple_http_chatapp/types"
)

func OnlineHandler(w http.ResponseWriter, r *http.Request) {
	var u types.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		logger.Debugln(err)
		w.WriteHeader(400)
		fmt.Fprintf(w, "could not decode input")
		return
	}

	users, err := db.AvailableUsers(u)
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
