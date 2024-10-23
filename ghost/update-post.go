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

	fmt.Print("Creating payload")

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

	fmt.Print("Payload created")

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshalling error: %s", err)
	}

	fmt.Print("Created json data")

	//prepare request

	fmt.Printf("Preparing request")

	endpoint := config.GhostURL + "/ghost/api/admin/posts/" + postID + "?source=html"

	fmt.Printf("Sending request to: %s", endpoint)

	client := &http.Client{}

	req, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to build request: %s", err)
	}

	//ghost authorization
	fmt.Print("Authorizing...")

	ghostKey := config.GhostKey
	fmt.Printf("Loaded API key: %s", ghostKey)
	ghostToken, err := CreateAdminToken(ghostKey)
	if err != nil {
		return fmt.Errorf("error generating ghost authorization token: %s", err)
	}
	fmt.Printf("Ghost token created: %s", ghostToken)

	//set headers
	fmt.Print("Setting headers")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Ghost "+ghostToken)

	//send request
	fmt.Print("Sending request")
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request error: %s", err)
	}
	defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return fmt.Errorf("error reading response: %s", err)
	// }

	// fmt.Println(string(body))

	return nil
}
