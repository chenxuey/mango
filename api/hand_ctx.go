package api

import (
	"mango/http_ctx"
	"mango/app/users"
	"encoding/json"
)

type RequestApi struct {
}

func NewRequestApi() *RequestApi {
	return &RequestApi{}
}

func (api *RequestApi) Get(ctx *http_ctx.HttpCtx) {

	uid := ctx.Request.Header.Peek("Unique-Id")
	u := users.NewUserHttp(ctx.Logger)
	u.Get()
	response, _ := json.Marshal(struct {
		Uid string
		Xid interface{}
	}{Uid:string(uid), Xid:"佛祖保右"})
	ctx.RequestCtx.Write(response)
}
