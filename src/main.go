package main

import (
	"fmt"
	"os"

	"net/http"

	"github.com/gorilla/mux"

	"github.com/Faaizz/simple_http_chatapp/business"
	"github.com/Faaizz/simple_http_chatapp/db"
	"github.com/Faaizz/simple_http_chatapp/misc"
)

func main() {
	// setup logger
	logger := misc.Logger()

	// setup DB
	tn := os.Getenv("TABLE_NAME")
	err := db.CheckExists(tn)
	if err != nil {
		logger.Fatalf("table does not exist %v", err)
	}

	// setup routing
	r := mux.NewRouter()

	r.HandleFunc("/connect", business.ConnectHandler).Methods("POST")

	// listen for connections
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "80"
	}

	listenIpPort := fmt.Sprintf(":%s", port)

	http.ListenAndServe(listenIpPort, r)
}
