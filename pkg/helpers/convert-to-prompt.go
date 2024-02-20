package helpers

import (
	"fmt"
	"os"
)

func ConvertToPrompt(filepath string) (string, error) {
	contentBytes, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("error reading the prompt file")
	}

	content := string(contentBytes)

	return content, nil
}
