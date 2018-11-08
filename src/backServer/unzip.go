package main

import(
	"fmt"
	gounzip "github.com/duncanhall/gounzip"
)


func main(){

	fmt.Printf("ss")
	err := gounzip.Unzip("demo.zip")
	if err != nil {
		return err
	}
	
}