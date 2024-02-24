package ghost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/internal"
	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func GhostPostHandler(c internal.ApiConfig) http.HandlerFunc {

	ghostKey := c.GhostKey

	ghostApiUrl := "http://localhost:8081/ghost/api/admin/posts"

	ghostToken, err := CreateAdminToken(ghostKey)
	if err != nil {
		fmt.Println(err)
	}

	handlerFunct := func(w http.ResponseWriter, r *http.Request) {

		postData := GhostPost{
			Posts: []Post{
				{
					Title:     "test",
					FeatImage: "https://upload.wikimedia.org/wikipedia/commons/thumb/b/bf/Monkey_drinking_coca-cola.jpg/640px-Monkey_drinking_coca-cola.jpg",
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

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		fmt.Println("Response Status:", resp.Status)
		fmt.Println("Response Body:", string(body))

	}
	return handlerFunct
}
