package main

import (
	"fmt"

	"github.com/vmihailenco/msgpack"
)

func main() {
	ExampleMarshal()
}

// Item ljlkjl
type Item struct {
	Foo string
}

func ExampleMarshal() {

	b, err := msgpack.Marshal(&Item{Foo: "bar"})
	if err != nil {
		panic(err)
	}

	var item Item
	err = msgpack.Unmarshal(b, &item)
	if err != nil {
		panic(err)
	}
	fmt.Println(item.Foo)
	// Output: bar
}
