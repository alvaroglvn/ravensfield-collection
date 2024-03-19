package internal

type ApiConfig struct {
	Port      string
	OpenAiKey string
	GhostKey  string
	GhostURL  string
	LeoKey    string
	ClaudeKey string
	MasterKey string
}

func BuildConfig(port, openAiKey, ghostKey, ghostUrl, leoKey, claudeKey, masterKey string) ApiConfig {

	config := ApiConfig{
		Port:      port,
		OpenAiKey: openAiKey,
		GhostKey:  ghostKey,
		GhostURL:  ghostUrl,
		LeoKey:    leoKey,
		ClaudeKey: claudeKey,
		MasterKey: masterKey,
	}
	return config
}
