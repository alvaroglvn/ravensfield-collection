package ghost

import (
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAdminToken(ghostKey string) (string, error) {
	splitKey := strings.Split(ghostKey, ":")

	keyID, hexKeySecret := splitKey[0], splitKey[1]

	keySecret, err := hex.DecodeString(hexKeySecret)
	if err != nil {
		return "", fmt.Errorf("failed to decode hex key secret: %v", err)
	}

	claims := jwt.MapClaims{
		"aud": "/admin/",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["kid"] = keyID

	tokenString, err := token.SignedString([]byte(keySecret))
	if err != nil {
		return "", fmt.Errorf("error signing token: %v", err)
	}
	return tokenString, nil
}
