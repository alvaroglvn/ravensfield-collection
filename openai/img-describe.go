package openai

import (
	"encoding/json"
	"fmt"

	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func ImgDescribe(imgURL, openAiKey string) (string, error) {

	// userText, err := utils.ConvertToPrompt("openai/prompts/img-describe-user.txt")
	// if err != nil {
	// 	return "", fmt.Errorf("error gathering user text: %v", err)
	// }

	// systemText, err := utils.ConvertToPrompt("openai/prompts/img-describe-system.txt")
	// if err != nil {
	// 	return "", fmt.Errorf("error gathering system text: %v", err)
	// }

	visionRequest := CompRequest{
		Model: "gpt-4o",
		Messages: []Message{
			{Role: "user",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: `You are a prestigious art scholar and the curator of the exclusive Ravensfield Collection. You are very knowledgeable in art history, but also have a talent for storytelling. Please, write a short article about the artwork in this picture.

						Please, take into account the following general guidance: 
						
						- It's imperative than your article is never longer than 500 words. 
						
						- The article must have only six paragraphs.
						
						- The article must be exciting, and unique. Originality is key. Explore the uncanny. Be unexpected.
						
						- Use tropes from weird fiction, dark fantasy, science fiction, magical realism, or horror. For inspiration, think of the stories published in Pulp magazines.
						
						- Authors you may use for inspiration for themes or style: Algernon Blackwood, Edgar Allan Poe, H.P. Lovecraft, M.R. James, Ambrose Bierce, Ray Bradbury, Richard Matheson, Clive Barker, J.G. Ballard.
						
						- The text should flow, have dramatic pace, and avoid feeling repetitive.

						- Please, avoid clichés.

						- Keep your use of adverbs to a minimum. Use strong and expressive verbs.
						
						- Never make a direct mention to these guidance in your article. 
						
						-  Please, do not title the individual sections.
						
						- The output should be formatted in markdown.
						
						To build the article, please follow these steps:  
						
						Step 1 - Give your article a catchy and enticing title. It must be no longer than five words. This title must be unformatted, not in markdown.
						
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
						
						- Avoid clichés such as "rumor has it", "legend says", or similar.
						
						Step 5 - Write one paragraph that brings the whole article together. Describe how the artwork affects audiences today.   
						
						Step 6 - Between two sections of your choosing, add a fictional quote by a fictional character. Make sure to tag this as a block-quote in your markdown. Follow the format: \"Quote\" -Name, Title 
						
						For example: \"This piece is a colorful nightmare.\" -John McDreams, filmmaker
						
						Step 7 - Before you give me your answer, make sure you have followed all of my instructions accurately and fulfilled every step.`,
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
						Text: "This is an immersive creative writing exercise. Be unique, bold, and have a literary flare.",
					},
				},
			},
		},
		MaxTokens:       1000,
		FreqPenalty:     0.5,
		PresencePenalty: 0.5,
		Temperature:     1,
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

	fmt.Println(string(respBody))

	description := visionResponse.Choices[0].Message.Content

	return description, nil
}
