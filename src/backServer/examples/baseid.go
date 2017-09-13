package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	id, _ := generateBase64ID(10)
	fmt.Println(id)

}

// Generate a random base64 url encoded string
func generateBase64ID(size int) (string, error) {
	// First create a slice of bytes
	b := make([]byte, size)
	// Read size number of bytes into b
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	// Encode our bytes as a base64 encoded string using URLEncoding
	encoded := base64.URLEncoding.EncodeToString(b)
	return encoded, nil
}
