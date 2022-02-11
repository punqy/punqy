package oauth

import (
	nethttp "net/http"

	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/model/http/oauth"
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
		punqy.Route{Path: "/token", Method: punqy.Post, Handler: h.getToken},
	}
}

// @Summary      Retrieve oauth token
// @Description  Retrieve oauth token
// @Tags         OAuth
// @Accept       json
// @Produce      json
// @Param        request body oauth.GrantAccessTokenRequest true "Body"
// @Success      200 {object} punqy.JsonResponseFormat{payload=oauth.GrantAccessTokenResponse}
// @Success      422 {object} punqy.JsonResponseFormat{payload=common.ValidationError} "validation error"
// @Failure      500 {object} punqy.JsonResponseFormat{payload=string}
// @Router       /oauth/token [post]
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
