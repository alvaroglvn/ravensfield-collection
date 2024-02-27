package openai

import (
	"os"
	"testing"

	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func _Test_ImgDescribe(t *testing.T) {
	utils.LoadEnv()
	apiKey := os.Getenv("OPENAI_API_KEY")

	_, err := ImgDescribe("https://upload.wikimedia.org/wikipedia/commons/thumb/3/34/MaskeAgamemnon.JPG/640px-MaskeAgamemnon.JPG", apiKey)
	if err != nil {
		t.Errorf("error creating image description: %v", err)
	}
}
