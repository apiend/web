package model

import "github.com/coocood/jas"

type UsersId struct{}

func (*UsersId) Photo(ctx *jas.Context) {
	// `GET /users/:id/photo`
	id := ctx.Id
	ctx.Data = id
}

func (*UsersId) Photolist(ctx *jas.Context) {
	id := ctx.Id
	ctx.Data = id
}
