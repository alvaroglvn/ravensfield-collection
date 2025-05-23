package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(status)
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Error writing data: %s", err)
		return
	}
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
		return
	}

	type errorResponse struct {
		Error string `json:"error"`
	}
	RespondWithJson(w, code, errorResponse{
		Error: msg,
	})
}
