package dev

import (
	punqy "github.com/punqy/core"
	nethttp "net/http"
)

type ProfilerHandler interface {
	Routes() punqy.RouteList
}

type profilerHandler struct {
	templating punqy.Engine
	manager    punqy.Manager
}

func NewProfilerHandler(manager punqy.Manager, templating punqy.Engine) ProfilerHandler {
	return &profilerHandler{
		templating: templating,
		manager:    manager,
	}
}

func (h *profilerHandler) Routes() punqy.RouteList {
	return punqy.RouteList{
		punqy.Route{Path: "/debug-charts", Method: punqy.GET, Handler: h.debugCharts},
		punqy.Route{Path: "/show/:id", Method: punqy.GET, Handler: h.get},
		punqy.Route{Path: "/last", Method: punqy.GET, Handler: h.last},
		punqy.Route{Path: "/", Method: punqy.GET, Handler: h.index},
	}
}

func (h *profilerHandler) debugCharts(req punqy.Request) punqy.Response {
	html, err := h.templating.Render("dev/profiler/debug_charts.gohtml", nil)
	if err != nil {
		return punqy.NewErrorHtmlResponse(err, nethttp.StatusInternalServerError)
	}
	return punqy.NewHtmlResponse(html.Bytes(), nethttp.StatusOK)
}

func (h *profilerHandler) index(req punqy.Request) punqy.Response {
	list, err := h.manager.List()
	if err != nil {
		return punqy.NewResponse([]byte(err.Error()), err, nethttp.StatusOK)
	}
	html, err := h.templating.Render("dev/profiler/list.gohtml", punqy.Vars{"pagination": pagination, "req": req})
	if err != nil {
		return punqy.NewErrorHtmlResponse(err, nethttp.StatusInternalServerError)
	}
	return punqy.NewHtmlResponse(html.Bytes(), nethttp.StatusOK)
}

func (h *profilerHandler) get(r punqy.Request) punqy.Response {
	last, err := h.manager.Get(r.Params.ByName("id"))
	if err != nil {
		return punqy.NewResponse([]byte(err.Error()), err, nethttp.StatusOK)
	}
	content, err := h.templating.Render("dev/profiler/show.gohtml", last)
	if err != nil {
		return punqy.NewResponse([]byte(err.Error()), err, nethttp.StatusOK)
	}
	return punqy.NewHtmlResponse(content.Bytes(), nethttp.StatusOK)
}

func (h *profilerHandler) last(r punqy.Request) punqy.Response {
	last, err := h.manager.Last()
	if err != nil {
		return punqy.NewResponse([]byte(err.Error()), err, nethttp.StatusOK)
	}
	content, err := h.templating.Render("dev/profiler/show.gohtml", last)
	if err != nil {
		return punqy.NewResponse([]byte(err.Error()), err, nethttp.StatusOK)
	}
	return punqy.NewHtmlResponse(content.Bytes(), nethttp.StatusOK)
}
