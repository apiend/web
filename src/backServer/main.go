package main

import (
	"fmt"
	"net/http"

	"github.com/coocood/jas"
)

type Hello struct{}

func (*Hello) Get(ctx *jas.Context) {
	// `GET /v1/hello`
	ctx.Data = "hello world"
	//response: `{"data":"hello world","error":null}`
}

type UsersId struct{}

func (*UsersId) Photo(ctx *jas.Context) {
	// `GET /users/:id/photo`
	id := ctx.Id
	ctx.Data = id
}

func main() {
	router := jas.NewRouter(new(Hello), new(UsersId))
	router.BasePath = "/v1/"
	fmt.Println(router.HandledPaths(true))

	//output: `GET /v1/hello`
	http.Handle(router.BasePath, router)
	http.ListenAndServe(":8080", nil)
}
