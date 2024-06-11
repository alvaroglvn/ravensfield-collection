package ghost

import (
	"fmt"
	"strings"

	"github.com/alvaroglvn/ravensfield-collection/internal"
	"github.com/alvaroglvn/ravensfield-collection/openai"
)

func GenTextChatgpt(config internal.ApiConfig) error {
	// Get next empty article on queue (oldest)
	postId, updatedAt, featImg, err := GetOldestPostID(config)
	if err != nil {
		return fmt.Errorf("failed to load article: %s", err)
	}

	// Generate text based on feature image
	genText, err := openai.GetTextFromImg(featImg, config.OpenAiKey)
	if err != nil {
		return fmt.Errorf("failed to generate text elements: %s", err)
	}

	// Edit Text
	caption, title, content, err := openai.FinalEdit(genText, config.OpenAiKey)
	if err != nil {
		return err
	}

	// Avoid markdown tags in title
	if title[0:1] == "#" {
		title = strings.Trim(title, "#")
	}

	// Capture author's voice
	sample1, sample2, sample3, err := GetOldestArticles(config)
	if err != nil {
		return err
	}

	finalStory, err := openai.CaptureVoice(sample1, sample2, sample3, content, config.OpenAiKey)
	if err != nil {
		return err
	}

	// fmt.Printf("%s\n%s\n%s", caption, title, content)

	//update post with generated text
	err = updatePost(postId, updatedAt, featImg, title, caption, finalStory, config)
	if err != nil {
		return fmt.Errorf("failed to update post: %s", err)
	}

	return nil
}
