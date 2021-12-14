package http

import (
	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/http/oauth"
)

type ModuleHttpHandlers interface {
	OAuthTokenHandler() oauth.TokenHandler
}

type module struct {
	oauthTokenHandler oauth.TokenHandler
}

func (m *module) OAuthTokenHandler() oauth.TokenHandler {
	return m.oauthTokenHandler
}

func NewModule(auth punqy.OAuth) ModuleHttpHandlers {
	var m module
	m.oauthTokenHandler = oauth.NewTokenHandler(auth)
	return &m
}
