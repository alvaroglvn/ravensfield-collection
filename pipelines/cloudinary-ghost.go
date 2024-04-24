package pipelines

import (
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/ghost"
	"github.com/alvaroglvn/ravensfield-collection/internal"
)

func CloudinaryToGhost(config internal.ApiConfig) (ghostImgUrl string, err error) {

	_, imgUrl, err := GetNextImage(config)
	if err != nil {
		return "", err
	}

	//fetch image from url
	resp, err := http.Get(imgUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ghostImgUrl, err = ghost.UploadImage(config, imgUrl)
	if err != nil {
		return "", err
	}

	return ghostImgUrl, err

}
