package ghost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/internal"
)

func SaveDraft(article GhostPost, config internal.ApiConfig) error {
	//ghost authorization
	ghostKey := config.GhostKey
	ghostToken, err := CreateAdminToken(ghostKey)
	if err != nil {
		return fmt.Errorf("error generating ghost authorization token: %s", err)
	}

	//set endpoint
	postEndpoint := config.GhostURL + "/ghost/api/admin/posts/?source=html"

	//marshal post data
	marshData, err := json.Marshal(article)
	if err != nil {
		return fmt.Errorf("marshalling error: %s", err)
	}

	//make request
	client := &http.Client{}
	req, err := http.NewRequest("POST", postEndpoint, bytes.NewBuffer(marshData))
	if err != nil {
		return fmt.Errorf("error building request: %s", err)
	}

	//set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Ghost "+ghostToken)

	//build response
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error building response: %s", err)
	}
	defer resp.Body.Close()

	return nil
}
