package main

import (
	"github.com/valyala/fasthttp"
	"mango/api"
	"time"
)

var appName = "api"

func main() {
	server := &fasthttp.Server{
		// see https://godoc.org/github.com/valyala/fasthttp#Server
		//Handler: fasthttp.RequestHandler(Hand),
		Handler: api.GetRouter().Handler,
		Name:    appName,
		// The maximum number of concurrent connections the server may serve.
		Concurrency:          100000,
		DisableKeepalive:     true,
		ReadTimeout:          60 * time.Second,
		WriteTimeout:         60 * time.Second,
		MaxConnsPerIP:        100000,
		MaxRequestsPerConn:   100000,
		MaxKeepaliveDuration: 120 * time.Second,
		MaxRequestBodySize:   512 * 1024 * 1024,
	}

	server.ListenAndServe("0.0.0.0:9999")
}
