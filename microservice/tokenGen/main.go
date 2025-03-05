package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	// Generate an Ed25519 key pair
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("Error generating key pair:", err)
		return
	}

	// Convert keys to Base64 for easy storage
	privateKeyBase64 := base64.StdEncoding.EncodeToString(privateKey)
	publicKeyBase64 := base64.StdEncoding.EncodeToString(publicKey)

	fmt.Println("Private Key:", privateKeyBase64)
	fmt.Println("Public Key:", publicKeyBase64)
}
