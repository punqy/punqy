package oauth

import (
	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/model/http/oauth"
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
		punqy.Route{Path: "/test", Method: punqy.POST, Handler: h.test},
	}
}

func (h *tokenHandler) getToken(req punqy.Request) punqy.Response {
	var request oauth.GrantAccessTokenRequest
	if err := req.ParseForm(&request); err != nil {
		return punqy.NewErrorJsonResponse(err)
	}
	if err := request.Validate(); err != nil {
		return punqy.NewValidationErrJsonResponse(err)
	}
	response, err := h.oauth.GrantAccessToken(req, request.ToPunqyRequest())
	if err != nil {
		return punqy.NewErrorJsonResponse(err)
	}
	return punqy.NewJsonResponse(oauth.GrantAccessTokenResponse(response), nethttp.StatusOK, nil)
}

func (h *tokenHandler) test(req punqy.Request) punqy.Response {
	return punqy.NewJsonResponse("ok", nethttp.StatusOK, nil)
}
