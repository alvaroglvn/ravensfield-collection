package openai

import (
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func Test_GetDalleImg(t *testing.T) {

	utils.LoadEnv()
	apiKey := os.Getenv("OPENAI_API_KEY")

	imgUrl, err := GetDalleImg("The fluffiest kitten", apiKey)
	if err != nil {
		t.Errorf("error creating dalle image: %v", err)
	}

	content, err := http.Get(imgUrl)
	if err != nil {
		t.Errorf("error accessing image: %v", err)
	}

	bytes, err := io.ReadAll(content.Body)
	if err != nil {
		t.Errorf("error reading image: %v", err)
	}
	defer content.Body.Close()

	mimeType := http.DetectContentType(bytes)

	if mimeType != "image/png" {
		t.Errorf("content is not an image: %v", err)
	}
}
