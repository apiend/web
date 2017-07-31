package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(RandInt64(100, 3000))
	var a1 int = 2000

	b1 := int64(a1)
	fmt.Println(b1)
}

func RandInt64(min, max int64) int64 {
	maxBigInt := big.NewInt(max)
	i, _ := rand.Int(rand.Reader, maxBigInt)
	if i.Int64() < min {
		RandInt64(min, max)
	}
	return i.Int64()
}
