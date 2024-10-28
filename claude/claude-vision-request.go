package claude

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/alvaroglvn/ravensfield-collection/internal"
	madlibsprompt "github.com/alvaroglvn/ravensfield-collection/madlibs-prompt"
	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func ClaudeVisionReq(imgUrl, claudeKey string) (string, error) {

	// Build random artist info
	artistInfo, err := madlibsprompt.GetArtistInfo()
	if err != nil {
		return "", err
	}

	// Build random story
	storyPrompt, err := madlibsprompt.ObjectHistory()
	if err != nil {
		return "", err
	}

	imageBase64, err := utils.ImgUrltoBase64(imgUrl)
	if err != nil {
		return "", fmt.Errorf("error converting image to base64: %s", err)
	}

	message := claudeMessage{
		Model:       "claude-3-5-sonnet-20241022",
		System:      "This is a creative writing exercise. The genre is weird fiction and speculative fiction. Be unique, bold, and have a strong literary flare. Originality is key.",
		MaxTokens:   1000,
		Temperature: 1,
		Messages: []Message{
			{Role: "user",
				Content: []interface{}{
					ContentImage{
						Type: "image",
						Source: Source{
							Type:      "base64",
							MediaType: "image/webp",
							Data:      imageBase64,
						},
					},
					ContentText{
						Type: "text",
						Text: fmt.Sprintf(`You are a prestigious art scholar and the curator of the exclusive Ravensfield Collection, an eclectic museum of the weird and wonderful.

						1. Style:
						1.1. Use an exciting and varied vocabulary.
						1.2. NEVER START with this introduction: "The [piece name] stands as one of the most enigmatic pieces...".
						1.3. Keep your use of adverbs to a minimum.
						1.4. Use strong and expressive verbs.
						1.5. Avoid clichés.
						1.6. Make the text enticicing and engaging.
						1.7. Create a great sense of dramatic pacing.
						1.8. Balance scholarly information with exciting storytelling.
						1.9. Every paragraph must give new information, never be repetitive.
						1.10. Make the characters' names uncanny and different.
						1.11. The article should be cohesive in style, genre, and dramatic themes.

						2. Structure:
						2.1. Title. It must be shorter than five words, catchy, and seductive.
						2.2. Museum Tag. Description of the artwork formatted as | Artist's name | Title (Year) | Medium |
						The year should match the art period or movement the artwork belongs to.
						About the artist, %s.
						The only exception is if the artwork is an archeological piece. In that case, the artist can be unknown and the year an approximation.

						2.3. Article content. This section must be formatted in markdown.
						2.3.1. Use the first two paragraphs to introduce the piece, its author, and highlight its uniqueness and relevancy.
						2.3.2. Use five paragraphs to describe the uncanny event related to this artwork: %s. Consider this section a piece of flash fiction that follows the three act structure of short stories. This story must be thematically related to the artwork.
						2.3.3. Use the last paragraph to bring the article together and explain how the piece affects audiences today.
						2.3.4. Between two paragraphs of your choosing, add a fictional blockquote about the artwork by a fictional expert. Format it as follows: "Quote" -Author (profession).

						Never title individual sections.

						Use this guidance to inspire your text, but never mention it directly in your result.

						Reply with the final version of your article when you're ready.
						`, artistInfo, storyPrompt),
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
