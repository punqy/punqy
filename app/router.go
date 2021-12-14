package app

import (
	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/app/routing"
	nethttp "net/http"
)

func RouterConfig(container *Container) punqy.RouterConfig {
	return punqy.RouterConfig{
		Middlewares: punqy.MiddlewareChain{
		},
		StaticFiles: &punqy.StaticFiles{
			Path:    "/static/*filepath",
			RootDir: "./public",
		},
		Routing:         routing.Root(container.ModuleHttpHandlers),
		NotFoundHandler: nethttp.DefaultServeMux,
		GlobalHandler: nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
			if r.Header.Get("Access-Control-Request-Method") != "" {
				w.Header().Set("Access-Control-Allow-Methods", w.Header().Get("Allow"))
				w.Header().Set("Access-Control-Allow-Origin", "*")
			}
			w.WriteHeader(nethttp.StatusNoContent)
		}),
	}
}
