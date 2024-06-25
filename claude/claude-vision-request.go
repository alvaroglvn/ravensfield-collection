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
		Model:       "claude-3-5-sonnet-20240620",
		System:      "This is a genre fiction creative writing exercise. Be unique, bold, and have a strong literary flare. Originality is key.",
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

						Please write a short article about the artwork in this picture following this guidance:
						
						- The article's length should be around 500 words.

						- Stylistically, use a varied vocabulary without sounding grandiloquent. Avoid the word "enigmatic", using a more exciting synonym instead. Also, keep you use of adverbs to a minimum, using strong and expressive verbs instead. Finally, avoid clichés.

						- Dramatically, make sure your article is engaging and enticing. Your article should have superb pacing and keep the readers interested. Balance your scholarly explanation as an art historian with some exciting storytelling.

						- Thematically, dive into the uncanny while using strong metaphors. Be exciting, unique, and unexpected. Make sure the whole article stays on theme, every paragraph offering information that strengthens the whole.

						- Structurally, the article should look like this:

						[Title: the catchy and seductive title of the article. Keep it shorter than five words.]

						[Museum tag: information about the piece formatted as
						| [Artist] | [Title (Year)] | [Medium] |

						If the piece is an archeological find, the author can be unkown and its year an approximation.

						Otherwise: %s. The artwork's year shuld reflect the art movement it belongs to and be historically accurate.]

						[In two paragraphs, introduce the piece and its author, highlighting its uniqueness and relevancy. If the artist is unkown, offer an interesting factoid instead.]

						[%s. In five paragraphs, describe a legend or supernatural event related to this artwork. Consider this section a piece of flash fiction with a beginning, middle, and end, but make sure to keep it on theme and cohesive with the rest of your article]

						[In one paragraph, bring your article together, explaining how it affects audiences today.]

						- Between two paragraphs of your choosing, add a fictional quote about the artwork by a fictional expert. Format it as a separate blockquote as follows: 
						"Quote" -Author (profession).

						- Never title each section. Make sure the text flows seamlessly, is cohesive, and maintains the same theme, tone, voice, and narrative pace.

						Please reply with the final version of your article when you're ready.
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
