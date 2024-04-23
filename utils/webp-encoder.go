package utils

// import (
// 	"bytes"
// 	"fmt"
// 	"image/png"
// 	"net/http"

// 	"github.com/kolesa-team/go-webp/encoder"
// 	"github.com/kolesa-team/go-webp/webp"
// )

// func PngToWebP(url string) ([]byte, error) {
// 	// Fetch the image via HTTP GET
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("bad status: %s", resp.Status)
// 	}

// 	// Decode the PNG image
// 	img, err := png.Decode(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Encode img to WebP
// 	var buf bytes.Buffer
// 	options, err := encoder.NewLosslessEncoderOptions(encoder.PresetDefault, 9)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := webp.Encode(&buf, img, options); err != nil {
// 		return nil, err
// 	}

// 	return buf.Bytes(), nil
// }
