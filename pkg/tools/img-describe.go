package tools

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type VisionRequest struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

type Message struct {
	Role    string        `json:"role"`
	Content []interface{} `json:"content"`
}

type TextContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type ImageContent struct {
	Type     string   `json:"type"`
	ImageURL ImageURL `json:"image_url"`
}

type ImageURL struct {
	URL string `json:"url"`
}

type VisionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func ImgDescribe(imgURL string, openAIKey string) (string, error) {

	visionRequest := VisionRequest{
		Model: "gpt-4-vision-preview",
		Messages: []Message{
			{Role: "user",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: "Describe the artwork in the picture in four short paragraphs",
					},
					ImageContent{
						Type: "image_url",
						ImageURL: ImageURL{
							URL: imgURL,
						},
					},
				},
			},
			{
				Role: "system",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: "You are an art scholar and curator, with a knack for the mystic.",
					},
				},
			},
		},
		MaxTokens: 300,
	}

	reqBody, err := json.Marshal(visionRequest)
	if err != nil {
		return "", err
	}

	visionEndpoint := "https://api.openai.com/v1/chat/completions"
	request, err := http.NewRequest("POST", visionEndpoint, bytes.NewBuffer(reqBody))
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

	var visionResponse VisionResponse
	err = json.Unmarshal(body, &visionResponse)
	if err != nil {
		return "", err
	}

	description := visionResponse.Choices[0].Message.Content

	return description, nil
}
