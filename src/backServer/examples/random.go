package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println(RandSeq(5))
}

func RandSeq(n int) string {
	const chars = "abcdefghijklmnopqrstuvwxyz01234567ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	i, l := 0, len(chars)
	b := make([]byte, n)
	for i < n {
		b[i] = chars[rand.Intn(l)]
		i++
	}
	return string(b)
}
