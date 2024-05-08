package handlers

import (
	"fmt"
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/internal"
	"github.com/alvaroglvn/ravensfield-collection/pipelines"
	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func ImageUploader(config internal.ApiConfig) http.HandlerFunc {
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		err := pipelines.CloudinaryToGhost(config)
		if err != nil {
			utils.RespondWithError(w, 500, fmt.Sprintf("uploading error: %s", err))
		}
	}
	return handlerFunc
}

func GenTextAndPost(config internal.ApiConfig) http.HandlerFunc {
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		err := pipelines.UpdateGentext(config)
		if err != nil {
			utils.RespondWithError(w, 500, fmt.Sprintf("failed to generate article: %s", err))
		}
	}
	return handlerFunc
}
