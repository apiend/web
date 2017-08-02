package main

import (
	"fmt"
)

type nodeType int

const (
	nodeRoot nodeType = iota
	nodeParam
	nodeNormal
	nodeCatchAll
	nodeEnd
)

func main() {

	fmt.Println(nodeRoot)
	fmt.Println(nodeParam)
	fmt.Println(nodeNormal)
	fmt.Println(nodeCatchAll)
	fmt.Println(nodeEnd)

}
