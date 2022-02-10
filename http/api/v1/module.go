package v1

import (
	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/http/api/v1/oauth"
	"github.com/punqy/punqy/http/api/v1/user"
	"github.com/punqy/punqy/repository"
	"github.com/punqy/punqy/service"
)

type ModuleApiV1 interface {
	OAuthTokenHandler() oauth.TokenHandler
	ProfileHandler() user.ProfileHandler
}

type module struct {
	oauthTokenHandler oauth.TokenHandler
	profileHandler    user.ProfileHandler
}

func (m *module) ProfileHandler() user.ProfileHandler {
	return m.profileHandler
}

func (m *module) OAuthTokenHandler() oauth.TokenHandler {
	return m.oauthTokenHandler
}

func NewModule(auth punqy.OAuth, repo repository.ModuleRepository, service service.ModuleService) ModuleApiV1 {
	var m module
	m.oauthTokenHandler = oauth.NewTokenHandler(auth)
	m.profileHandler = user.NewProfileHandler(service.ProfileManager())
	return &m
}
