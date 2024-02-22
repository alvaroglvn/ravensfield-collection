package helpers

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

var projectDirName = "ravensfield-collection"

// LoadEnv loads env vars from .env
func LoadEnv() {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Fatal("unable to load env file")
	}
}
