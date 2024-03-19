package openai

//DALLE

type DalleRequest struct {
	Prompt         string `json:"prompt"`
	Model          string `json:"model"`
	NumberImgs     int    `json:"n"`
	Quality        string `json:"quality"`
	ResponseFormat string `json:"response_format"`
	Size           string `json:"size"`
	Style          string `json:"style"`
}

type DalleResponse struct {
	Created int     `json:"created"`
	Data    []Image `json:"data"`
}

type Image struct {
	URL string `json:"url"`
}

//COMPLETION

type CompRequest struct {
	Model           string    `json:"model"`
	Messages        []Message `json:"messages"`
	MaxTokens       int       `json:"max_tokens"`
	FreqPenalty     float64   `json:"frequency_penalty,omitempty"`
	PresencePenalty float64   `json:"presence_penalty,omitempty"`
	Temperature     float64   `json:"temperature,omitempty"`
}

type Message struct {
	Role    string        `json:"role"`
	Content []interface{} `json:"content"`
}

type TextContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type ImageContent struct {
	Type     string   `json:"type"`
	ImageURL ImageURL `json:"image_url"`
}

type ImageURL struct {
	URL string `json:"url"`
}

type CompResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
