package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alvaroglvn/ravensfield-collection/handlers"
	"github.com/alvaroglvn/ravensfield-collection/internal"
	"github.com/alvaroglvn/ravensfield-collection/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	//Load enviromental variables
	utils.LoadEnv()
	port := os.Getenv("PORT")
	openAiKey := os.Getenv("OPENAI_API_KEY")

	//Load database
	db, err := sql.Open("sqlite3", "/sqlite/db/art-museum.sqlite")
	if err != nil {
		log.Fatal("unable to load database")
	}
	defer db.Close()

	//Build config
	config := internal.BuildConfig(port, openAiKey, db)

	//Load router with CORS
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{}))
	//Endpoints
	router.Get("/test", handlers.BuildSite)

	//Start server
	server := &http.Server{
		Addr:              config.Port,
		Handler:           router,
		ReadHeaderTimeout: 2 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
