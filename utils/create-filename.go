package utils

import (
	"fmt"
	"strings"
)

func CreateFileName(url string) (fileName string, err error) {
	splitUrl := strings.Split(url, "/")
	fileName = fmt.Sprintf("%s-%s", splitUrl[9], splitUrl[10][0:8])

	return fileName, nil
}
