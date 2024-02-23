package openai

import (
	"encoding/json"
	"fmt"
)

func ArticleCreator(openAiKey string) (string, error) {

	userText, err := convertToPrompt("openai/prompts/article-user.txt")
	if err != nil {
		return "", fmt.Errorf("error gathering user text: %v", err)
	}

	systemText, err := convertToPrompt("openai/prompts/img-describe-system.txt")
	if err != nil {
		return "", fmt.Errorf("error gathering system text: %v", err)
	}

	compRequest := CompRequest{
		Model: "gpt-4",
		Messages: []Message{
			{Role: "user",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: userText,
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

	var compResponse CompResponse

	compEndpoint := "https://api.openai.com/v1/chat/completions"

	respBody, err := OpenAIConnect(compRequest, compEndpoint, openAiKey)
	if err != nil {
		return "", fmt.Errorf("error connecting to OpenAI's API: %v", err)
	}

	err = json.Unmarshal(respBody, &compResponse)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response's body: %v", err)
	}

	article := compResponse.Choices[0].Message.Content

	return article, nil
}
