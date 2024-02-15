package main

import (
	// "fmt"
	"log"
	"net/http"
	"time"

	"github.com/alvaroglvn/ravensfield-collection/configs"
	"github.com/alvaroglvn/ravensfield-collection/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("warning: assuming default configuration. .env unreadable: %v", err)
	}
	config := configs.BuildConfig()

	// fmt.Println("running")

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{}))

	router.Get("/img", handlers.GetDalleImg)

	server := &http.Server{
		Addr:              config.Port,
		Handler:           router,
		ReadHeaderTimeout: 2 * time.Second,
	}

	// fmt.Println(server.Addr)

	log.Fatal(server.ListenAndServe())
}
