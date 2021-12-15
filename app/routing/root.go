package routing

import (
	punqy "github.com/punqy/core"
	httphandler "github.com/punqy/punqy/http"
)

func Root(handlers httphandler.ModuleHttpHandlers) punqy.Route {
	routes := punqy.Route{
		Path: "/",
		Inner: punqy.RouteList{
			punqy.Route{
				Path:  "/dev",
				Inner: Dev(handlers),
			},
			punqy.Route{
				Path:  "/api",
				Inner: Api(handlers),
			},
		},
	}

	return routes
}
