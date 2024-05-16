package ghost

import (
	"fmt"

	"github.com/alvaroglvn/ravensfield-collection/claude"
	"github.com/alvaroglvn/ravensfield-collection/internal"
)

func GenTextClaude(config internal.ApiConfig) error {
	//get next empty article on queue (oldest)
	postId, updatedAt, featImg, err := GetOldestPostID(config)
	if err != nil {
		return fmt.Errorf("failed to load article: %s", err)
	}

	// //generate text based on feature image
	title, caption, content, err := claude.ClaudeTextElements(featImg, config)
	if err != nil {
		return fmt.Errorf("failed to generate text elements: %s", err)
	}

	//update post with generated text
	err = updatePost(postId, updatedAt, featImg, title, caption, content, config)
	if err != nil {
		return fmt.Errorf("failed to update post: %s", err)
	}

	return nil
}
