package pipelines

import (
	//"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/alvaroglvn/ravensfield-collection/ghost"
	"github.com/alvaroglvn/ravensfield-collection/internal"
)

func CloudinaryToGhost(config internal.ApiConfig) (err error) {
	//fetch image from Cloudinary
	_, imgUrl, err := GetNextImage(config)
	if err != nil {
		return err
	}
	resp, err := http.Get(imgUrl) //#nosec G107
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fileContents, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	//create file name
	splitUrl := strings.Split(imgUrl, "/")
	fileName := splitUrl[9] + "-" + splitUrl[10][0:8]

	//upload image to Ghost
	ghostImgUrl, err := ghost.UploadImage(config, fileContents, fileName)
	if err != nil {
		return err
	}

	//create empty draft with featured image
	postData := CreateArticle(ghostImgUrl, fileName, "", "", config)

	_, err = UploadArticletoGhost(postData, config)

	return nil
}
