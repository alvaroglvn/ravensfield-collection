package leonardo

import (
	"fmt"

	"github.com/alvaroglvn/ravensfield-collection/utils"
)

type LeonardoReq struct {
	Prompt         string  `json:"prompt"`
	NegativePr     string  `json:"negative_prompt"`
	GuidanceScale  int     `json:"guidance_scale"`
	Width          int     `json:"width"`
	Height         int     `json:"height"`
	NofImages      int     `json:"num_images"`
	Steps          int     `json:"num_inference_steps"`
	Public         bool    `json:"public"`
	Alchemy        bool    `json:"alchemy"`
	PhotoReal      bool    `json:"photoReal"`
	PhRealStrength float32 `json:"photoRealStrength"`
	PromptMagic    bool    `json:"promptMagic"`
	PresetStyle    string  `json:"presetStyle"`
	SDVersion      string  `json:"sd_version"`
}

func GetLeonardoImg(prompt, leoKey string) error {

	leoEndpoint := "https://cloud.leonardo.ai/api/rest/v1/generations"

	width, height := utils.GetRandomSize()
	fmt.Println(width)
	fmt.Println(height)

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

	respBody, err := utils.ExternalAIConnect(leoReq, leoEndpoint, leoKey)
	if err != nil {
		return fmt.Errorf("error connecting to Leonardo's API: %v", err)
	}

	fmt.Println(string(respBody))

	return nil
}
