package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alvaroglvn/ravensfield-collection/configs"
	"github.com/alvaroglvn/ravensfield-collection/initiate"
	"github.com/alvaroglvn/ravensfield-collection/pkg/helpers"
)

func main() {
	helpers.LoadEnv()

	config := configs.BuildConfig()

	router := initiate.RouterInit()

	server := &http.Server{
		Addr:              config.Port,
		Handler:           router,
		ReadHeaderTimeout: 2 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
