package routing

import (
	punqy "github.com/punqy/core"
	v1 "github.com/punqy/punqy/http/api/v1"
)

func apiV1(hdl v1.ModuleApiV1) punqy.RouteList {
	return punqy.RouteList{
		punqy.Route{Path: "/oauth", Inner: hdl.OAuthTokenHandler().Routes()},
		punqy.Route{Path: "/user", Inner: hdl.ProfileHandler().Routes()},
	}
}
