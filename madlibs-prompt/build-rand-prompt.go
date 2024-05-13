package madlibsprompt

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func getRandFromTxt(itemList []string) (string, error) {

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

	mood, err := getRandFromTxt(genreMood)
	if err != nil {
		return "", fmt.Errorf("error getting random mood: %s", err)
	}

	artStyle, err := getRandFromTxt(artStyles)
	if err != nil {
		return "", fmt.Errorf("error getting random art style: %s", err)
	}

	object, err := getRandFromTxt(artObjects)
	if err != nil {
		return "", fmt.Errorf("error getting random object: %s", err)
	}

	prompt := fmt.Sprintf("museum piece showcase: %s %s %s. The artwork must be by itself an centered in plain color background that opposes the object's color palette.", mood, artStyle, object)

	fmt.Println(prompt)

	return prompt, nil
}
