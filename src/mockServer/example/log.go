package main

import "log"

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("000")
}
