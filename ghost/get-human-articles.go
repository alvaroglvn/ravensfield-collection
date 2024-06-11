package ghost

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/alvaroglvn/ravensfield-collection/internal"
)

func GetOldestArticles(config internal.ApiConfig) (article1, article2, article3 string, err error) {
	// Ghost authorization
	ghostKey := config.GhostKey
	ghostToken, err := CreateAdminToken(ghostKey)
	if err != nil {
		return "", "", "", fmt.Errorf("error generating ghost authorization token: %s", err)
	}

	// Set endopoint
	getPostsEp := config.GhostURL + "/ghost/api/admin/posts/?limit=12&formats=html&order=updated_at%20asc"

	//make request
	client := &http.Client{}
	req, err := http.NewRequest("GET", getPostsEp, nil)
	if err != nil {
		return "", "", "", fmt.Errorf("failed to create request: %s", err)
	}

	//set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Ghost "+ghostToken)

	//send request
	resp, err := client.Do(req)
	if err != nil {
		return "", "", "", fmt.Errorf("request error: %s", err)
	}
	defer resp.Body.Close()

	//read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", fmt.Errorf("failed to read response: %s", err)
	}

	//parse response
	var result struct {
		Posts []struct {
			Content string `json:"html"`
		} `json:"posts"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return "", "", "", fmt.Errorf("JSON parsing error: %s", err)
	}

	// get random articles from the list
	src := rand.NewSource(time.Now().UnixNano()) // #nosec G404
	randGen := rand.New(src)                     // #nosec G404

	perm := randGen.Perm(19)

	uniqueNumbers := perm[:3]

	article1 = result.Posts[uniqueNumbers[0]].Content
	article2 = result.Posts[uniqueNumbers[1]].Content
	article3 = result.Posts[uniqueNumbers[2]].Content

	return article1, article2, article3, nil
}
