package oauth

import (
	punqy "github.com/punqy/core"
	nethttp "net/http"
)

type TokenHandler interface {
	Routes() punqy.RouteList
}

type tokenHandler struct {
	oauth punqy.OAuth
}

func NewTokenHandler(oauth punqy.OAuth) TokenHandler {
	return &tokenHandler{
		oauth: oauth,
	}
}

func (h *tokenHandler) Routes() punqy.RouteList {
	return punqy.RouteList{
		punqy.Route{Path: "/token", Method: punqy.POST, Handler: h.getToken},
	}
}

func (h *tokenHandler) getToken(req punqy.Request) punqy.Response {
	var request punqy.GrantAccessTokenRequest
	if err := req.ParseForm(&request); err != nil {
		return punqy.NewErrorJsonResponse(err)
	}
	//if err := form.GetError(request); err != nil {
	//	return punqy.NewValidationErrJsonResponse(err)
	//}
	response, err := h.oauth.GrantAccessToken(req.Context(), request)
	if err != nil {
		return punqy.NewErrorJsonResponse(err)
	}
	return punqy.NewJsonResponse(response, nethttp.StatusOK)
}
