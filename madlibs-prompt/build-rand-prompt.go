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

	prompt := fmt.Sprintf("A captivating and thought-provoking museum piece photographed by itself centered on a plain color background that matches its color palette: %s %s %s.", mood, artStyle, object)

	return prompt, nil
}

func BuildRandStory() (string, error) {

	subgenre, err := getRandFromList(subgenres)
	if err != nil {
		return "", err
	}

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

	storyPrompt := fmt.Sprintf("This object carries a %s story. The story's theme is %s and it explores %s and %s. It features a %s who interacts with the artwork and ultimately %s. The story's surprising and inevitable ending is a %s", subgenre, theme, adj1, adj2, protagonist, fate, ending)

	return storyPrompt, nil
}

func GetArtistInfo() (string, error) {

	artistTypes := []string{"man", "woman", "collective", "unknown"}
	artistAges := []string{"young", "adult", "middle-aged", "mature"}
	artistOrigins := []string{"Europe", "North America", "South America", "Africa", "Asia", "Oceania", "Another dimension"}

	artistType, err := getRandFromList(artistTypes)
	if err != nil {
		return "", err
	}
	artistAge, err := getRandFromList(artistAges)
	if err != nil {
		return "", err
	}

	artistOrigin, err := getRandFromList(artistOrigins)
	if err != nil {
		return "", err
	}

	artistInfo := ""

	if artistType == "man" || artistType == "woman" {
		artistInfo = fmt.Sprintf("The artist is an imaginary %s %s originally from %s", artistAge, artistType, artistOrigin)
	} else if artistType == "collective" {
		artistInfo = fmt.Sprintf("This piece is attributed to a %s", artistType)
	} else {
		artistInfo = fmt.Sprintf("The author of this piece is %s", artistType)
	}

	return artistInfo, nil
}
