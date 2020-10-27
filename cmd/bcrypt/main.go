package main

import (
	"fmt"
	"log"

	"github.com/wascript3r/gocipher/bcrypt"
)

func main() {
	hash, err := bcrypt.Compute([]byte("secret"), 10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("hashed:", string(hash))

	err = bcrypt.Validate(hash, []byte("secret"))
	fmt.Println("validate:", err)
}
