package leonardo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func CreateLeoImg(prompt, leoKey string) (imgId string, err error) {

	leoEndpoint := "https://cloud.leonardo.ai/api/rest/v1/generations"

	width, height := utils.GetRandomSize()

	leoReq := LeonardoReq{
		Prompt:         prompt,
		NegativePr:     "text, fonts, lettering",
		GuidanceScale:  10,
		Width:          width,
		Height:         height,
		NofImages:      1,
		Steps:          30,
		Public:         false,
		Alchemy:        true,
		PhotoReal:      true,
		PhRealStrength: 0.45,
		PromptMagic:    false,
		PresetStyle:    "PHOTOGRAPHY",
		SDVersion:      "v1_5",
	}

	respBody, err := utils.ExternalAIPostReq(leoReq, leoEndpoint, leoKey)
	if err != nil {
		return "", fmt.Errorf("error connecting to Leonardo's API: %v", err)
	}

	var resp LeonardoResp
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response: %v", err)
	}

	imgId = resp.Job.GenId

	return imgId, nil
}

func GetLeoImgUrl(imgId, leoKey string) (imgUrl string, err error) {

	endpoint := fmt.Sprintf("https://cloud.leonardo.ai/api/rest/v1/generations/%s", imgId)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return "", fmt.Errorf("error sending GET request to Leonardo: %v", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Bearer "+leoKey)

	resp, err := http.DefaultClient.Do(req)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}
	defer resp.Body.Close()

	var gens Generations
	err = json.Unmarshal(body, &gens)

	status := gens.GenerationsByPk.Status
	fmt.Println(status)

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	timeout := time.After(30 * time.Second)

	for {
		select {
		case <-ticker.C:
			if status == "COMPLETE" {
				imgUrl = gens.GenerationsByPk.GeneratedImages[0].URL
				return imgUrl, nil
			} else if status == "FAILED" {
				return "", fmt.Errorf("unable to create new image: %v", err)
			}
		case <-timeout:
			return "", fmt.Errorf("timeout reached before completing image generation: %v", err)
		}
	}
}
