package configs

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type ApiConfig struct {
	Port      string
	OpenAIKey string
	DalleURL  string
}

func BuildConfig() ApiConfig {

	db, err := sql.Open("sqlite3", "/sqlite/db/art-museum.sqlite")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	config := ApiConfig{
		Port:      ":" + os.Getenv("PORT"),
		OpenAIKey: os.Getenv("OPENAI_API_KEY"),
		DalleURL:  "https://api.openai.com/v1/images/generations",
	}
	return config
}
