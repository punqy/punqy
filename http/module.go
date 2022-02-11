package http

import (
	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/app/config"
	v1 "github.com/punqy/punqy/http/api/v1"
	"github.com/punqy/punqy/http/dev"
)

type ModuleHttpHandlers interface {
	ProfilerHandler() dev.ProfilerHandler
	ApiV1() v1.ModuleApiV1
}

type module struct {
	profilerHandler dev.ProfilerHandler
	apiV1           v1.ModuleApiV1
}

func (m *module) ApiV1() v1.ModuleApiV1 {
	return m.apiV1
}

func (m *module) ProfilerHandler() dev.ProfilerHandler {
	return m.profilerHandler
}

func NewModule(profiler punqy.ProfilerManager, templating punqy.TemplatingEngine, apiV1 v1.ModuleApiV1, config config.ModuleConfig) ModuleHttpHandlers {
	var m module
	m.profilerHandler = dev.NewProfilerHandler(profiler, templating, config.Config())
	m.apiV1 = apiV1
	return &m
}
