package openai

import (
	"encoding/json"
	"fmt"
)

func PromptFromArticle(article, openAiKey string) (string, error) {

	compRequest := CompRequest{
		Model: "gpt-4",
		Messages: []Message{
			{Role: "user",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: fmt.Sprintf("Write a prompt optimized for Dalle3 to create the image of the artwork described in this article: %v. The prompt needs to be short and to the point and optmized to make an image that is absolutely realistic.", article),
					},
				},
			},
		},
		MaxTokens: 500,
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

	prompt := compResponse.Choices[0].Message.Content

	return prompt, nil
}
