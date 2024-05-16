package openai

import (
	"fmt"
	"strings"

	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func GetTextFromImg(imgUrl, openaiKey string) (title, caption, content string, err error) {
	fullText, err := ImgDescribe(imgUrl, openaiKey)
	if err != nil {
		return "", "", "", fmt.Errorf("unable to create text based on image: %v", err)
	}

	splitText := strings.Split(fullText, "\n")
	caption = splitText[2]
	title = splitText[0]
	contentParts := splitText[4:]
	content = strings.Join(contentParts, "\n")

	htmlDescript := utils.MarkdownToHTML([]byte(content))

	content = string(htmlDescript)

	return title, caption, content, nil
}
