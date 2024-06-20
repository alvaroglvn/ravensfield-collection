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
						Text: fmt.Sprintf(`You are a prestigious art scholar and the curator of the exclusive Ravensfield Collection.
						
						Please write a short article about the artwork in this picture. Its length should be around 500 words.

						Use the following guidance to build your article, but never allude to the guidance itself in your text.

						%s. Keep this theme consistent throughout the whole article, making sure every section adds interest and meaning to the whole. Never directly mention the genere or subgenre of the story in your text.

						Stylistically, use a varied vocabulary without sounding grandiloquent. Avoid the word "enigmatic", using a more exciting synonym instead. Also, keep you use of adverbs to a minimum, using strong and expressive verbs instead. Finally, avoid clichés.
						
						Dramatically, make sure your article is engaging and enticing. Your article should have superb pacing and keep the readers interested. Balance your scholarly explanation as an art historian with some exciting storytelling.

						Structurally, the article should look like this:

						[Title: the catchy and seductive title of the article. Keep it shorter than five words.]

						[Museum tag: information about the piece formatted as
						| [Artist] | [Title (Year)] | [Medium] |
						%s
						The artwork's year should reflect the art movement it belongs to, which you should deduce by the picture.

						If you consider the piece to be an archeological find, you can re-write the tag to reflect this.
						]

						From this point, format your text in markdown.

						[Introduction: introduce the piece and its author in two paragraphs, highlighting its uniqueness and relevancy.]

						[Artwork's legend: %s This story should be around four paragraphs long, and have a beginning, middle, and end.]

						[Final words: bring your article together in one paragraph, explaining how it affects audiences today.]

						Between two paragraphs of your choosing, add a fictional quote about the piece in the picture by a fictional expert. Format it as a separate blockquote as follows: "Quote" -Author (profession).

						Never title each section. Make sure the text flows seamlessly, is cohesive, and maintains the same theme, tone, and narrative pace.

						Please reply with the final version of your article when you're ready.
						
						`, artistInfo, storyPrompt, objectAnecdote),
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
						
						6. Rewrite nonsensical paragraphs, and cut those who are repeating a point already made. Nevertheless, stay as loyal to the original text as possible.

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
