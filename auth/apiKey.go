package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateAPIKey() (string, error) {
	apikey := make([]byte, 32)
	_, err := rand.Read(apikey)
	if err != nil {
		return "", fmt.Errorf("Can't Generate APIKEY")
	}
	return base64.RawURLEncoding.EncodeToString(apikey), nil

}
