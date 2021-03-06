package routing

import (
	punqy "github.com/punqy/core"
	httphandler "github.com/punqy/punqy/http"
)

func dev(handlers httphandler.ModuleHttpHandlers) punqy.RouteList {
	return punqy.RouteList{
		punqy.Route{Path: "/profiler", Inner: handlers.ProfilerHandler().Routes()},
	}
}
