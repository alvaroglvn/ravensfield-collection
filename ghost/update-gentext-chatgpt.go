package ghost

import (
	"fmt"
	"log"
	"strings"

	"github.com/alvaroglvn/ravensfield-collection/internal"
	"github.com/alvaroglvn/ravensfield-collection/openai"
)

func GenTextChatgpt(config internal.ApiConfig) error {

	// Get next empty article on queue (oldest)
	log.Print("Loading article...")

	postId, updatedAt, featImg, err := GetOldestPostID(config)
	if err != nil {
		return fmt.Errorf("failed to load article: %s", err)
	}

	log.Printf("Article loaded:\nPostId: %s\nUpdated At: %s\nFeat Image:%s", postId, updatedAt, featImg)

	// Generate text based on feature image
	log.Print("Generating text...")

	genText, err := openai.GetTextFromImg(featImg, config.OpenAiKey)
	if err != nil {
		return fmt.Errorf("failed to generate text elements: %s", err)
	}

	log.Printf("Text generated:%s...", genText[0:100])

	// Edit Text

	log.Print("Editing text and formatting...")

	caption, title, content, err := openai.FinalEdit(genText, config.OpenAiKey)
	if err != nil {
		return err
	}

	log.Printf("Text edited:\nCaption: %s\nTitle: %s\nContent:%s[...]", caption, title, content[0:50])

	// Avoid markdown tags in title

	log.Print("Cleaning up title format...")

	if title[0:1] == "#" {
		title = strings.Trim(title, "#")
	}

	if title[0:1] == "*" {
		title = strings.Trim(title, "*")
	}

	log.Printf("Title formatted: %s", title)

	// Capture author's voice

	log.Print("Adding author's voice...")
	sample1, sample2, sample3, err := GetOldestArticles(config)
	if err != nil {
		return err
	}

	finalStory, err := openai.CaptureVoice(sample1, sample2, sample3, content, config.OpenAiKey)
	if err != nil {
		return err
	}

	log.Printf("Final text ready: %s...", finalStory[0:50])

	//update post with generated text

	log.Print("Updating post and saving draft...")

	err = updatePost(postId, updatedAt, featImg, title, caption, finalStory, config)
	if err != nil {
		return fmt.Errorf("failed to update post: %s", err)
	}

	log.Print("Post created")

	return nil
}
