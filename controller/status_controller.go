package controller

import (
	"encoding/json"
	"net/http"
)

//FindStatusEndpoint return status
func FindStatusEndpoint(w http.ResponseWriter, r *http.Request) {
	response(w, http.StatusOK)
}

func response(w http.ResponseWriter, code int) {
	response, _ := json.Marshal(true)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
