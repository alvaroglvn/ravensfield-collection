package handlers

import (
	"html/template"
	"net/http"
	"os"

	"github.com/alvaroglvn/ravensfield-collection/openai"
	"github.com/alvaroglvn/ravensfield-collection/utils"
)

type PageData struct {
	ImgUrl      string
	Description string
}

func BuildSite(w http.ResponseWriter, r *http.Request) {

	utils.LoadEnv()
	openAiKey := os.Getenv("OPENAI_API_KEY")

	tmpl := template.Must(template.ParseFiles("static/templates/layout.html"))

	prompt := openai.PromptBuilder()

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
