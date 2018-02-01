package main

import (
	"fmt"
	"net/http"

	"github.com/kn9ts/frodo"
)

func main() {
	app := frodo.New()

	app.Get("/", one, two, three)
	app.Get("/hello/:name", one, nameFunction)

	app.Serve()

}

func one(w http.ResponseWriter, r *frodo.Request) {

	fmt.Println("Hello, am the 1st middleware!")
	// fmt.Fprint(w, "Hello, I'm 1st!\n")
	r.Next()
}

func two(w http.ResponseWriter, r *frodo.Request) {
	fmt.Println("Hello, am function no. 2!")
	// fmt.Fprint(w, "Hello, am function no. 2!\n")
	r.Next()
}

func three(w http.ResponseWriter, r *frodo.Request) {
	fmt.Println("Hello, am function no 3!")
	fmt.Fprint(w, "Hey, am function no. 3!\n")

}

func nameFunction(w http.ResponseWriter, r *frodo.Request) {
	fmt.Println("Hello there, ", r.GetParam("name"))
	fmt.Fprintf(w, "Hello there, %s!\n", r.GetParam("name"))
}
