package main

import (
	"fmt"
	"net/http"

	"backServer/model"
	"backServer/random"

	"backServer/lib/coocood/jas"
)

type Hello struct {
	title string
}
type remsg struct {
	name string
}

func (*Hello) Get(ctx *jas.Context) {
	// `GET /v1/hello`
	ctx.Data = random.RandString(12)
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
