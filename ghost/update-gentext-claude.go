package ghost

import (
	"fmt"

	"github.com/alvaroglvn/ravensfield-collection/claude"
	"github.com/alvaroglvn/ravensfield-collection/internal"
)

func GenTextClaude(config internal.ApiConfig) error {
	// Get oldest article on queue
	fmt.Print("\nLoading post to be edited...\n")
	postId, updatedAt, featImg, err := GetOldestPostID(config)
	if err != nil {
		return fmt.Errorf("failed to load article: %s", err)
	}
	fmt.Printf("\nArticle loaded\n")

	// Generate text based on feature image
	fmt.Print("\nGenerating text\n")
	title, caption, content, err := claude.ClaudeTextElements(featImg, config)
	if err != nil {
		return fmt.Errorf("failed to generate text elements: %s", err)
	}

	// Edit content to mimic author's voice
	fmt.Print("\nCapturing author's voice...\n")
	sample1, sample2, sample3, err := GetOldestArticles(config)
	if err != nil {
		return err
	}

	tunedText, err := claude.ClaudeAuthorVoice(sample1, sample2, sample3, content, config.ClaudeKey)
	if err != nil {
		return err
	}

	// Autoedit
	fmt.Print("\nAuto-editing...\n")

	editedText, err := claude.ClaudeAutoEdit(tunedText, config.ClaudeKey)
	if err != nil {
		return err
	}

	// Update post with generated text
	fmt.Print("\nUpdating post...\n")
	err = updatePost(postId, updatedAt, featImg, title, caption, editedText, config)
	if err != nil {
		return fmt.Errorf("failed to update post: %s", err)
	}
	fmt.Print("\nNew post ready!\n")

	return nil
}
