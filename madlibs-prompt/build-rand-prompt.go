package madlibsprompt

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func getRandFromList(itemList []string) (string, error) {

	//get random item from list
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(itemList))))
	if err != nil {
		return "", fmt.Errorf("error generating rand index: %s", err)
	}

	randIndex := int(nBig.Int64())
	randItem := itemList[randIndex]

	return randItem, nil
}

func BuildRandPrompt() (string, error) {

	mood, err := getRandFromList(genreMood)
	if err != nil {
		return "", fmt.Errorf("error getting random mood: %s", err)
	}

	artStyle, err := getRandFromList(artStyles)
	if err != nil {
		return "", fmt.Errorf("error getting random art style: %s", err)
	}

	object, err := getRandFromList(artObjects)
	if err != nil {
		return "", fmt.Errorf("error getting random object: %s", err)
	}

	prompt := fmt.Sprintf("A captivating and thought-provoking museum piece photographed by itself against a plain background that matches its color palette: %s %s %s.", mood, artStyle, object)

	return prompt, nil
}

func BuildRandStory() (string, error) {
	adj1, err := getRandFromList(generalMood)
	if err != nil {
		return "", err
	}

	adj2, err := getRandFromList(generalMood)
	if err != nil {
		return "", err
	}

	protagonist, err := getRandFromList(protagonists)
	if err != nil {
		return "", err
	}

	fate, err := getRandFromList(fates)
	if err != nil {
		return "", err
	}

	theme, err := getRandFromList(themes)
	if err != nil {
		return "", err
	}

	ending, err := getRandFromList(endings)
	if err != nil {
		return "", err
	}

	storyPrompt := fmt.Sprintf("The story should be about %s and %s. Its protagonist must be a %s, who %s. The story must explore %s. The ending must be a %s", adj1, adj2, protagonist, fate, theme, ending)

	return storyPrompt, nil
}
