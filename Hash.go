package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	text := "José Benito Ibarria Topete"
	hasher := sha256.New()
	hasher.Write([]byte(text))

	hash := hasher.Sum(nil)

	fmt.Printf("The hash of the text is: %x/n", hash)
}
