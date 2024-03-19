package claude

type claudeMessage struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	System      string    `json:"system"`
	MaxTokens   int       `json:"max_tokens"`
	Temperature float64   `json:"temperature"`
}

type Message struct {
	Role    string        `json:"role"`
	Content []interface{} `json:"content"`
}

type ContentText struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type ContentImage struct {
	Type   string `json:"type"`
	Source Source `json:"source"`
}

type Source struct {
	Type      string `json:"type"`
	MediaType string `json:"media_type"`
	Data      string `json:"data"`
}

type ClaudeResponse struct {
	Content []RespContent `json:"content"`
}

type RespContent struct {
	Text string `json:"text"`
}
