package pipelines

import (
	"fmt"

	"github.com/alvaroglvn/ravensfield-collection/cloudinary"
	"github.com/alvaroglvn/ravensfield-collection/internal"
)

func GetNextImage(config internal.ApiConfig) (imgId, imgUrl string, err error) {
	//load cloud images
	images, err := cloudinary.GetImgsFromCloud()
	if err != nil {
		return "", "", fmt.Errorf("error loading cloud folder: %s", err)
	}

	//if there are no images in cloud
	if len(images.Assets) == 0 {
		//generate new image and upload
		err = LeoToCloud(config)
		if err != nil {
			return "", "", fmt.Errorf("error generating new image: %s", err)
		}
	}

	//get image id
	imgId = images.Assets[len(images.Assets)-1].PublicID

	//get img url
	imgUrl, err = cloudinary.GetNextImgUrl(images)
	if err != nil {
		return "", "", fmt.Errorf("error getting image url: %s", err)
	}

	return imgId, imgUrl, nil
}
