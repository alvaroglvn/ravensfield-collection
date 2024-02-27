package openai

import (
	"encoding/json"
	"fmt"

	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func GetDalleImg(prompt, openAiKey string) (string, error) {

	dalleRequest := DalleRequest{
		Prompt:         prompt,
		Model:          "dall-e-3",
		NumberImgs:     1,
		Quality:        "standard",
		ResponseFormat: "url",
		Size:           "1024x1024",
		Style:          "vivid",
	}

	var dalleResponse DalleResponse

	var dalleEndpoint = "https://api.openai.com/v1/images/generations"

	respBody, err := utils.ExternalAIConnect(dalleRequest, dalleEndpoint, openAiKey)
	if err != nil {
		return "", fmt.Errorf("error connecting to OpenAI's API: %v", err)
	}

	err = json.Unmarshal(respBody, &dalleResponse)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response's body: %v", err)
	}

	createdUrl := dalleResponse.Data[0].URL

	return createdUrl, nil
}
