package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/ghost"
	"github.com/alvaroglvn/ravensfield-collection/internal"
	"github.com/alvaroglvn/ravensfield-collection/pipelines"
	"github.com/alvaroglvn/ravensfield-collection/utils"
)

// MASTER KEY MIDDLEWARE
func CreateMasterKeyWare(c internal.ApiConfig) func(http.Handler) http.Handler {
	function := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Api-Key")
			if header == c.MasterKey {
				next.ServeHTTP(w, r)
				return
			}

			utils.RespondWithError(w, 401, "bad api key")

		})
	}
	return function
}

// HEALTH CHECK
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

// IMAGE GENERATOR
func ImgGenerator(config internal.ApiConfig) http.HandlerFunc {
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		err := pipelines.LeoToCloud(config)
		if err != nil {
			utils.RespondWithError(w, 500, fmt.Sprintf("image generation process failed: %s", err))
		}
	}
	return handlerFunc
}

// IMAGE UPLOADER
func ImageUploader(config internal.ApiConfig) http.HandlerFunc {
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		err := pipelines.CloudinaryToGhost(config)
		if err != nil {
			utils.RespondWithError(w, 500, fmt.Sprintf("uploading error: %s", err))
		}
	}
	return handlerFunc
}

// GENERATE POST WITH CLAUDE
func GenTextClaude(config internal.ApiConfig) http.HandlerFunc {
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		err := ghost.GenTextClaude(config)
		if err != nil {
			utils.RespondWithError(w, 500, fmt.Sprintf("failed to generate article: %s", err))
		}
	}
	return handlerFunc
}

// GENERATE POST WITH CHATGPT
func GenTextChatGpt(config internal.ApiConfig) http.HandlerFunc {
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		err := ghost.GenTextChatgpt(config)
		if err != nil {
			utils.RespondWithError(w, 500, fmt.Sprintf("failed to generate article: %s", err))
		}
	}
	return handlerFunc
}
