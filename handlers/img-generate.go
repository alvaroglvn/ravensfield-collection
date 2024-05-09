package handlers

import (
	"fmt"
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/internal"
	"github.com/alvaroglvn/ravensfield-collection/pipelines"
	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func ImgGenerator(config internal.ApiConfig) http.HandlerFunc {
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		err := pipelines.LeoToCloud(config)
		if err != nil {
			utils.RespondWithError(w, 500, fmt.Sprintf("image generation process failed: %s", err))
		}
	}
	return handlerFunc
}
