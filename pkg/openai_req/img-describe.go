package openai_req

import (
	"encoding/json"
	"fmt"

	"github.com/alvaroglvn/ravensfield-collection/pkg/helpers"
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

func ImgDescribe(imgURL string) (string, error) {

	userText, err := helpers.ConvertToPrompt("pkg/openai_req/prompts/img-describe-user.txt")
	if err != nil {
		return "", fmt.Errorf("error gathering user text: %v", err)
	}

	systemText, err := helpers.ConvertToPrompt("pkg/openai_req/prompts/img-describe-system.txt")
	if err != nil {
		return "", fmt.Errorf("error gathering system text: %v", err)
	}

	visionRequest := VisionRequest{
		Model: "gpt-4-vision-preview",
		Messages: []Message{
			{Role: "user",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: userText,
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
						Text: systemText,
					},
				},
			},
		},
		MaxTokens: 300,
	}

	var visionResponse VisionResponse

	visionEndpoint := "https://api.openai.com/v1/chat/completions"

	respBody, err := OpenAIConnect(visionRequest, visionEndpoint)
	if err != nil {
		return "", fmt.Errorf("error connecting to OpenAI's API: %v", err)
	}

	err = json.Unmarshal(respBody, &visionResponse)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response's body: %v", err)
	}

	description := visionResponse.Choices[0].Message.Content

	return description, nil
}
