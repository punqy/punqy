package app

import (
	"context"
	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/app/config"
	httphandler "github.com/punqy/punqy/http"
	"github.com/punqy/punqy/repository"
	"github.com/punqy/punqy/service"
)

type Container struct {
	OAuth              punqy.OAuth
	HttpFirewall       punqy.Firewall
	SecurityConfig     punqy.SecurityConfig
	PasswordEncoder    punqy.PasswordEncoder
	OAuthAuthenticator punqy.OAuthAuthenticator
	ProfilerManager    punqy.Manager
	ProfilerMiddleware punqy.HttpProfilerMiddleware
	TemplatingEngine   punqy.Engine
	punqy.ModuleHttpServer
	punqy.ModuleStorage
	config.ModuleConfig
	repository.ModuleRepository
	service.ModuleService
	httphandler.ModuleHttpHandlers
}

func BuildRegistry(ctx context.Context) (*Container, error) {

	app := &Container{}
	app.ModuleConfig = config.NewModule()

	app.ModuleStorage = punqy.NewModule(
		app.ModuleConfig.Config().DatabaseDriverName,
		app.ModuleConfig.Config().DatabaseDSN)

	app.ModuleRepository = repository.NewModule(app.ModuleStorage)
	app.ModuleService = service.NewModule()
	app.PasswordEncoder = punqy.NewPasswordEncoder()

	app.OAuth = punqy.NewOAuth(
		app.ModuleRepository.ClientRepository(),
		app.ModuleRepository.AccessTokenRepository(),
		app.ModuleRepository.RefreshTokenRepository(),
		app.PasswordEncoder,
		app.ModuleConfig.Config().OauthAccessTokenTTL,
		app.ModuleConfig.Config().OauthRefreshTokenTTL)

	app.ProfilerManager = punqy.NewManager(app.ModuleConfig.Config().ProfilerDir)
	app.ProfilerMiddleware = punqy.NewProfilerMiddleware(
		app.ModuleConfig.Config().ProfilerMiddlewareEnabled,
		app.ProfilerManager)
	app.TemplatingEngine = punqy.NewEngine(app.ModuleConfig.Config().TemplateDir, TemplatingConfig())
	app.ModuleHttpHandlers = httphandler.NewModule(app.OAuth, app.ProfilerManager, app.TemplatingEngine)
	app.OAuthAuthenticator = punqy.NewOAuthAuthenticator(
		app.ModuleRepository.AccessTokenRepository(),
		app.ModuleRepository.ClientRepository(),
		app.ModuleRepository.UserRepository())
	app.SecurityConfig = SecurityConfig(app.OAuthAuthenticator)

	app.HttpFirewall = punqy.NewFirewall(
		app.ModuleConfig.Config().HttpFirewallEnabled,
		app.SecurityConfig.Firewall)

	app.ModuleHttpServer = punqy.NewHttpModule(
		app.ModuleConfig.Config().HttpServerPort,
		RouterConfig(app))

	return app, nil
}
