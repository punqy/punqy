package routing

import (
	punqy "github.com/punqy/core"
	httphandler "github.com/punqy/punqy/http"
)

func Api(handlers httphandler.ModuleHttpHandlers) punqy.RouteList {
	return punqy.RouteList{
		punqy.Route{Path: "/oauth", Inner: handlers.OAuthTokenHandler().Routes()},
	}
}
