package tools

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
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

func GetDalleImg(prompt string, openAIKey string) (string, error) {
	dalleRequest := DalleRequest{
		Prompt:         prompt,
		Model:          "dall-e-3",
		NumberImgs:     1,
		Quality:        "hd",
		ResponseFormat: "url",
		Size:           "1024x1024",
		Style:          "vivid",
	}

	requestBody, err := json.Marshal(dalleRequest)
	if err != nil {
		return "", err
	}

	var dalleEndpoint = "https://api.openai.com/v1/images/generations"
	request, err := http.NewRequest("POST", dalleEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+openAIKey)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var dalleResponse DalleResponse
	err = json.Unmarshal(body, &dalleResponse)
	if err != nil {
		return "", err
	}

	createdImgUrl := dalleResponse.Data[0].URL

	return createdImgUrl, nil
}
