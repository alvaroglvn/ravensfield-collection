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

	imgURL := "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a6/BremerDom-Taufbecken.jpg/640px-BremerDom-Taufbecken.jpg"

	helpers.ImgDescribe(imgURL)

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
