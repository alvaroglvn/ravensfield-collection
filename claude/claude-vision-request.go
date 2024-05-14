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

	imageBase64, err := imgUrltoBase64(imgUrl)
	if err != nil {
		return "", fmt.Errorf("error converting image to base64: %s", err)
	}

	message := claudeMessage{
		Model:       "claude-3-opus-20240229",
		System:      "Lean into creative writing and have a literary flare.",
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
						Text: "You are a prestigious art scholar and the curator of the exculsive Ravensfield Collection. You are scholarly but also slightly quirky. Please, write a short article about the artwork in this picture.",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Of course! I would be delighted to talk about this unique piece and its enigmatic story. Before I start, could you tell me the length of the article you are requesting?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Sure. The article should be around 500 words",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Understood. Something engaging and succint. Do you have any stylistic preferences?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Yes, I do. I would like you to use an exciting and varied vocabulary that uses jergon from the world of art curation without feeling pedantic. The article must feel both informative and literary. I want the text to flow seamlessly, maintain a tight literary pace, and avoid feeling repetitive.",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Do not worry. I can be a masterful storyteller. My article will not only inform, but also intrigue, entice, and excite its readers. What about the article's genre and mood?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "I would like the article to be very mysterious and fantastic. Let's use tropes from dark fantasy, horror, science fiction, magical realism, and weird fiction in general.",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Of course. All the objects from the Ravensfield Collection have uncanny origin stories, so that will be just perfect. Just to bring everything together, could you give me some literary references I can use for inspiration?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Yes, that's a great idea. What about Ray Bradbury, Algernon Blackwood, Harlan Ellison, H.P. Lovecraft, M.R. James, Ambrose Bierce, Edgar Allan Poe, J.G. Ballard, Richard Matheson, and Clive Barker? You can use these authors for inspiration for themes, mood, and stylistic tropes.",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "What a comprehensive list! No doubt these authors are great literary inspiration for me, as they dwell in the world of the weird and wonderful, just like the objects in this collection. Shall we discuss the structure of the article itself next?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Yes, let's do it. First, I want you to write a catchy and engaging title that will capture the reader's attention. Keep it short and to the point, 4 words maximum.",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Absolutely. What's the next step?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Next, write a museum tag that follows this format: [Artist] | [Title (Year)] | [Medium] |",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "I understand, but would like to clarify some points. The artist should be an imaginary person that exists in the universe of the Ravensfield Collection. Correct?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Yes, exactly. In some rare cases, the artist might be unknown or a collective, but only if that fits the object's story better.",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "That sounds great. Do you want to give me some guidance about the year field?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Yes. Let's keep the year related to the art movement the object belongs to.",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Yes, that makes sense. What about if the piece is an archeological object?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "In that case, an approximation should be just fine.",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Yes, that should work. Should the medium include specific materials?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Only when the artwork calls for that specificity. For example, a sculpture that uses different materials.",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "I understand. Maybe you can provide me with an example just to double check your guidance for the museum tag.",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Here are some examples: | John Jonason | Nightmare in Pink (1965) | Acrylic on canvas |, | Unknown | Bejewelled King Skull (c. 330 BC) | Obsidian |, | Mark and James Thompkins | The Gentlemen (2000) | Plexiglass and marble |",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Perfect. I think that's all I need regarding the museum tag. Shall we move on to the next step?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Yes. Now we are moving on to the article itself. The first paragraph should be an introduction to the artwork: why this piece is relevant, how it relates to its art style or a particular moment in time, and a bit about the artist.",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "I understand. What should I talk about if the artist is not known?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "What about a historical or legendary factoid related to the piece?",
					},
				},
			},
			{

				Role: "assistant",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "That's a great solution. Should we move on to the next step?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Yes. Next, we need to dive into the object's story. It will tie the artwork to a supernatural legend or event, all of it built around the creative process behind the object and the artist's state of mind.",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Fantastic! All the objects in the collection have amazing and uncanny origin stories that I cannot wait to share. How should we structure this part of the article?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Let's think of this part as its own flash fiction story, but making sure it flows seamlessly with the rest of the content.",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "I agree. The artwork's story should be what entices the reader the most but it should feel organic within the article. Should I follow a classic three act structure for this particular part?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "That's a great idea. That way this section will have some narrative independence. Let's keep it simple. Act 1 will have the setup and inciting incident. Act 2 will show rising action and a midpoint reversal. Act 3 will describe the climax and the resolution.",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "I understand. Who should the protagonist be?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "The artist, an art collector, someone who owned the artwork at a particular moment, or even the object itself.",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Understood. What about themes?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Here are some themes you can explore: wild creativity, artistic obsession, existential dread, professional jealousy, societal envy, rise and fall, love gone wrong, tragic love, creativity into madness, seeing through the veil, the afterlife, things that crawl at night, tapping into other realities, ancient gods, the mythical and the mundane, the forbidden, lost civilizations, the arcane, hidden magic. You can pick one of these, mix and match, or use something else that is similar.",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Fantastic. I think I have all the information I need to write this section of the article. Shall we move on?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Sure. For the last section, write a paragraph that brings the article together. It should also describe how the artwork affects the visitors to the Ravensfield Collection today, and also keep the readers guessing.",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "That sounds like a marvelous way to wrap the article up. Is there anything else you would like to add?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Yes, please. between two sections of your choosing, add a fictional quote by a fictional character. Make sure to mark this as a blockquote in your markdown. Follow the format: Quote -Name, Title. For example:This piece is a colorful nightmare. -John McDreams, filmmaker",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Of course. I will add that quote to the article. Would you like me to format the output?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "Yes, please. Format the output in markdown. But please, leave the article's title unformatted.",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "No problem. I will share the article in markdown. Anything else?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					ContentText{
						Type: "text",
						Text: "I believe that's it. Share the article when you're ready. Before you deliver it, make sure to double check you've taken into account everything we discussed.",
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
