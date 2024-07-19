package leonardo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	madlibsprompt "github.com/alvaroglvn/ravensfield-collection/madlibs-prompt"
	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func generateLeoImg(prompt, leoKey string) (imgId string, err error) {

	leoEndpoint := "https://cloud.leonardo.ai/api/rest/v1/generations"

	width, height := utils.GetRandomSize()

	leoReq := LeonardoReq{
		Alchemy:       true,
		ContrastRatio: 1,
		GuidanceScale: 7,
		Height:        height,
		HighContrast:  true,
		HighRes:       true,
		ModelId:       "5c232a9e-9061-4777-980a-ddc8e65647c6", //Leonardo Vision XL
		NegativePr:    "text, fonts, lettering, letters",
		NImages:       1,
		NSteps:        60,
		PhotoReal:     true,
		PhRealV:       "v2",
		PresetStyle:   "PHOTOGRAPHY",
		Prompt:        prompt,
		PromptMagic:   false,
		//PromptMagicV:  "v3",
		Public:    false,
		SDVersion: "SDXL_1_0",
		Width:     width,
	}

	respBody, err := utils.ExternalAIPostReq(leoReq, leoEndpoint, leoKey)
	if err != nil {
		return "", fmt.Errorf("error connecting to Leonardo's API: %s", err)
	}

	var resp LeonardoResp
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response: %s", err)
	}

	imgId = resp.Job.GenId

	return imgId, nil
}

func getImgData(imgId, leoKey string) (Generations, error) {
	var gens Generations

	endpoint := fmt.Sprintf("https://cloud.leonardo.ai/api/rest/v1/generations/%s", imgId)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return gens, fmt.Errorf("error sending GET request to Leonardo: %s", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Bearer "+leoKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return gens, fmt.Errorf("response error: %s", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return gens, fmt.Errorf("error reading response body: %s", err)
	}
	defer resp.Body.Close()

	err = json.Unmarshal(body, &gens)
	if err != nil {
		return gens, fmt.Errorf("unmarshalling error: %s", err)
	}
	return gens, nil
}

func getImgUrl(imgId, leoKey string) (imgUrl string, err error) {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	timeout := time.After(30 * time.Second)

	for {
		select {
		case <-ticker.C:
			data, err := getImgData(imgId, leoKey)
			if err != nil {
				return "", fmt.Errorf("error retrieving image data: %s", err)
			}

			status := data.GenerationsByPk.Status
			if status == "COMPLETE" {
				return data.GenerationsByPk.GeneratedImages[0].URL, nil
			} else if status == "FAILED" {
				return "", fmt.Errorf("unable to create new image: %s", err)
			}
		case <-timeout:
			return "", fmt.Errorf("timeout reached before completing image generation: %s", err)
		}
	}
}

func LeonardoPipeline(leoKey string) (imgUrl string, err error) {
	prompt, err := madlibsprompt.BuildRandPrompt()
	if err != nil {
		return "", fmt.Errorf("error creating prompt: %s", err)
	}

	imgId, err := generateLeoImg(prompt, leoKey)
	if err != nil {
		return "", fmt.Errorf("error generating image: %s", err)
	}

	imgUrl, err = getImgUrl(imgId, leoKey)
	if err != nil {
		return "", fmt.Errorf("unable to retrieve image url: %s", err)
	}

	return imgUrl, nil
}
