package ghost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/internal"
	"github.com/alvaroglvn/ravensfield-collection/leonardo"
	"github.com/alvaroglvn/ravensfield-collection/openai"
	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func GhostPostHandler(c internal.ApiConfig) http.HandlerFunc {

	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		//create leonardo image
		img, err := leonardo.LeonardoPipeline(c.LeoKey)
		if err != nil {
			utils.RespondWithError(w, 500, fmt.Sprintf("image error: %s", err))
		}

		//create openAi text
		tag, title, descript, err := openai.GetTextFromImg(img, c.OpenAiKey)
		if err != nil {
			utils.RespondWithError(w, 500, fmt.Sprintf("text error: %s", err))
		}

		//create Ghost post
		ghostKey := c.GhostKey
		ghostToken, err := createAdminToken(ghostKey)
		if err != nil {
			utils.RespondWithError(w, 500, fmt.Sprintf("unable to create ghost auth: %s", err))
		}
		ghostEndpoint := c.GhostURL + "/ghost/api/admin/posts/?source=html"

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

		marshPostData, err := json.Marshal(postData)
		if err != nil {
			utils.RespondWithError(w, 500, "error marshalling post data")
		}

		client := &http.Client{}
		req, err := http.NewRequest("POST", ghostEndpoint, bytes.NewBuffer(marshPostData))
		if err != nil {
			utils.RespondWithError(w, 500, fmt.Sprintf("error with ghost request: %s", err))
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Ghost "+ghostToken)

		resp, err := client.Do(req)
		if err != nil {
			utils.RespondWithError(w, 500, fmt.Sprintf("error with ghost response: %s", err))
		}
		defer resp.Body.Close()
	}
	return handlerFunc
}
