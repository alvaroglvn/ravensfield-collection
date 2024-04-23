package pipelines

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"

// 	"github.com/alvaroglvn/ravensfield-collection/ghost"
// 	"github.com/alvaroglvn/ravensfield-collection/internal"
// )

// func CloudinaryToGhost(config internal.ApiConfig) error {

// 	_, imgUrl, err := GetNextImage(config)
// 	if err != nil {
// 		return err
// 	}

// 	//fetch image from url
// 	resp, err := http.Get(imgUrl)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	fmt.Println(imgUrl)

// 	file, err := os.Create("ghost/test.webp")
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	_, err = io.Copy(file, resp.Body)
// 	if err != nil {
// 		return err
// 	}

// 	//get image as data
// 	imgData, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return err
// 	}

// //upload image to ghost
// err = ghost.UploadImgToGhost(imgData, config)
// if err != nil {
// 	return err
// }

//build multipart dataform

// dataform, contentType, err := ghost.ImgToMPDataform(imgData, config)
// if err != nil {
// 	return err
// }

// //upload to ghost
// err = ghost.UploadImgToGhost(dataform, contentType, config)
// if err != nil {
// 	return err
// }

// 	return nil
// }
