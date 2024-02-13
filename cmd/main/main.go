package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := ":" + os.Getenv("PORT")

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{}))

	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
