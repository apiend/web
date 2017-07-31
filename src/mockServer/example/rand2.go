package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := 100
	i := 0
	for i < n {
		x := rand.Intn(11)
		fmt.Println(x)
		i += 1
	}
}
