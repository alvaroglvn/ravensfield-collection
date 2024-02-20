package handlers

import (
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/pkg/helpers"
	"github.com/alvaroglvn/ravensfield-collection/pkg/openai_req"
)

type ImgText struct {
	Image string `json:"img"`
	Text  string `json:"text"`
}

func GetImgAndText(w http.ResponseWriter, r *http.Request) {

	prompt := helpers.PromptBuilder()

	imageUrl, err := openai_req.GetDalleImg(prompt)
	if err != nil {
		helpers.RespondWithError(w, 500, "Error connecting to Dalle generator")
		return
	}

	imgDescript, err := openai_req.ImgDescribe(imageUrl)
	if err != nil {
		helpers.RespondWithError(w, 500, "Error connecting to Vision")
	}

	response := ImgText{
		Image: imageUrl,
		Text:  imgDescript,
	}

	helpers.RespondWithJson(w, 202, response)
}
