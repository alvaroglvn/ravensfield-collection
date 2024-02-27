package utils

import (
	"crypto/rand"
	"math/big"
)

func GetRandomSize() (width, height int) {

	aspectRatios := []string{"square", "portrait", "landscape", "wide"}

	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(aspectRatios))))
	if err != nil {
		width, height = 1024, 1024
		return
	}

	randIndex := int(nBig.Int64())
	randAR := aspectRatios[randIndex]

	switch randAR {
	case "square":
		width, height = 1024, 1024
	case "portrait":
		width, height = 820, 1024
	case "landscape":
		width, height = 1024, 820
	case "wide":
		width, height = 1024, 536
	}

	return width, height
}
