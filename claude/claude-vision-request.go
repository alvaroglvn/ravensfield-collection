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

	// system, err := utils.ConvertToPrompt("claude/prompts/system.txt")
	// if err != nil {
	// 	return "", fmt.Errorf("error loading system prompt: %s", err)
	// }

	// user, err := utils.ConvertToPrompt("claude/prompts/user.txt")
	// if err != nil {
	// 	return "", fmt.Errorf("error loading user prompt: %s", err)
	// }

	imageBase64, err := utils.ImgUrltoBase64(imgUrl)
	if err != nil {
		return "", fmt.Errorf("error converting image to base64: %s", err)
	}

	message := claudeMessage{
		Model:       "claude-3-opus-20240229",
		System:      "This is an immersive creative writing exercise. Be unique, bold, and have a literary flare.",
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
						Text: `You are a prestigious art scholar and the curator of the exclusive Ravensfield Collection. You are very knowledgeable in art history, but also have a talent for storytelling. Please, write a short article about the artwork in this picture.

						Please, take into account the following general guidance: 
						
						- It's imperative than your article is never longer than 500 words. 
						
						- The article must have only six paragraphs.
						
						- The article must be exciting, and unique. Originality is key. Explore the uncanny. Be unexpected.
						
						- Use tropes from weird fiction, dark fantasy, science fiction, magical realism, or horror. For inspiration, think of the stories published in Pulp magazines.
						
						- Authors you may use for inspiration for themes or style: Algernon Blackwood, Edgar Allan Poe, H.P. Lovecraft, M.R. James, Ambrose Bierce, Ray Bradbury, Richard Matheson, Clive Barker, J.G. Ballard.
						
						- The text should flow, have dramatic pace, and avoid feeling repetitive.
						
						- Never make a direct mention to these guidance in your article. 
						
						-  Please, do not title the individual sections.
						
						- The output should be formatted in markdown.
						
						To build the article, please follow these steps:  
						
						Step 1 - Give your article a catchy and enticing title. It must be no longer than five words. This title must be unformatted.
						
						Step 2 - Write a museum tag that follows this structure: 
						
						| [Artist] | [Title (Year)] | [Medium] | 
						
						- The artist must be an imaginary person. 
						
						- In some rare cases, the artist can be unknown or a collective. 
						
						- The year must relate to the art style of the artwork. If the artwork is an archeological piece, it can be an approximation. 
						
						- The medium might include materials if the artwork calls for it. 
						
						Here are some examples: 
						
						| John Jonason | Nightmare in Pink (1965) | Acrylic on canvas |  
						
						| Unknown | Bejewelled King Skull (c. 330 BC) | Obsidian |  
						
						| Mark and James Thompkins | The Gentlemen (2000) | Plexiglass and marble 
						
						Step 3 - Write one paragraph introducing the artwork. Describe why this piece is relevant and introduce us to the artist behind it. If the piece doesn't have a known author, give us a fictional historical factoid related to the piece. 
						
						Step 4 - Write four paragraphs narrating a supernatural event or legend related to this artwork.  This section must follow the previous paragraph seamlessly, while also showcasing its own narrative independence using a three-act structure with a setup, an inciting incident, rising action, midpoint reversal, climax and resolution.
						
						- Use these themes for inspiration:  wild creativity, artistic obsession, existential dread, professional jealousy, societal envy, rise and fall, love gone wrong, tragic love, creativity into madness, seeing through the veil, the afterlife, things that crawl at night, tapping into other realities, ancient gods, the mythical and the mundane, the forbidden, lost civilizations, the arcane, hidden magic. You may also mix more than one theme.
						
						- Avoid clichés such as \"rumor has it\", \"legend says\", or similar.
						
						Step 5 - Write one paragraph that brings the whole article together. Describe how the artwork affects audiences today.   
						
						Step 6 - Between two sections of your choosing, add a fictional quote by a fictional character. Make sure to tag this as a block-quote in your markdown. Follow the format: \"Quote\" -Name, Title 
						
						For example: \"This piece is a colorful nightmare.\" -John McDreams, filmmaker
						
						Step 7 - Before you give me your answer, make sure you have followed all of my instructions accurately and fulfilled every step.`,
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

	// fmt.Println(string(body))

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
