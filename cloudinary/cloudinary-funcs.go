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

func cloudCredentials() (*cloudinary.Cloudinary, context.Context) {
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

func GetImgsFromCloud() (*admin.AssetsResult, error) {
	cld, ctx := cloudCredentials()

	resp, err := cld.Admin.Assets(ctx, admin.AssetsParams{
		DeliveryType: "upload",
		Prefix:       "ravensfield-objects",
	})
	if err != nil {
		return resp, fmt.Errorf("error loading images: %s", err)
	}

	resp, err = cld.Admin.AssetsByTag(ctx, admin.AssetsByTagParams{
		Tag: "new",
	})
	if err != nil {
		return resp, fmt.Errorf("error loading images by tag: %s", err)
	}

	// resp.Assets = sortList(resp.Assets)

	return resp, nil
}

func GetNextImgUrl(resp *admin.AssetsResult) (string, error) {
	cld, _ := cloudCredentials()
	img, err := cld.Image(resp.Assets[len(resp.Assets)-1].PublicID)
	if err != nil {
		return "", fmt.Errorf("error loading image: %s", err)
	}
	//cloudinary auto transformation
	img.Transformation = "f_webp/q_auto"
	//generate url
	imgUrl, err := img.String()
	if err != nil {
		return "", fmt.Errorf("error getting image url: %s", err)
	}

	return imgUrl, nil
}

func DeleteUsedImg(imgId string, resp *admin.AssetsResult) error {

	cld, ctx := cloudCredentials()
	_, err := cld.Admin.DeleteAssets(ctx, admin.DeleteAssetsParams{
		PublicIDs: api.CldAPIArray{imgId},
	})
	if err != nil {
		return fmt.Errorf("error deleting image: %s", err)
	}
	return nil
}

func UploadImgFromUrl(url string) error {
	cld, ctx := cloudCredentials()

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
	cld, ctx := cloudCredentials()

	_, err := cld.Upload.RemoveTag(ctx, uploader.RemoveTagParams{
		PublicIDs: []string{imgId},
		Tag:       "new",
	})
	if err != nil {
		return fmt.Errorf("error untagging image: %s", err)
	}
	return nil
}
