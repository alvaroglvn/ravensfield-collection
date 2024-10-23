package openai

import (
	"encoding/json"
	"fmt"

	madlibsprompt "github.com/alvaroglvn/ravensfield-collection/madlibs-prompt"
	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func ImgDescribe(imgURL, openAiKey string) (string, error) {

	//Make new random story prompt
	storyPrompt, err := madlibsprompt.ObjectHistory()
	if err != nil {
		return "", err
	}
	// fmt.Println(storyPrompt)

	objectAnecdote, err := madlibsprompt.ObjectAnecdote()
	if err != nil {
		return "", err
	}
	// fmt.Println(objectAnecdote)

	artistInfo, err := madlibsprompt.GetArtistInfo()
	if err != nil {
		return "", err
	}
	// fmt.Println(artistInfo)

	imgData, err := utils.ImgUrltoBase64(imgURL)
	if err != nil {
		return "", err
	}

	visionRequest := CompRequest{
		Model: "o1-mini",
		Messages: []Message{
			{
				Role: "system",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: "This is a creative writing exercise. The genre is weird fiction and speculative fiction. Be unique, bold, and have a strong literary flare. Originality is key.",
					},
				},
			},
			{Role: "user",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: fmt.Sprintf(`You are a prestigious art scholar and the curator of the exclusive Ravensfield Collection.
						
						Please write an article about the artwork in this picture. The article must be between 500 and 600 words. %s. Follow this guidance:

						1. Style:
						1.1. Use an exciting and varied vocabulary.
						1.2. Keep your use of adverbs to a minimum.
						1.3. Use strong and expressive verbs.
						1.4. Avoid clichés.
						1.5. Make the text enticicing and engaging.
						1.6. Create a great sense of dramatic pacing.
						1.7. Balance scholarly information with exciting storytelling.
						1.8. Every paragraph must give new information, never be repetitive.
						1.9. The article should be cohesive in style, genre, and dramatic themes. 

						2. Structure:
						2.1. Title. It must be shorter than five words, catchy, and seductive.
						2.2. Museum Tag. Description of the artwork formatted as | Artist's name | Title (Year) | Medium |
						The year should match the art period or movement the artwork belongs to.
						About the artist, %s.
						The only exception is if the artwork is an archeological piece. In that case, the artist can be unknown and the year an approximation.
						2.3. Article content. This section must be formatted in markdown.
						2.3.1. Use the first two paragraphs to introduce the piece, its author, and highlight its uniqueness and relevancy.
						2.3.2. Use five paragraphs to describe the uncanny event related to this artwork: %s. Consider this section a piece of flash fiction that follows the three act structure of short stories.
						2.3.3. Use the last paragraph to bring the article together and explain how the piece affects audiences today.
						2.3.4. Between two paragraphs of your choosing, add a fictional blockquote about the artwork by a fictional expert. Format it as follows: "Quote" -Author (profession).

						Never title individual sections.

						Use this guidance to inspire your text, but never mention it directly in your result.

						Reply with the final version of your article when you're ready.
						
						`, storyPrompt, objectAnecdote, artistInfo),
					},
					ImageContent{
						Type: "image_url",
						ImageURL: ImageURL{
							URL: fmt.Sprintf("data:image/webp;base64,%s", imgData),
						},
					},
				},
			},
		},
		MaxTokens:       1000,
		FreqPenalty:     0.6,
		PresencePenalty: 0.6,
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

func CaptureVoice(sample1, sample2, sample3, genText, openAiKey string) (tunedText string, err error) {
	chatRequest := CompRequest{
		Model: "gpt-4o",
		Messages: []Message{
			{
				Role: "system",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: "Be a precise literary editor.",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: fmt.Sprintf(`Please follow these steps:
						
						Step 1: Read these samples and obtain the author's voice from them: %s, %s, %s.

						Step 2: Edit this text using the author's voice you obtained, so this text and the samples sound like written by the same person: %s.
						
						Please, be as loyal to the original text as possible.

						Please, respond only with the final version of the text.
						
						`, sample1, sample2, sample3, genText),
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

	//fmt.Println(string(respBody))

	tunedText = chatResponse.Choices[0].Message.Content

	return tunedText, nil
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

						1. Please, never edit the article's title or its museum tag.

						3. If the main text is much longer than 500 words, shorten it to fit a word count closer to that limit.

						4. If you catch an un unformatted blockquote, format it in markdown.
						
						5. Improve the article's readability if it is too verbose. Delete clichés and unnecessary transitions.
						
						6. Rewrite nonsensical paragraphs, and cut those which are repeating a point already made. Nevertheless, stay as loyal to the original text as possible.

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

	//fmt.Println(string(respBody))

	editedText = chatResponse.Choices[0].Message.Content

	return editedText, nil
}
