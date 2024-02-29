package openai

import (
	"encoding/json"
	"fmt"

	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func imgDescribe(imgURL, openAiKey string) (string, error) {

	userText, err := convertToPrompt("openai/prompts/img-describe-user.txt")
	if err != nil {
		return "", fmt.Errorf("error gathering user text: %v", err)
	}

	systemText, err := convertToPrompt("openai/prompts/img-describe-system.txt")
	if err != nil {
		return "", fmt.Errorf("error gathering system text: %v", err)
	}

	visionRequest := CompRequest{
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
		MaxTokens: 1000,
	}

	var visionResponse CompResponse

	visionEndpoint := "https://api.openai.com/v1/chat/completions"

	respBody, err := utils.ExternalAIPostReq(visionRequest, visionEndpoint, openAiKey)
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
