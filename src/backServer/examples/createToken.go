package main

import (
	"fmt"
	"time"

	t "github.com/sh3rp/token"
)

func main() {
	store, _ := t.NewTokenStore("tmp", "tokens", time.Second*1)
	store.AddUser("test", "password")
	token, _ := store.NewToken("test", "password")
	if store.IsValid(token) {
		fmt.Println("Valid!")
	}
	time.Sleep(time.Second * 3)
	if !store.IsValid(token) {
		fmt.Println("Invalid!")
	}

}
