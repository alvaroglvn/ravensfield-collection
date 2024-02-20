package openai_req

import (
	"encoding/json"
	"fmt"
)

type DalleRequest struct {
	Prompt         string `json:"prompt"`
	Model          string `json:"model"`
	NumberImgs     int    `json:"n"`
	Quality        string `json:"quality"`
	ResponseFormat string `json:"response_format"`
	Size           string `json:"size"`
	Style          string `json:"style"`
}

type DalleResponse struct {
	Created int     `json:"created"`
	Data    []Image `json:"data"`
}

type Image struct {
	URL string `json:"url"`
}

func GetDalleImg(prompt string) (string, error) {
	dalleRequest := DalleRequest{
		Prompt:         prompt,
		Model:          "dall-e-3",
		NumberImgs:     1,
		Quality:        "hd",
		ResponseFormat: "url",
		Size:           "1024x1024",
		Style:          "vivid",
	}

	var dalleResponse DalleResponse

	var dalleEndpoint = "https://api.openai.com/v1/images/generations"

	respBody, err := OpenAIConnect(dalleRequest, dalleEndpoint)
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
