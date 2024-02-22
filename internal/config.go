package internal

import (
	"database/sql"
)

type ApiConfig struct {
	Port      string
	OpenAiKey string
	Database  *sql.DB
}

func BuildConfig(port, openAiKey string, database *sql.DB) ApiConfig {

	config := ApiConfig{
		Port:      ": " + port,
		OpenAiKey: openAiKey,
		Database:  database,
	}
	return config
}
