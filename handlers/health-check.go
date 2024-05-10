package handlers

import (
	"log"
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/internal"
)

func HealthCheck(config internal.ApiConfig) http.HandlerFunc {
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, err := w.Write([]byte("OK"))
		if err != nil {
			log.Printf("Error writing data: %s", err)
			return
		}
	}

	return handlerFunc
}
