package handlers

import (
	"html/template"
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/internal"
	"github.com/alvaroglvn/ravensfield-collection/openai"
	"github.com/alvaroglvn/ravensfield-collection/utils"
)

type PageData struct {
	ImgUrl      string
	Description string
}

func BuildTestSiteHandler(c internal.ApiConfig) http.HandlerFunc {

	openAiKey := c.OpenAiKey

	handlerFunct := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("static/templates/layout.html"))

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

		err = tmpl.Execute(w, data)
		if err != nil {
			utils.RespondWithError(w, 500, "error executing html template")
		}

	}
	return handlerFunct
}
