package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/wascript3r/gocipher/aes"
)

func main() {
	cipher, err := aes.NewCipher("0123456789012345")
	if err != nil {
		log.Fatalln(err)
	}

	encrypted, err := cipher.Encrypt([]byte("plain text"))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("encrypted:", base64.StdEncoding.EncodeToString(encrypted))

	decrypted, err := cipher.Decrypt(encrypted)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("decrypted:", string(decrypted))
}
