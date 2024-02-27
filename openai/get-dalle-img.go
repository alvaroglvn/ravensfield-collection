package openai

import (
	"encoding/json"
	"fmt"
	//"math/rand"
)

func GetDalleImg(prompt, openAiKey string) (string, error) {

	// sizes := [3]string{"1024x1024", "1792x1024", "1024x1792"}
	// randIndex := rand.Intn(len(sizes))
	// size := sizes[randIndex]

	dalleRequest := DalleRequest{
		Prompt:         prompt,
		Model:          "dall-e-3",
		NumberImgs:     1,
		Quality:        "standard",
		ResponseFormat: "url",
		Size:           "1792x1024",
		Style:          "vivid",
	}

	var dalleResponse DalleResponse

	var dalleEndpoint = "https://api.openai.com/v1/images/generations"

	respBody, err := OpenAIConnect(dalleRequest, dalleEndpoint, openAiKey)
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
