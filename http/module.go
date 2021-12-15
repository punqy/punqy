package http

import (
	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/http/dev"
	"github.com/punqy/punqy/http/oauth"
)

type ModuleHttpHandlers interface {
	OAuthTokenHandler() oauth.TokenHandler
	ProfilerHandler() dev.ProfilerHandler
}

type module struct {
	oauthTokenHandler oauth.TokenHandler
	profilerHandler   dev.ProfilerHandler
}

func (m *module) ProfilerHandler() dev.ProfilerHandler {
	return m.profilerHandler
}

func (m *module) OAuthTokenHandler() oauth.TokenHandler {
	return m.oauthTokenHandler
}

func NewModule(auth punqy.OAuth, profiler punqy.Manager, templating punqy.Engine) ModuleHttpHandlers {
	var m module
	m.oauthTokenHandler = oauth.NewTokenHandler(auth)
	m.profilerHandler = dev.NewProfilerHandler(profiler, templating)
	return &m
}
