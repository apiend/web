package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ColorGroup struct {
	ID     int      `json:"ID"`
	Name   string   `json:"Name"`
	Colors []string `json:"Colors"`
}

func checkerr(e error) {
	if e != nil {
		fmt.Println("error:", e)
		// panic(e)
	}
}
func main() {
	// open input file
	file, _ := os.Open("temp/color.json")
	decoder := json.NewDecoder(file)
	configuration := ColorGroup{}
	err := decoder.Decode(&configuration)
	checkerr(err)
	fmt.Println(configuration.ID)
	// fmt.Println(configuration)
}
