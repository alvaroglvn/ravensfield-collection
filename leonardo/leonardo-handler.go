package leonardo

import (
	"fmt"
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/internal"
	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func LeoHandler(c internal.ApiConfig) http.HandlerFunc {
	leoKey := c.LeoKey

	handlerFunct := func(w http.ResponseWriter, r *http.Request) {
		prompt := utils.PromptBuilder()

		imgId, err := CreateLeoImg(prompt, leoKey)
		if err != nil {
			utils.RespondWithError(w, 500, "error creating image")
		}

		imgUrl, err := GetLeoImgUrl(imgId, leoKey)
		if err != nil {
			utils.RespondWithError(w, 500, "error loading image url")
		}
		fmt.Println(imgUrl)
	}

	return handlerFunct
}
