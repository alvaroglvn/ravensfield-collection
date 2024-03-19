package utils

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func GetRandFromTxt(file string) (string, error) {

	//generate string from txt
	content, err := os.ReadFile(file)
	if err != nil {
		return "nil", fmt.Errorf("error reading file: %s", err)
	}
	list := strings.Split(string(content), "\n")

	//get random item from list
	randIndex := rand.Intn(len(list))
	randItem := list[randIndex]

	return randItem, nil
}
