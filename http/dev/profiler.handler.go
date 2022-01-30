package dev

import (
	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/model/http/common"
	"github.com/valyala/fasthttp"
	nethttp "net/http"
)

type ProfilerHandler interface {
	Routes() punqy.RouteList
}

type profilerHandler struct {
	templating punqy.TemplatingEngine
	manager    punqy.ProfilerManager
}

func NewProfilerHandler(manager punqy.ProfilerManager, templating punqy.TemplatingEngine) ProfilerHandler {
	return &profilerHandler{
		templating: templating,
		manager:    manager,
	}
}

func (h *profilerHandler) Routes() punqy.RouteList {
	return punqy.RouteList{
		punqy.Route{Path: "/debug-charts", Method: punqy.Get, Handler: h.debugCharts},
		punqy.Route{Path: "/show/{id}", Method: punqy.Get, Handler: h.get},
		punqy.Route{Path: "/last", Method: punqy.Get, Handler: h.last},
		punqy.Route{Path: "/", Method: punqy.Get, Handler: h.index},
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
	var page []punqy.Profile
	list, err := h.manager.List()
	if err != nil {
		return punqy.NewResponse([]byte(err.Error()), err, nethttp.StatusOK)
	}
	pagination := common.PaginationFromReq(req, 10)
	total := len(list)
	ol := pagination.ToStorage()
	rBorder := int(ol.Offset + ol.Limit)
	lBorder := ol.Offset
	if rBorder > len(list) {
		rBorder = len(list)
	}
	for _, profile := range list[lBorder:rBorder] {
		page = append(page, profile)
	}
	htmlPagination := common.NewHtmlPagination(pagination, page, total, "/dev/profiler/")
	html, err := h.templating.Render("dev/profiler/list.gohtml", punqy.Vars{"pagination": htmlPagination, "req": req})
	if err != nil {
		return punqy.NewErrorHtmlResponse(err, nethttp.StatusInternalServerError)
	}
	return punqy.NewHtmlResponse(html.Bytes(), fasthttp.StatusOK)
}

func (h *profilerHandler) get(r punqy.Request) punqy.Response {
	last, err := h.manager.Get(r.UserValue("id").(string))
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
