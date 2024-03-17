package pipelines

import (
	"fmt"

	"github.com/alvaroglvn/ravensfield-collection/internal"
	"github.com/alvaroglvn/ravensfield-collection/openai"
)

func BuildTextFromImg(imgUrl string, config internal.ApiConfig) (tag, title, descript string, err error) {

	tag, title, descript, err = openai.GetTextFromImg(imgUrl, config.OpenAiKey)
	if err != nil {
		return "", "", "", fmt.Errorf("error creating text from image: %s", err)
	}

	return tag, title, descript, nil
}
