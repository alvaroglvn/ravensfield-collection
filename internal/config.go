package internal

import (
	"database/sql"
)

type ApiConfig struct {
	Port      string
	OpenAiKey string
	GhostKey  string
	LeoKey    string
	Database  *sql.DB
}

func BuildConfig(port, openAiKey, ghostKey, leoKey string, database *sql.DB) ApiConfig {

	config := ApiConfig{
		Port:      ": " + port,
		OpenAiKey: openAiKey,
		GhostKey:  ghostKey,
		LeoKey:    leoKey,
		Database:  database,
	}
	return config
}
