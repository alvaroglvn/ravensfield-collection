package internal

import (
	"database/sql"
)

type ApiConfig struct {
	Port      string
	OpenAiKey string
	GhostKey  string
	Database  *sql.DB
}

func BuildConfig(port, openAiKey, ghostKey string, database *sql.DB) ApiConfig {

	config := ApiConfig{
		Port:      ": " + port,
		OpenAiKey: openAiKey,
		GhostKey:  ghostKey,
		Database:  database,
	}
	return config
}
