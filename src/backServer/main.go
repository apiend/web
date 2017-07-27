package main

import (
	"fmt"
	"net/http"

	"backServer/model"

	"github.com/coocood/jas"
)

type Hello struct {
	title string
}

func (*Hello) Get(ctx *jas.Context) {
	// `GET /v1/hello`
	ctx.Data = "hello world"
	//response: `{"data":"hello world","error":null}`
}

func main() {

	router := jas.NewRouter(new(Hello), new(model.UsersId))
	router.BasePath = "/v1/"
	fmt.Println(router.HandledPaths(true))
	//output: `GET /v1/hello`
	http.Handle(router.BasePath, router)
	http.ListenAndServe(":8080", nil)
}
