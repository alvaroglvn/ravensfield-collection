package handlers

import (
	"html/template"
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/internal"
	"github.com/alvaroglvn/ravensfield-collection/openai"
	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func BuildTest2Handler(c internal.ApiConfig) http.HandlerFunc {

	openAiKey := c.OpenAiKey

	handlerFunct := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("static/templates/layout.html"))

		article, err := openai.ArticleCreator(c.OpenAiKey)
		if err != nil {
			utils.RespondWithError(w, 500, "error creating GPT text")
		}
		prompt, err := openai.PromptFromArticle(article, c.OpenAiKey)
		if err != nil {
			utils.RespondWithError(w, 500, "error creating prompt")
		}

		imgUrl, err := openai.GetDalleImg(prompt, openAiKey)
		if err != nil {
			utils.RespondWithError(w, 500, "error creating Dalle image")
		}

		data := PageData{
			ImgUrl:      imgUrl,
			Description: article,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			utils.RespondWithError(w, 500, "error executing html template")
		}

	}
	return handlerFunct
}
