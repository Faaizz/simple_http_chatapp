package business

import (
	"fmt"

	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	logger.Debugln("checking health...")

	w.WriteHeader(200)
	fmt.Fprintln(w, "healthy")
}
