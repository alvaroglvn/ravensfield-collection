package ghost

import (
	"bytes"
	"encoding/json"
	"fmt"
	//"io"
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/internal"
)

func UpdatePost(postID, updatedAt, featImg, title, content, capt string, config internal.ApiConfig) error {

	//prepare updated data
	payload := GhostPost{Posts: []Post{{
		UpdatedAt: updatedAt,
		FeatImage: featImg,
		Title:     title,
		HTML:      content,
		ImgCapt:   capt,
		Status:    "draft",
		Tags:      []Tag{},
	},
	},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshalling error: %s", err)
	}

	//prepare request

	endpoint := config.GhostURL + "/ghost/api/admin/posts/" + postID + "?source=html"

	client := &http.Client{}

	req, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to build request: %s", err)
	}

	// fmt.Println(string(jsonData))

	//ghost authorization
	ghostKey := config.GhostKey
	ghostToken, err := CreateAdminToken(ghostKey)
	if err != nil {
		return fmt.Errorf("error generating ghost authorization token: %s", err)
	}

	//set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Ghost "+ghostToken)

	//send request
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request error: %s", err)
	}
	defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return fmt.Errorf("error reading response: %s", err)
	// }

	//fmt.Println(string(body))

	return nil
}
