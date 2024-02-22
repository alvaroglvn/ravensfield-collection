package internal

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type ApiConfig struct {
	Port      string
	OpenAiKey string
}

func BuildConfig() ApiConfig {

	db, err := sql.Open("sqlite3", "/sqlite/db/art-museum.sqlite")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	config := ApiConfig{
		Port:      ":" + os.Getenv("PORT"),
		OpenAiKey: os.Getenv("OPENAI_API_KEY"),
	}
	return config
}
