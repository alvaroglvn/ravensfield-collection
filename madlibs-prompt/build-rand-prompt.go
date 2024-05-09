package madlibsprompt

import (
	"fmt"

	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func BuildRandPrompt() (string, error) {

	mood, err := utils.GetRandFromTxt("madlibs-prompt/word-lists/genre-mood.txt")
	if err != nil {
		return "", fmt.Errorf("error getting random mood: %s", err)
	}

	artStyle, err := utils.GetRandFromTxt("madlibs-prompt/word-lists/art-styles.txt")
	if err != nil {
		return "", fmt.Errorf("error getting random art style: %s", err)
	}

	object, err := utils.GetRandFromTxt("madlibs-prompt/word-lists/art-objects.txt")
	if err != nil {
		return "", fmt.Errorf("error getting random object: %s", err)
	}

	prompt := fmt.Sprintf("museum piece showcase: %s %s %s. The artwork must be by itself an centered in plain color background that opposes the object's color palette.", mood, artStyle, object)

	fmt.Println(prompt)

	return prompt, nil
}
