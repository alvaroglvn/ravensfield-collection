package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alvaroglvn/ravensfield-collection/configs"
	"github.com/alvaroglvn/ravensfield-collection/initiate"
	"github.com/alvaroglvn/ravensfield-collection/pkg/helpers"
	"github.com/joho/godotenv"
)

func main() {

	helpers.PromptBuilder()

	err := godotenv.Load()
	if err != nil {
		log.Printf("warning: assuming default configuration. .env unreadable: %v", err)
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
