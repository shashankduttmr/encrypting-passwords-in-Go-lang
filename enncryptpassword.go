package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	key := []byte("mysecretkey")

	message := []byte("mypassword")

	h := hmac.New(sha256.New, key)

	h.Write(message)

	result := h.Sum(nil)

	fmt.Println(result)
}
