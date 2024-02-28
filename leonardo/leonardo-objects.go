package leonardo

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

type LeonardoResp struct {
	Job struct {
		GenId string `json:"generationId"`
		Cost  int    `json:"apiCreditCost"`
	} `json:"sdGenerationJob"`
}

type Generations struct {
	GenerationsByPk struct {
		GeneratedImages []struct {
			URL string `json:"url"`
			//Nsfw         bool          `json:"nsfw"`
			ID string `json:"id"`
			//LikeCount    int           `json:"likeCount"`
			//MotionMP4URL string        `json:"motionMP4URL"`
			GenImgVarGen []ImgGenerics `json:"generated_image_variation_generics"`
		} `json:"generated_images"`
		// ModelId             string        `json:"modelId"`
		// Motion              bool          `json:"motion"`
		// MotionModel         string        `json:"motionModel"`
		// MotionStrength      int           `json:"motionStrength"`
		// Prompt              string        `json:"prompt"`
		// NegativePrompt      string        `json:"negativePrompt"`
		// ImageHeight         int           `json:"imageHeight"`
		// ImageToVideo        bool          `json:"imageToVideo"`
		// ImageWidth          int           `json:"imageWidth"`
		// InferenceSteps      int           `json:"inferenceSteps"`
		// Seed                int           `json:"seed"`
		// Public              bool          `json:"public"`
		// Scheduler           string        `json:"scheduler"`
		// SdVersion           string        `json:"sdVersion"`
		Status string `json:"status"`
		// PresetStyle         string        `json:"presetStyle"`
		// InitStrength        float64       `json:"initStrength"`
		// GuidanceScale       float64       `json:"guidanceScale"`
		// ID                  string        `json:"id"`
		// CreatedAt           string    `json:"createdAt"`
		// PromptMagic         bool          `json:"promptMagic"`
		// PromptMagicVersion  string        `json:"promptMagicVersion"`
		// PromptMagicStrength float64       `json:"promptMagicStrength"`
		// PhotoReal           bool          `json:"photoReal"`
		// PhotoRealStrength   float64       `json:"photoRealStrength"`
		// FantasyAvatar       interface{}   `json:"fantasyAvatar"`
		// GenerationElements  []interface{} `json:"generation_elements"`
	} `json:"generations_by_pk"`
}

type ImgGenerics struct {
	Id     string `json:"id"`
	Status string `json:"status"`
	//Transform string `json:"transformType"`
	Url string `json:"url"`
}
