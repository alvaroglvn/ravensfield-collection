package ghost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/internal"
	"github.com/alvaroglvn/ravensfield-collection/openai"
	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func GhostPostHandler(c internal.ApiConfig) http.HandlerFunc {

	ghostKey := c.GhostKey

	ghostApiUrl := "http://localhost:8081/ghost/api/admin/posts/?source=html"

	ghostToken, err := CreateAdminToken(ghostKey)
	if err != nil {
		fmt.Println(err)
	}

	handlerFunct := func(w http.ResponseWriter, r *http.Request) {
		img, tag, title, descript, err := openai.GetArticlePieces(c.OpenAiKey)
		if err != nil {
			fmt.Println(err)
		}

		postData := GhostPost{
			Posts: []Post{
				{
					Title:     title,
					HTML:      descript,
					FeatImage: img,
					ImgCapt:   tag,
					Status:    "draft",
				},
			},
		}

		postDataBytes, err := json.Marshal(postData)
		if err != nil {
			utils.RespondWithError(w, 500, "error marshalling post data")
		}

		client := &http.Client{}
		req, err := http.NewRequest("POST", ghostApiUrl, bytes.NewBuffer(postDataBytes))
		if err != nil {
			utils.RespondWithError(w, 500, "error connecting to ghost")
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Ghost "+ghostToken)

		resp, err := client.Do(req)
		if err != nil {
			utils.RespondWithError(w, 500, "error sending request")
		}
		defer resp.Body.Close()
	}
	return handlerFunct
}
