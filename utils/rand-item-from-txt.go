package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"path/filepath"

	"os"
	"strings"
)

func GetRandFromTxt(file string) (string, error) {

	//generate string from txt
	content, err := os.ReadFile(filepath.Clean(file))
	if err != nil {
		return "nil", fmt.Errorf("error reading file: %s", err)
	}
	list := strings.Split(string(content), "\n")

	//get random item from list
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(list))))
	if err != nil {
		return "", fmt.Errorf("error generating rand index: %s", err)
	}

	randIndex := int(nBig.Int64())
	randItem := list[randIndex]

	return randItem, nil
}
