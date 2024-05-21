package utils

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

func ImgUrltoBase64(imgUrl string) (string, error) {
	resp, err := http.Get(imgUrl) //#nosec G107
	if err != nil {
		return "", fmt.Errorf("error loading image from url: %s", err)
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading image from url: %s", err)
	}

	var base64Encoding string

	base64Encoding += "data:image/webp;base64,"

	encoded := base64.StdEncoding.EncodeToString(bytes)

	return encoded, nil
}
