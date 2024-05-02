package ghost

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/internal"
)

func GetOldestPostID(config internal.ApiConfig) (postId, featImgUrl string, err error) {
	//ghost authorization
	ghostKey := config.GhostKey
	ghostToken, err := CreateAdminToken(ghostKey)
	if err != nil {
		return "", "", fmt.Errorf("error generating ghost authorization token: %s", err)
	}

	//set endpoint
	getPostsEp := config.GhostURL + "/ghost/api/admin/posts/?limit=1&filter=status:draft&fields=id,feature_image&order=published_at%20asc"

	//make request
	client := &http.Client{}
	req, err := http.NewRequest("GET", getPostsEp, nil)
	if err != nil {
		return "", "", fmt.Errorf("failed to create request: %s", err)
	}

	//set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Ghost "+ghostToken)

	//send request
	resp, err := client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("request error: %s", err)
	}
	defer resp.Body.Close()

	//read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("failed to read response: %s", err)
	}

	//parse response
	var result struct {
		Posts []struct {
			ID           string `json:"id"`
			FeatImageUrl string `json:"feature_image"`
		} `json:"posts"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return "", "", fmt.Errorf("JSON parsing error: %s", err)
	}

	postId = result.Posts[0].ID
	featImgUrl = result.Posts[0].FeatImageUrl
	fmt.Printf("Post ID: %s, Feature Image: %s", postId, featImgUrl)

	return postId, featImgUrl, nil
}
