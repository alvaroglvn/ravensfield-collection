package pipelines

// import (
// 	"context"
// 	"fmt"

// 	"github.com/cloudinary/cloudinary-go/v2"

// 	"github.com/alvaroglvn/ravensfield-collection/internal"
// )

// func GetNextImage(cld *cloudinary.Cloudinary, ctx context.Context, config internal.ApiConfig) (imgId, imgUrl string, err error) {
// 	//load cloud images
// 	images, err := cloudinary.GetImgsFromCloud()
// 	if err != nil {
// 		return "", "", fmt.Errorf("error loading cloud folder: %s", err)
// 	}

// 	//if there are no images in cloud
// 	if len(images.Assets) == 0 {
// 		//generate new image and upload
// 		err = LeoToCloud(config)
// 		if err != nil {
// 			return "", "", fmt.Errorf("error generating new image: %s", err)
// 		}

// 		//reload gallery
// 		images, err = cloudinary.GetImgsFromCloud()
// 		if err != nil {
// 			return "", "", fmt.Errorf("error loading cloud folder: %s", err)
// 		}
// 	}

// 	//get image id
// 	index := len(images.Assets) - 1
// 	imgId = images.Assets[index].PublicID

// 	//get img url
// 	imgUrl, err = cloudinary.GetNextImgUrl(images)
// 	if err != nil {
// 		return "", "", fmt.Errorf("error getting image url: %s", err)
// 	}

// 	return imgId, imgUrl, nil
// }
