package ghost

import (
	"fmt"

	"github.com/alvaroglvn/ravensfield-collection/claude"
	"github.com/alvaroglvn/ravensfield-collection/internal"
)

func GenTextClaude(config internal.ApiConfig) error {
	// Get oldest article on queue
	postId, updatedAt, featImg, err := GetOldestPostID(config)
	if err != nil {
		return fmt.Errorf("failed to load article: %s", err)
	}

	// Generate text based on feature image
	title, caption, content, err := claude.ClaudeTextElements(featImg, config)
	if err != nil {
		return fmt.Errorf("failed to generate text elements: %s", err)
	}

	// Edit content to mimic author's voice
	sample1, sample2, sample3, err := GetOldestArticles(config)
	if err != nil {
		return err
	}

	tunedText, err := claude.ClaudeAuthorVoice(sample1, sample2, sample3, content, config.ClaudeKey)
	if err != nil {
		return err
	}

	// Update post with generated text
	err = updatePost(postId, updatedAt, featImg, title, caption, tunedText, config)
	if err != nil {
		return fmt.Errorf("failed to update post: %s", err)
	}

	return nil
}
