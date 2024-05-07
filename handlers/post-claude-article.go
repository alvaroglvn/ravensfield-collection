package handlers

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/alvaroglvn/ravensfield-collection/claude"
// 	"github.com/alvaroglvn/ravensfield-collection/cloudinary"
// 	"github.com/alvaroglvn/ravensfield-collection/internal"
// 	"github.com/alvaroglvn/ravensfield-collection/pipelines"
// 	"github.com/alvaroglvn/ravensfield-collection/utils"
// )

// func PostClaudeArticle(config internal.ApiConfig) http.HandlerFunc {
// 	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
// 		//get image from cloud
// 		imgId, imgUrl, err := pipelines.GetNextImage(config)
// 		if err != nil {
// 			utils.RespondWithError(w, 500, fmt.Sprintf("error getting image from cloud: %s", err))
// 			return
// 		}

// 		//get text from image
// 		title, tag, descript, err := claude.ClaudeTextElements(imgUrl, config)
// 		if err != nil {
// 			utils.RespondWithError(w, 500, fmt.Sprintf("error generating text from image: %s", err))
// 			return
// 		}

// 		//build article
// 		article := pipelines.CreateArticle(imgUrl, title, tag, descript, config)

// 		//upload article to Ghost
// 		_, err = pipelines.UploadArticletoGhost(article, config)
// 		if err != nil {
// 			utils.RespondWithError(w, 500, fmt.Sprintf("error uploading article to ghost: %s", err))
// 			return
// 		}

// 		//after upload is complete, untag image
// 		images, err := cloudinary.GetImgsFromCloud()
// 		if err != nil {
// 			utils.RespondWithJson(w, 500, fmt.Sprintf("error: %s", err))
// 		}

// 		err = cloudinary.UntagImage(imgId, images)
// 		if err != nil {
// 			utils.RespondWithError(w, 500, fmt.Sprintf("error: %s", err))
// 		}
// 	}
// 	return handlerFunc
// }
