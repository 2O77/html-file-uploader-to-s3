package idgenerator

import (
	"crypto/rand"
	"encoding/base64"
)

func NewIDGenerator() (string, error) {
	return generateRandomID()
}

func generateRandomID() (string, error) {
	id := make([]byte, 16)

	_, err := rand.Read(id)
	if err != nil {
		return "", err
	}

	encodedID := base64.URLEncoding.EncodeToString(id)

	return encodedID, nil
}
