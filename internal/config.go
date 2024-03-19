package internal

type ApiConfig struct {
	Port      string
	OpenAiKey string
	GhostKey  string
	GhostURL  string
	LeoKey    string
	MasterKey string
}

func BuildConfig(port, openAiKey, ghostKey, ghostUrl, leoKey, masterKey string) ApiConfig {

	config := ApiConfig{
		Port:      port,
		OpenAiKey: openAiKey,
		GhostKey:  ghostKey,
		GhostURL:  ghostUrl,
		LeoKey:    leoKey,
		MasterKey: masterKey,
	}
	return config
}
