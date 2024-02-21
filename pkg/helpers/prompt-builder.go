package helpers

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/alvaroglvn/ravensfield-collection/sqlite/museum_db"
)

func PromptBuilder() string {

	connect, err := sql.Open("sqlite3", "sqlite/db/art-museum.sqlite")
	if err != nil {
		log.Printf("error loading database: %v", err)
	}
	defer connect.Close()

	queries := museum_db.New(connect)

	ctx := context.Background()

	choice1, err := queries.SelectRandMood(ctx)
	if err != nil {
		log.Printf("error querying mood: %v", err)
	}

	mood := choice1.String

	choice2, err := queries.SelectRandStyle(ctx)
	if err != nil {
		log.Printf("error querying style: %v", err)
	}

	style := choice2.String

	choice3, err := queries.SelectRandArtwork(ctx)
	if err != nil {
		log.Printf("error querying artwork: %v", err)
	}

	artwork := choice3.String

	prompt := fmt.Sprintf("I NEED to test how the tool works with extremely simple prompts. DO NOT add any detail, just use it AS-IS. NEVER add text, fonts, or lettering: photo for museum's art catalog using Sigma 85mm f/1.4 lens showcasing a %s %s %s. The object must be centered.", mood, style, artwork)

	return prompt
}
