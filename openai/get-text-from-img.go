package openai

import (
	"fmt"
	"strings"

	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func GetTextFromImg(imgUrl, openaiKey string) (tag, title, descript string, err error) {
	fullText, err := imgDescribe(imgUrl, openaiKey)
	if err != nil {
		return "", "", "", fmt.Errorf("unable to create text based on image: %v", err)
	}

	splitText := strings.Split(fullText, "\n")
	tag = splitText[2]
	title = splitText[0]
	descriptParts := splitText[4:]
	descript = strings.Join(descriptParts, "\n")

	htmlDescript := utils.MarkdownToHTML([]byte(descript))

	descript = string(htmlDescript)

	return tag, title, descript, nil
}
