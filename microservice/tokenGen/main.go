package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/curve25519"
)

func main() {
	// Generate a random 32-byte secret key
	secretKey := make([]byte, 32)
	_, err := rand.Read(secretKey)
	if err != nil {
		fmt.Println("Error generating secret key:", err)
		return
	}

	// Derive the public key from the secret key
	var publicKey [32]byte
	curve25519.ScalarBaseMult(&publicKey, (*[32]byte)(secretKey))

	
	// Convert keys to Base64 for easy storage
	secretKeyBase64 := base64.StdEncoding.EncodeToString(secretKey)
	publicKeyBase64 := base64.StdEncoding.EncodeToString(publicKey[:])

	fmt.Println("Secret Key:", secretKeyBase64)
	fmt.Println("Public Key:", publicKeyBase64)
}
