package openai

import (
	"fmt"
	"strings"

	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func GetArticlePieces(openaiKey string) (imgUrl, tag, title, descript string, err error) {

	prompt := PromptBuilder()

	imgUrl, err = GetDalleImg(prompt, openaiKey)
	if err != nil {
		return imgUrl, tag, title, descript, fmt.Errorf("error creating dalle image: %v", err)
	}

	imgText, err := ImgDescribe(imgUrl, openaiKey)
	if err != nil {
		return imgUrl, tag, title, descript, fmt.Errorf("error creating image description: %v", err)
	}

	splitText := strings.Split(imgText, "\n")
	tag = splitText[2]
	title = splitText[0]
	descriptParts := splitText[4:]
	descript = strings.Join(descriptParts, "\n")

	htmlDescript := utils.MarkdownToHTML([]byte(descript))

	return imgUrl, tag, title, string(htmlDescript), nil
}
