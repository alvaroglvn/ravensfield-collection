package ghost

import (
	"time"

	"github.com/alvaroglvn/ravensfield-collection/internal"
)

func CreateArticle(img, title, tag, descript string, config internal.ApiConfig) GhostPost {

	postData := GhostPost{
		Posts: []Post{
			{
				Title:     title,
				UpdatedAt: time.Now().UTC().Format(time.RFC3339),
				HTML:      descript,
				FeatImage: img,
				ImgCapt:   tag,
				Status:    "draft",
				Tags: []Tag{
					{Name: "no text"},
				},
			},
		},
	}

	return postData
}
