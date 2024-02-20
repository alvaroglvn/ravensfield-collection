package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alvaroglvn/ravensfield-collection/configs"
	"github.com/alvaroglvn/ravensfield-collection/initiate"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("error loading env file: %v", err)
		return
	}

	config := configs.BuildConfig()

	router := initiate.RouterInit()

	server := &http.Server{
		Addr:              config.Port,
		Handler:           router,
		ReadHeaderTimeout: 2 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
