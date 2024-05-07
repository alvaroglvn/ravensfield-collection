package pipelines

import (
	"time"

	"github.com/alvaroglvn/ravensfield-collection/ghost"
	"github.com/alvaroglvn/ravensfield-collection/internal"
)

func CreateArticle(img, title, tag, descript string, config internal.ApiConfig) ghost.GhostPost {

	postData := ghost.GhostPost{
		Posts: []ghost.Post{
			{
				Title:     title,
				UpdatedAt: time.Now().UTC().Format(time.RFC3339),
				HTML:      descript,
				FeatImage: img,
				ImgCapt:   tag,
				Status:    "draft",
				Tags: []ghost.Tag{
					{Name: "no text"},
				},
			},
		},
	}

	return postData
}
