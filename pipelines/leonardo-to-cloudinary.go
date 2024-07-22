package pipelines

import (
	"fmt"

	"github.com/alvaroglvn/ravensfield-collection/cloudinary"
	"github.com/alvaroglvn/ravensfield-collection/internal"
	"github.com/alvaroglvn/ravensfield-collection/leonardo"
)

func LeoToCloud(config internal.ApiConfig) error {
	//create leo image
	imgUrl, prompt, err := leonardo.LeonardoPipeline(config.LeoKey)
	if err != nil {
		return fmt.Errorf("error: %s", err)
	}

	//upload created image to cloudinary
	err = cloudinary.UploadImgFromUrl(imgUrl, prompt)
	if err != nil {
		return fmt.Errorf("error: %s", err)
	}

	return nil
}
