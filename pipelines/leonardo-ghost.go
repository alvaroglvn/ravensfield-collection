package pipelines

// import (
// 	"io"
// 	"os"

// 	"github.com/alvaroglvn/ravensfield-collection/ghost"
// 	"github.com/alvaroglvn/ravensfield-collection/internal"
// )

// func LocalToGhost(config internal.ApiConfig) error {
// 	imgFile, err := os.Open("ghost/artwork.jpg")
// 	if err != nil {
// 		return nil
// 	}
// 	defer imgFile.Close()

// 	imgData, err := io.ReadAll(imgFile)
// 	if err != nil {
// 		return nil
// 	}

// 	err = ghost.UploadImgToGhost(imgData, config)

// 	return nil
// }
