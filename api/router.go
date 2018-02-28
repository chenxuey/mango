package api

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"mango/mango_tool"
	"mango/http_ctx"
	"mango/config"
	"mango/mango_logger"
)

type ApiHandle func(ctx *http_ctx.HttpCtx)

type ApiRoute struct {
	method []string
	handle ApiHandle
}

type Decorator struct {
	RunFuc  ApiHandle
	PathStr string
}

var dct Decorator

// 包装 request ctx
func (d Decorator) Decorator(ctx *fasthttp.RequestCtx) {
	var mangoLog = config.NewLogger(false, mango_logger.AddLog)
	uid := mango_tool.GetUID()
	mangoLog.Info = uid
	mangoLog.Println(string(ctx.Request.RequestURI()) + "\r\n" + ctx.Request.PostArgs().String())
	ctx.Response.Header.Set("x-response-id", uid)
	var httpCtx = http_ctx.New(ctx, mangoLog, nil)
	d.RunFuc(httpCtx)
	mangoLog.Println(ctx.Response.String())
	return
}

func GetRouter() *fasthttprouter.Router {
	// add api route in here.
	var r = NewRequestApi()
	ApiRouteList := map[string]ApiRoute{
		"/get": {[]string{"GET", "POST"}, r.Get},
	}

	router := &fasthttprouter.Router{
		RedirectTrailingSlash:  true,
		RedirectFixedPath:      true,
		HandleMethodNotAllowed: true,
		HandleOPTIONS:          true,
	}

	for k, v := range ApiRouteList {
		for _, value := range v.method {
			dct.PathStr = k
			dct.RunFuc = v.handle
			router.Handle(value, k, dct.Decorator)
		}
	}

	return router
}
