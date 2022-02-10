package routing

import (
	punqy "github.com/punqy/core"
	httphandler "github.com/punqy/punqy/http"
)

func Routes(handlers httphandler.ModuleHttpHandlers) punqy.Route {
	routes := punqy.Route{
		Path: "/",
		Inner: punqy.RouteList{
			punqy.Route{
				Path:  "/dev",
				Inner: dev(handlers),
			},
			punqy.Route{
				Path:  "/api",
				Inner: punqy.RouteList{
					punqy.Route{
						Path:  "/v1",
						Inner: apiV1(handlers.ApiV1()),
					},
				},
			},
		},
	}

	return routes
}
