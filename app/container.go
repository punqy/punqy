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
	OAuth          punqy.OAuth
	HttpFirewall   punqy.Firewall
	SecurityConfig punqy.SecurityConfig
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
	app.ModuleStorage = punqy.NewModule(app.ModuleConfig.Config().DatabaseDriverName,app.ModuleConfig.Config().DatabaseDSN)
	app.ModuleRepository = repository.NewModule()
	app.ModuleService = service.NewModule()
	app.ModuleHttpHandlers = httphandler.NewModule()
	app.ModuleHttpServer = punqy.NewHttpModule(app.ModuleConfig.Config().HttpServerPort, RouterConfig(app))
	app.OAuth = punqy.NewOAuth()
	app.SecurityConfig = SecurityConfig()
	app.HttpFirewall = punqy.NewFirewall(app.ModuleConfig.Config().HttpFirewallEnabled, SecurityConfig())

	return app, nil
}
