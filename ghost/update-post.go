package ghost

import (
	"bytes"
	"encoding/json"
	"fmt"
	// "io"
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/internal"
)

func updatePost(postID, updatedAt, featImg, title, caption, content string, config internal.ApiConfig) error {

	fmt.Print("\nCreating payload...\n")

	//prepare updated data
	payload := GhostPost{Posts: []Post{{
		UpdatedAt: updatedAt,
		FeatImage: featImg,
		Title:     title,
		HTML:      content,
		ImgCapt:   caption,
		Status:    "draft",
		Tags:      []Tag{},
	},
	},
	}

	fmt.Print("\nPayload created\n")

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshalling error: %s", err)
	}

	fmt.Print("\nJson data ready\n")

	//prepare request

	fmt.Printf("\nPreparing request...\n")

	endpoint := config.GhostURL + "/ghost/api/admin/posts/" + postID + "?source=html"

	fmt.Printf("\nSending request to: %s\n", endpoint)

	client := &http.Client{}

	req, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to build request: %s", err)
	}

	//ghost authorization
	fmt.Print("\nAuthorizing...\n")

	ghostKey := config.GhostKey
	fmt.Print("Loaded API key")
	ghostToken, err := CreateAdminToken(ghostKey)
	if err != nil {
		return fmt.Errorf("error generating ghost authorization token: %s", err)
	}
	fmt.Print("\nGhost token ready\n")

	//set headers
	fmt.Print("\nSetting headers\n")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Ghost "+ghostToken)

	//send request
	fmt.Print("\nSending request...\n")
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request error: %s", err)
	}
	defer resp.Body.Close()

	fmt.Print("\nPost created\n")

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return fmt.Errorf("error reading response: %s", err)
	// }

	// fmt.Println(string(body))

	return nil
}
