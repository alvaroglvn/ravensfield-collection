package utils

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

	prompt := fmt.Sprintf("museum piece photographed for art catalog: %s %s %s museum piece photographed for art catalog.", mood, style, artwork)

	return prompt
}

//dalle version
// "I NEED to test how the tool works with extremely simple prompts. DO NOT add any detail, just use it AS-IS: %s %s %s museum piece photographed for art catalog using a Sigma 85mm f/1.4 lens and studio light. Don't include any lettering, fonts, or text.", mood, style, artwork
