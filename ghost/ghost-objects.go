package ghost

type GhostPost struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	Title     string `json:"title"`
	FeatImage string `json:"feature_image"`
	Status    string `json:"status"`
}
