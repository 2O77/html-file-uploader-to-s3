package idgenerator

import (
	"crypto/rand"
	"encoding/base64"
)

func NewIDGenerator() (string, error) {
	return generateRandomID()
}

func generateRandomID() (string, error) {
	// ID'nin byte dizisi için bir tampon oluştur
	id := make([]byte, 16)

	// Rastgele verilerle tamponu doldur
	_, err := rand.Read(id)
	if err != nil {
		return "", err
	}

	// Base64 ile kodlayarak rastgele veriyi dizeye dönüştür
	encodedID := base64.URLEncoding.EncodeToString(id)

	return encodedID, nil
}
