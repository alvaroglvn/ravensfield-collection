package handlers

import (
	"html/template"
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/pkg/helpers"
	"github.com/alvaroglvn/ravensfield-collection/pkg/openai_req"
)

type PageData struct {
	ImgUrl      string
	Description string
}

func BuildSite(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/templates/layout.html"))

	prompt := helpers.PromptBuilder()

	imgUrl, err := openai_req.GetDalleImg(prompt)
	if err != nil {
		helpers.RespondWithError(w, 500, "error creating Dalle image")
	}

	imgText, err := openai_req.ImgDescribe(imgUrl)
	if err != nil {
		helpers.RespondWithError(w, 500, "error creating GPT text")
	}

	data := PageData{
		ImgUrl:      imgUrl,
		Description: imgText,
	}

	tmpl.Execute(w, data)
}
