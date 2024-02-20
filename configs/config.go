package configs

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type ApiConfig struct {
	Port string
}

func BuildConfig() ApiConfig {

	db, err := sql.Open("sqlite3", "/sqlite/db/art-museum.sqlite")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	config := ApiConfig{
		Port: ":" + os.Getenv("PORT"),
	}
	return config
}
