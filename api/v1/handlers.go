package api

import (
	"net/http"
)

func DroneHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ahahhaasdasdasdasda122121121hahaha"))
	w.WriteHeader(http.StatusOK)
}
