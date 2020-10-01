package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/wascript3r/gocipher/sha256"
)

func main() {
	bs, err := sha256.ComputeHMAC([]byte("plain text"), []byte("secret"))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("hashed:", hex.EncodeToString(bs))
}
