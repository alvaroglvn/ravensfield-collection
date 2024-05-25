package openai

import (
	"encoding/json"
	"fmt"

	madlibsprompt "github.com/alvaroglvn/ravensfield-collection/madlibs-prompt"
	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func ImgDescribe(imgURL, openAiKey string) (string, error) {

	//Make new random story prompt
	storyPrompt, err := madlibsprompt.BuildRandStory()
	if err != nil {
		return "", err
	}
	fmt.Println(storyPrompt)

	imgData, err := utils.ImgUrltoBase64(imgURL)
	if err != nil {
		return "", err
	}

	visionRequest := CompRequest{
		Model: "gpt-4o",
		Messages: []Message{
			{Role: "user",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: fmt.Sprintf(`You are a prestigious art scholar and the curator of the exclusive Ravensfield Collection. You are very knowledgeable in art history, but also have a talent for storytelling. Please, write a short article about the artwork in this picture.

						Please, take into account the following general guidance: 
						
						- The article must be 500 words. 
						
						- The article must be six paragraphs long.
						
						- The article must be exciting, and unique. Originality is key. Explore the uncanny. Be unexpected.
						
						- Use tropes from weird fiction, dark fantasy, science fiction, magical realism, or horror. For inspiration, think of the stories published in Pulp magazines.
						
						- Authors you may use for inspiration for themes or style: Algernon Blackwood, Edgar Allan Poe, H.P. Lovecraft, M.R. James, Ambrose Bierce, Ray Bradbury, Richard Matheson, Clive Barker, J.G. Ballard.
						
						- The text should flow, have dramatic pace, and avoid feeling repetitive.

						- Please, avoid clichés.

						- Keep your use of adverbs to a minimum. Use strong and expressive verbs.

						- Please, do not title the individual sections.
						
						- Never make a direct mention to these guidance in your article. 
			
						To build the article, please follow these steps:  
						
						Step 1 - Give your article a catchy and enticing title. It must be no longer than five words.
						
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
						
						Step 4 - Write four paragraphs narrating a supernatural event or legend related to this artwork that ties seamlessly with the content so far. This story should fit organically into the rest of your article.
						
						%s
						
						Step 5 - Write one paragraph that brings the whole article together. Describe how the artwork affects audiences today.   
						
						Step 6 - Between two sections of your choosing, add a fictional quote by a fictional character.
						Follow the structure: 
						"Quote" -Name, Title 
						
						For example: 
						"This piece is a colorful nightmare." 
						-John McDreams, filmmaker
						Format this quote in markdown blockquote.`, storyPrompt),
					},
					ImageContent{
						Type: "image_url",
						ImageURL: ImageURL{
							URL: fmt.Sprintf("data:image/webp;base64,%s", imgData),
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
		PresencePenalty: 0.75,
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

	// fmt.Println(string(respBody))

	description := visionResponse.Choices[0].Message.Content

	return description, nil
}

func AutoEdit(text, openAiKey string) (editedText string, err error) {
	chatRequest := CompRequest{
		Model: "gpt-4o",
		Messages: []Message{
			{
				Role: "system",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: "Be a thorough literary editor.",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: fmt.Sprintf(`Please edit this article following these steps:

						1. If the article is much longer than 500 words, shorten it to fit a word count closer to that limit.

						3. If the article's title is formatted in markdown, remove its formatting.

						4. If you find a blockquote, make sure to format it as such in markdown.
						
						5. Go through a developmental and copy edit to improve the articles narrative and flow. Improve its readability if it is too verbose. Delete clichés and unnecessary transitions.
						
						You have full freedom to rewrite or cut paragraphs that are subpar or don't make sense, but do your best to be as loyal to the original text as possible.

						7. Go through a proofread. Fix typos and grammatical errors.
						
						Please, respond only with the edited text, no need to add editor's notes.
						
						Here's the original text: %s`, text),
					},
				},
			},
		},
		MaxTokens: 1000,
	}

	var chatResponse CompResponse

	gptEndpoint := "https://api.openai.com/v1/chat/completions"

	respBody, err := utils.ExternalAIPostReq(chatRequest, gptEndpoint, openAiKey)
	if err != nil {
		return "", fmt.Errorf("error connecting to OpenAI's API: %v", err)
	}

	err = json.Unmarshal(respBody, &chatResponse)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response's body: %v", err)
	}

	// fmt.Println(string(respBody))

	editedText = chatResponse.Choices[0].Message.Content

	return editedText, nil
}
