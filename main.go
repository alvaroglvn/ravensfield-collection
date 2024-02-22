package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alvaroglvn/ravensfield-collection/handlers"
	"github.com/alvaroglvn/ravensfield-collection/internal"
	"github.com/alvaroglvn/ravensfield-collection/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	utils.LoadEnv()

	config := internal.BuildConfig()

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{}))

	router.Get("/test", handlers.BuildSite)

	server := &http.Server{
		Addr:              config.Port,
		Handler:           router,
		ReadHeaderTimeout: 2 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
