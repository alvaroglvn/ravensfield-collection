package internal

import (
	"database/sql"
)

type ApiConfig struct {
	Port      string
	OpenAiKey string
	GhostKey  string
	GhostURL  string
	LeoKey    string
	MasterKey string
	Database  *sql.DB
}

func BuildConfig(port, openAiKey, ghostKey, ghostUrl, leoKey, masterKey string, database *sql.DB) ApiConfig {

	config := ApiConfig{
		Port:      port,
		OpenAiKey: openAiKey,
		GhostKey:  ghostKey,
		GhostURL:  ghostUrl,
		LeoKey:    leoKey,
		MasterKey: masterKey,
		Database:  database,
	}
	return config
}
