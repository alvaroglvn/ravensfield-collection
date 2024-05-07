package cloudinary

// Import Cloudinary and other necessary libraries
//===================
import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"

	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func CloudCredentials() (*cloudinary.Cloudinary, context.Context) {
	// Add your Cloudinary credentials, set configuration parameter
	// Secure=true to return "https" URLs, and create a context
	//===================
	cld, err := cloudinary.New()
	if err != nil {
		fmt.Printf("error loading cloudinary credentials: %s", err)
	}
	cld.Config.URL.Secure = true
	ctx := context.Background()
	return cld, ctx
}

func getImgsFromCloud(cld *cloudinary.Cloudinary, ctx context.Context) (*admin.AssetsResult, error) {
	// cld, ctx := CloudCredentials()

	resp, err := cld.Admin.Assets(ctx, admin.AssetsParams{
		DeliveryType: "upload",
		Prefix:       "ravensfield-objects",
	})
	if err != nil {
		return resp, fmt.Errorf("error loading assets: %s", err)
	}

	return resp, nil
}

func GetImgsData(cld *cloudinary.Cloudinary, ctx context.Context) (imgIds, imgUrls []string, err error) {
	//1 - Load images archive
	imgArchive, err := getImgsFromCloud(cld, ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("cloudinary error: %s", err)
	}

	//2 - Loop through assets, transform into wepb, and get data
	imgIds = []string{}
	imgUrls = []string{}

	if len(imgArchive.Assets) != 0 {
		for i, asset := range imgArchive.Assets {
			//add image's id to ids list
			imgIds = append(imgIds, asset.PublicID)
			//get image
			img, err := cld.Image(imgArchive.Assets[i].PublicID)
			if err != nil {
				return nil, nil, fmt.Errorf("error loading image: %s", err)
			}
			//transform image into webp
			img.Transformation = "f_webp/q_auto"
			//get image's url
			imgUrl, err := img.String()
			if err != nil {
				return nil, nil, fmt.Errorf("error loading image's url: %s", err)
			}
			//add image's url to list
			imgUrls = append(imgUrls, imgUrl)
		}
	}

	return imgIds, imgUrls, nil
}

func DeleteImg(imgId string, resp *admin.AssetsResult) error {

	cld, ctx := CloudCredentials()
	_, err := cld.Admin.DeleteAssets(ctx, admin.DeleteAssetsParams{
		PublicIDs: api.CldAPIArray{imgId},
	})
	if err != nil {
		return fmt.Errorf("error deleting image: %s", err)
	}
	return nil
}

func UploadImgFromUrl(url string) error {
	cld, ctx := CloudCredentials()

	_, err := cld.Upload.Upload(ctx, url, uploader.UploadParams{
		Folder: "ravensfield-objects",
		Tags:   []string{"new"},
	})

	if err != nil {
		return fmt.Errorf("upload error: %s", err)
	}

	return nil
}

func UntagImage(imgId string, resp *admin.AssetsResult) error {
	cld, ctx := CloudCredentials()

	_, err := cld.Upload.RemoveTag(ctx, uploader.RemoveTagParams{
		PublicIDs: []string{imgId},
		Tag:       "new",
	})
	if err != nil {
		return fmt.Errorf("error untagging image: %s", err)
	}
	return nil
}
