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
	fmt.Println(storyPrompt)

	objectAnecdote, err := madlibsprompt.ObjectAnecdote()
	if err != nil {
		return "", err
	}
	fmt.Println(objectAnecdote)

	artistInfo, err := madlibsprompt.GetArtistInfo()
	if err != nil {
		return "", err
	}
	fmt.Println(artistInfo)

	imgData, err := utils.ImgUrltoBase64(imgURL)
	if err != nil {
		return "", err
	}

	visionRequest := CompRequest{
		Model: "gpt-4o",
		Messages: []Message{
			{
				Role: "system",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: "This is a genre fiction, creative writing exercise. Be unique, bold, and have a strong literary flare. Originality is key.",
					},
				},
			},
			{Role: "user",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: "You are a prestigious art scholar and the curator of the exclusive Ravensfield Collection. Please write a short article about the artwork in this picture.",
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
				Role: "assistant",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: fmt.Sprintf(`Of course. I would be more than happy to talk about this unique and mysterious piece. It's, no doubt, one of the highlights of the Ravensfield Collection.

						%s
						
						How long would you like my article to be?`, storyPrompt),
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: "I was thinking a maximum of 500 words.",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: "Yes, 500 words shall be enough. Is there any stylistic considerations you would want me to take into account?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: "Yes, please. Use a varied vocabulary without sounding grandiloquent. Also, keep you use of adverbs to a minimum, using strong and expressive verbs instead. Finally, avoid clichés.",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: "Understood. What about dramatic considerations?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: "Just make sure your article is engaging and enticing. Your article should have superb pacing and keep the readers interested.",
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: "Of course. I will make sure to balance my scholarly explanation as an art historian with some exciting storytelling. Shall we discuss the structure of the article next?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: `Yes. Please, follow these steps:
						
						- Step 1: Give your article a short and catchy title in five words or less.
						
						- Step 2: Add the artwork's museum tag following this format: | [Artist] | [Title (Year)] | [Medium] | If the piece is an archeological find, the artist can be unkown and the year an approximation.
						
						- Step 3: Write the article's content itself. First, introduce us to the artwork and its author. Then, describe the uncanny anecdote surrounding the object. Finally, describe how it still affects audiences today.`,
					},
				},
			},
			{
				Role: "assistant",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: "Of course. This object has a long history, though. Is there a particular anecdote you would like me to describe in my article?",
					},
				},
			},
			{
				Role: "user",
				Content: []interface{}{
					TextContent{
						Type: "text",
						Text: fmt.Sprintf("%s", objectAnecdote),
					},
				},
			},
			// {
			// 	Role: "assistant",
			// 	Content: []interface{}{
			// 		TextContent{
			// 			Type: "text",
			// 			Text: "Before I respond with my article, would you like me to check some examples so I can mimic a particular author's voice?",
			// 		},
			// 	},
			// },
			// {
			// 	Role: "user",
			// 	Content: []interface{}{
			// 		TextContent{
			// 			Type: "text",
			// 			Text: fmt.Sprintf("That's a great idea. Please use these as examples: %s, %s, %s", sample1, sample2, sample3) ,
			// 		},
			// 	},
			// },
		},
		MaxTokens:       1000,
		FreqPenalty:     0.7,
		PresencePenalty: 0.7,
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
						
						Please, maintain the same text structure, its title, and its museum tag.

						Please, respond only with the final version of the text.
						
						`, sample1, sample2, sample3, genText),
					},
				},
			},
		},
		MaxTokens:       1000,
		FreqPenalty:     0.7,
		PresencePenalty: 0.7,
		Temperature:     1,
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

						2. If the main text is much longer than 500 words, shorten it to fit a word count closer to that limit.

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

	//fmt.Println(string(respBody))

	editedText = chatResponse.Choices[0].Message.Content

	return editedText, nil
}
