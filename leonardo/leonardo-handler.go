package leonardo

import (
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/internal"
	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func LeoHandler(c internal.ApiConfig) http.HandlerFunc {
	leoKey := c.LeoKey

	handlerFunct := func(w http.ResponseWriter, r *http.Request) {
		prompt := "museum piece photographed for art catalog: brutalist sculpture made of brushed bronze"

		err := GetLeonardoImg(prompt, leoKey)
		if err != nil {
			utils.RespondWithError(w, 500, "error creating image")
		}

		utils.RespondWithJson(w, 201, "image created")
	}

	return handlerFunct
}
