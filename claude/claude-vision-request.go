package claude

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func ClaudeVisionReq(imgUrl, claudeKey string) (string, error) {

	system, err := utils.ConvertToPrompt("claude/prompts/system.txt")
	if err != nil {
		return "", fmt.Errorf("error loading system prompt: %s", err)
	}

	user, err := utils.ConvertToPrompt("claude/prompts/user.txt")
	if err != nil {
		return "", fmt.Errorf("error loading user prompt: %s", err)
	}

	imageBase64, err := imgUrltoBase64(imgUrl)
	if err != nil {
		return "", fmt.Errorf("error converting image to base64: %s", err)
	}

	message := claudeMessage{
		Model:     "claude-3-opus-20240229",
		System:    system,
		MaxTokens: 1000,
		Messages: []Message{
			{Role: "user",
				Content: []interface{}{
					ContentImage{
						Type: "image",
						Source: Source{
							Type:      "base64",
							MediaType: "image/png",
							Data:      imageBase64,
						},
					},
					ContentText{
						Type: "text",
						Text: user,
					},
				},
			},
		},
	}

	reqBody, err := json.Marshal(message)
	if err != nil {
		return "", fmt.Errorf("marshalling error: %s", err)
	}

	claudeEndpoint := "https://api.anthropic.com/v1/messages"

	req, err := http.NewRequest("POST", claudeEndpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", claudeKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("response error: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %s", err)
	}

	fmt.Println(string(body))

	var claudeResponse ClaudeResponse

	err = json.Unmarshal(body, &claudeResponse)
	if err != nil {
		return "", fmt.Errorf("unmarshalling error: %s", err)
	}

	return claudeResponse.Content[0].Text, nil
}
