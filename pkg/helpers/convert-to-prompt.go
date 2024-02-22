package helpers

import (
	"fmt"
	"os"
	"path/filepath"
)

func ConvertToPrompt(textFile string) (string, error) {
	contentBytes, err := os.ReadFile(filepath.Clean(textFile))
	if err != nil {
		return "", fmt.Errorf("error reading the prompt file")
	}

	content := string(contentBytes)

	return content, nil
}
