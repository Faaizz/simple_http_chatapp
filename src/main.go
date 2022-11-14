package main

import (
	"fmt"
	"os"

	"net/http"

	"github.com/gorilla/mux"

	"github.com/Faaizz/simple_http_chatapp/business"
	"github.com/Faaizz/simple_http_chatapp/db"
	"github.com/Faaizz/simple_http_chatapp/misc"
	"github.com/Faaizz/simple_http_chatapp/msg"
	"github.com/Faaizz/simple_http_chatapp/types"
)

func main() {
	// setup logger
	logger := misc.Logger()

	// setup DB
	dbType := os.Getenv("DB_TYPE")
	tn := os.Getenv("TABLE_NAME")

	var dba types.DBAdapter

	switch dbType {

	case "MONGODB":
		ctx, cancel, mc, dbName := misc.InitMongoDB()
		defer cancel()
		defer mc.Disconnect(ctx)
		dba = &types.MongoDBAdapter{
			Client: mc,
			DBName: dbName,
		}

	case "", "DYNAMODB":
		dba = &types.DynamoDBAdapter{}

	default:
		dba = &types.DynamoDBAdapter{}
	}

	db.SetDatabaseAdapter(dba)
	err := db.CheckExists(tn)
	if err != nil {
		logger.Fatalf("table does not exist %v", err)
	}

	// setup message gateway adapter
	var mga types.MsgGwAdapter
	mga = &types.AWSApiGwAdapter{}

	msg.SetMsgGwAdapter(mga)

	// setup routing
	r := mux.NewRouter()

	r.HandleFunc("/connect", business.ConnectHandler).Methods("POST")
	r.HandleFunc("/online", business.OnlineHandler).Methods("GET")
	r.HandleFunc("/disconnect", business.DisconnectHandler).Methods("GET")
	r.HandleFunc("/message", business.MessageHandler).Methods("POST")

	// listen for connections
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "80"
	}

	listenIpPort := fmt.Sprintf(":%s", port)

	logger.Infoln("starting server")
	logger.Fatal(http.ListenAndServe(listenIpPort, r))
}
