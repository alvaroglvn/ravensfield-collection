package ghost

type GhostPost struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	Title     string `json:"title"`
	UpdatedAt string `json:"updated_at"`
	HTML      string `json:"html"`
	FeatImage string `json:"feature_image"`
	ImgCapt   string `json:"feature_image_caption"`
	Status    string `json:"status"`
	Tags      []Tag  `json:"tags"`
}

type Tag struct {
	Name string `json:"name"`
}
