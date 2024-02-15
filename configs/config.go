package configs

import "os"

type ApiConfig struct {
	Port      string
	OpenAIKey string
	DalleURL  string
}

func BuildConfig() ApiConfig {
	config := ApiConfig{
		Port:      ":" + os.Getenv("PORT"),
		OpenAIKey: os.Getenv("OPENAI_API_KEY"),
		DalleURL:  "https://api.openai.com/v1/images/generations",
	}
	return config
}
