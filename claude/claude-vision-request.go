package claude

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/alvaroglvn/ravensfield-collection/internal"
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
		Model:       "claude-3-opus-20240229",
		System:      system,
		MaxTokens:   1000,
		Temperature: 1,
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

	var claudeResponse ClaudeResponse

	err = json.Unmarshal(body, &claudeResponse)
	if err != nil {
		return "", fmt.Errorf("unmarshalling error: %s", err)
	}

	return claudeResponse.Content[0].Text, nil
}

func ClaudeTextElements(imgUrl string, config internal.ApiConfig) (title, tag, descript string, err error) {
	fullText, err := ClaudeVisionReq(imgUrl, config.ClaudeKey)
	if err != nil {
		return "", "", "", fmt.Errorf("error creating claude text: %s", err)
	}

	splitText := strings.Split(fullText, "\n")
	title = splitText[0]
	tag = splitText[2]
	descriptParts := splitText[4:]
	descript = strings.Join(descriptParts, "\n")

	htmlDescript := utils.MarkdownToHTML([]byte(descript))

	descript = string(htmlDescript)

	return title, tag, descript, nil
}
