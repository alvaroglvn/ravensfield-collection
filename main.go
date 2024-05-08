package main

import (
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
	ghostKey := os.Getenv("GHOST_KEY")
	ghostUrl := os.Getenv("GHOST_URL")
	leoKey := os.Getenv("LEONARDO_KEY")
	claudeKey := os.Getenv("CLAUDE_KEY")
	masterKey := os.Getenv("MASTER_KEY")

	//Build config
	config := internal.BuildConfig(port, openAiKey, ghostKey, ghostUrl, leoKey, claudeKey, masterKey)

	//Load router with CORS
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{}))

	//Endpoints
	//Upload images to Ghost and create empty articles
	router.With(handlers.CreateMasterKeyWare(config)).Post("/uploader", handlers.ImageUploader(config))
	//Generate new article from feature image
	router.With(handlers.CreateMasterKeyWare(config)).Post("/generate", handlers.GenTextAndPost(config))
	//post using openai
	// router.With(handlers.CreateMasterKeyWare(config)).Post("/ghostpost", handlers.PostArticle(config))
	// //post using claude
	// router.With(handlers.CreateMasterKeyWare(config)).Post("/claudepost", handlers.PostClaudeArticle(config))

	//Start server
	server := &http.Server{
		Addr:              ":" + config.Port,
		Handler:           router,
		ReadHeaderTimeout: 2 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
