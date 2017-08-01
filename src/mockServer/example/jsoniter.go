package main

import (
	"bufio"
	"fmt"
	"os"

	jsoniter "github.com/json-iterator/go"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
	
}

func checkerr(e error) {
	if e != nil {
		fmt.Println("error:", e)
		panic(e)
	}
}

func main() {

	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := jsoniter.Marshal(group)
	checkerr(err)

	fmt.Printf(jsoniter.Get(b, "Name").ToString())
	fmt.Println()
	os.Stdout.Write(b)

	// 写入到JSON文件中
	fo, err := os.Create("temp/color.json")
	checkerr(err)

	// make a write buffer
	w := bufio.NewWriter(fo)
	w.Write(b)
	w.Flush()
	fmt.Println()

}
