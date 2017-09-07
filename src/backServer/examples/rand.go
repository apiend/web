package main

import "crypto/rand"
import "fmt"

func main() {
	key := make([]byte, 64)

	_, err := rand.Read(key)
	if err != nil {
		// handle error here
	}
	fmt.Println(string(key))
}
