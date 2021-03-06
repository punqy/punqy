package app

import (
	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/app/routing"
	"github.com/valyala/fasthttp"
)

func RouterConfig(container *Container) punqy.RouterConfig {
	return punqy.RouterConfig{
		Middlewares: punqy.MiddlewareChain{
			container.ProfilerMiddleware.Handle,
			container.HttpFirewall.Handle,
			container.UserValuesMiddleware().Handle,
		},
		StaticFiles: &punqy.StaticFiles{
			Path:    "/static/{filepath:*}",
			RootDir: "public",
		},
		Routing: routing.Routes(container.ModuleHttpHandlers),
		NotFoundHandler: func(ctx *fasthttp.RequestCtx) {
			ctx.Response.SetStatusCode(fasthttp.StatusNotFound)
			ctx.Response.ResetBody()
		},
		GlobalHandler: func(ctx *fasthttp.RequestCtx) {
			ctx.Request.Header.Set("Access-Control-Allow-Methods", string(ctx.Request.Header.Peek("Allow")))
			ctx.Request.Header.Set("Access-Control-Allow-Origin", "*")
			ctx.Response.SetStatusCode(fasthttp.StatusOK)
		},
	}
}
