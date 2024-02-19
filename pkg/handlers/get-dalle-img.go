package handlers

import (
	"bytes"
	"encoding/json"

	"io"
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/configs"
	"github.com/alvaroglvn/ravensfield-collection/pkg/helpers"
)

type DalleRequest struct {
	Prompt         string `json:"prompt"`
	Model          string `json:"model"`
	NumberImgs     int    `json:"n"`
	Quality        string `json:"quality"`
	ResponseFormat string `json:"response_format"`
	Size           string `json:"size"`
	Style          string `json:"style"`
}

type DalleResponse struct {
	Created int     `json:"created"`
	Data    []Image `json:"data"`
}

type Image struct {
	URL string `json:"url"`
}

func GetDalleImg(w http.ResponseWriter, r *http.Request) {

	var config = configs.BuildConfig()

	prompt := helpers.PromptBuilder()

	dalleRequest := DalleRequest{
		Prompt:         prompt,
		Model:          "dall-e-3",
		NumberImgs:     1,
		Quality:        "hd",
		ResponseFormat: "url",
		Size:           "1024x1024",
		Style:          "vivid",
	}

	requestBody, err := json.Marshal(dalleRequest)
	if err != nil {
		helpers.RespondWithError(w, 500, "Error marshalling request body")
		return
	}

	// fmt.Println(string(requestBody))

	request, err := http.NewRequest("POST", config.DalleEp, bytes.NewBuffer(requestBody))
	if err != nil {
		helpers.RespondWithError(w, 500, "Error building request")
		return
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+config.OpenAIKey)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		helpers.RespondWithError(w, 500, "Error making request")
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		helpers.RespondWithError(w, 500, "Error reading response")
		return
	}

	// fmt.Println(string(body))

	var dalleResponse DalleResponse
	err = json.Unmarshal(body, &dalleResponse)
	if err != nil {
		helpers.RespondWithError(w, 500, "Error unmarshalling response")
		return
	}

	helpers.RespondWithJson(w, 201, dalleResponse)
}
