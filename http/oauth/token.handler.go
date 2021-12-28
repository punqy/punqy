package oauth

import (
	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/model/http/oauth"
	"github.com/punqy/punqy/service/form"
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
	var request oauth.GrantAccessTokenRequest
	if err := req.ParseForm(&request); err != nil {
		return punqy.NewErrorJsonResponse(err)
	}
	if err := form.GetError(request); err != nil {
		return form.NewValidationErrJsonResponse(err)
	}
	response, err := h.oauth.GrantAccessToken(req.Context(), request.ToPunqyRequest())
	if err != nil {
		return punqy.NewErrorJsonResponse(err)
	}
	return punqy.NewJsonResponse(oauth.GrantAccessTokenResponse(response), nethttp.StatusOK, nil)
}
