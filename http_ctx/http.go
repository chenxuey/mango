package http_ctx

import (
	"github.com/valyala/fasthttp"
	"mango/config"
)

// http 请求载体
type HttpCtx struct {
	*fasthttp.RequestCtx
	*config.Logger
	Session config.Session
}

func New(ctx *fasthttp.RequestCtx, logger *config.Logger, session config.Session) *HttpCtx {
	httpCtx := HttpCtx{}
	httpCtx.RequestCtx = ctx
	httpCtx.Logger = logger
	httpCtx.Session = session
	return &httpCtx
}

