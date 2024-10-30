package claude

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ClaudeAuthorVoice(sample1, sample2, sample3, content, claudeKey string) (tunedText string, err error) {

	message := claudeMessage{
		Model:     "claude-3-5-sonnet-20241022",
		System:    "Be a literary editor.",
		MaxTokens: 1000,
		//Temperature: 1,
		Messages: []Message{
			{Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: fmt.Sprintf(`Analyze these samples and obtain the author's voice from them: %s, %s, %s.
						
						Apply that same author voice from the samples to this text, so both the samples and the text sound like as if written by the same person: %s.
						
						Please, be as loyal to the original text as possible.
						
						Respond only with the final version of the text.`, sample1, sample2, sample3, content),
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

	//fmt.Println(string(body))

	var claudeResponse ClaudeResponse

	err = json.Unmarshal(body, &claudeResponse)
	if err != nil {
		return "", fmt.Errorf("unmarshalling error: %s", err)
	}

	return claudeResponse.Content[0].Text, nil
}

func ClaudeAutoEdit(content, claudeKey string) (editedText string, err error) {
	message := claudeMessage{
		Model:     "claude-3-5-sonnet-20241022",
		System:    "Be a thorough but evenhanded literary editor.",
		MaxTokens: 1000,
		//Temperature: 1,
		Messages: []Message{
			{Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: fmt.Sprintf(`Please edit this text following this guidance without losing its structure or essence:

						1. If the first sentence begins: "[Artwork name] stands as one of the most enigmatic pieces in The Ravensfield Collection", change it.
						
						1. If the text is much longer than 600 words, shorten it to be closer to that limit.
						
						2. Improve the article's readability if it is too verbose. Delete clichés and unnecessary transitions.
						
						3. Cut paragraphs which are repeating a point already made. Change words that sound repetitive to a synonym.
						
						4. Make sure the text is correctly formatted in html, separating the paragraphs and formatting any blockquotes accordingly.
			
						Please respond with the edited text ONLY, no need to add the editor's notes.
						
						Here's the original text: %s`, content),
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

	//fmt.Println(string(body))

	var claudeResponse ClaudeResponse

	err = json.Unmarshal(body, &claudeResponse)
	if err != nil {
		return "", fmt.Errorf("unmarshalling error: %s", err)
	}

	return claudeResponse.Content[0].Text, nil
}
