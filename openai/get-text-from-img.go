package openai

import (
	"fmt"
	"github.com/alvaroglvn/ravensfield-collection/utils"
	"strings"
)

func GetTextFromImg(imgUrl, sample1, sample2, sample3, openaiKey string) (genText string, err error) {
	//Get text content from image
	genText, err = ImgDescribe(imgUrl, sample1, sample2, sample3, openaiKey)
	if err != nil {
		return "", fmt.Errorf("unable to create text based on image: %v", err)
	}

	return genText, nil
}

func FinalEdit(tunedText, openAiKey string) (caption, title, content string, err error) {
	editedText, err := AutoEdit(tunedText, openAiKey)
	if err != nil {
		return "", "", "", err
	}

	splitText := strings.Split(editedText, "\n")
	caption = splitText[2]
	title = splitText[0]
	contentParts := splitText[4:]
	content = strings.Join(contentParts, "\n")

	htmlDescript := utils.MarkdownToHTML([]byte(content))

	content = string(htmlDescript)

	return caption, title, content, nil
}
