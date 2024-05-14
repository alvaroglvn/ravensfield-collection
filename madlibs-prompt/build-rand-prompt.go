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

	prompt := fmt.Sprintf("image from a museum's art catalog showcasing a %s %s %s.", mood, artStyle, object)

	return prompt, nil
}
