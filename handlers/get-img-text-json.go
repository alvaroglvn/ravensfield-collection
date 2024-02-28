package handlers

import (
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/internal"
	"github.com/alvaroglvn/ravensfield-collection/openai"
	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func GetImgTextJson(c internal.ApiConfig) http.HandlerFunc {
	openAiKey := c.OpenAiKey

	handlerFunct := func(w http.ResponseWriter, r *http.Request) {
		prompt := utils.PromptBuilder()

		imgUrl, err := openai.GetDalleImg(prompt, openAiKey)
		if err != nil {
			utils.RespondWithError(w, 500, "error creating Dalle image")
		}

		imgText, err := openai.ImgDescribe(imgUrl, openAiKey)
		if err != nil {
			utils.RespondWithError(w, 500, "error creating GPT text")
		}

		data := PageData{
			ImgUrl:      imgUrl,
			Description: imgText,
		}

		utils.RespondWithJson(w, 202, data)

	}
	return handlerFunct
}
