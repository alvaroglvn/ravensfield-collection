package pipelines

import (
	"fmt"
	"io"
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/cloudinary"
	"github.com/alvaroglvn/ravensfield-collection/ghost"
	"github.com/alvaroglvn/ravensfield-collection/internal"
	"github.com/alvaroglvn/ravensfield-collection/utils"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
)

func CloudinaryToGhost(config internal.ApiConfig) error {
	//Get Cloudinary credentials
	cld, ctx := cloudinary.CloudCredentials()

	//Get images data
	imgIds, imgUrls, err := cloudinary.GetImgsData(cld, ctx)
	if err != nil {
		return fmt.Errorf("cloudinary error: %s", err)
	}

	//Upload images to Ghost
	for _, url := range imgUrls {
		imgResp, err := http.Get(url) // #nosec G107
		if err != nil {
			return fmt.Errorf("failed to get response from image url: %s", err)
		}
		imgContent, err := io.ReadAll(imgResp.Body)
		if err != nil {
			return fmt.Errorf("error reading image content: %s", err)
		}
		//create file name
		fileName, err := utils.CreateFileName(url)
		if err != nil {
			return fmt.Errorf("failed to create file name: %s", err)
		}
		//ghost upload
		ghostImgUrl, err := ghost.UploadImage(config, imgContent, fileName)
		if err != nil {
			return fmt.Errorf("failed to upload image to ghost: %s", err)
		}

		//create empty article with feature image
		postContent := ghost.CreateArticle(ghostImgUrl, fileName, "", "", config)
		err = ghost.SaveDraft(postContent, config)
		if err != nil {
			return fmt.Errorf("error creating new article: %s", err)
		}
	}

	//After upload, delete images in Cloudinary
	for _, id := range imgIds {
		_, err := cld.Admin.DeleteAssets(ctx, admin.DeleteAssetsParams{
			PublicIDs: api.CldAPIArray{id},
		})
		if err != nil {
			return fmt.Errorf("failed to delete image from cloud: %s", err)
		}
	}
	return nil
}
