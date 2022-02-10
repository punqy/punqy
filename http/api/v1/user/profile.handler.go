package user

import (
	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/model/http/user"
	"github.com/punqy/punqy/model/storage"
	"github.com/punqy/punqy/service/app/http"
	userservice "github.com/punqy/punqy/service/user"
	nethttp "net/http"
)

type ProfileHandler interface {
	Routes() punqy.RouteList
}

type profileHandler struct {
	profileManager userservice.ProfileManager
}

func NewProfileHandler(
	profileManager userservice.ProfileManager,
) ProfileHandler {
	return &profileHandler{
		profileManager: profileManager,
	}
}

func (h *profileHandler) Routes() punqy.RouteList {
	return punqy.RouteList{
		punqy.Route{Path: "/profile", Method: punqy.Get, Handler: h.get},
		punqy.Route{Path: "/profile", Method: punqy.Patch, Handler: h.patch},
		punqy.Route{Path: "/profile/change-pass", Method: punqy.Post, Handler: h.changePass},
	}
}

func (h *profileHandler) get(req punqy.Request) punqy.Response {
	usr := req.UserValue(punqy.SecurityContextKey).(punqy.SecurityContext).Token.User()
	return punqy.NewJsonResponse(usr, nethttp.StatusOK, nil)
}

func (h *profileHandler) patch(req punqy.Request) punqy.Response {
	var request user.PatchProfileRequest
	if err := req.ParseForm(&request); err != nil {
		return punqy.NewErrorJsonResponse(err)
	}
	if err := request.Validate(); err != nil {
		return punqy.NewValidationErrJsonResponse(err)
	}
	if err := h.profileManager.PatchProfile(req, request, req.UserValue(http.UserValueUser).(storage.User)); err != nil {
		return punqy.NewErrorJsonResponse(err)
	}
	return punqy.NewOKJsonResponse()
}

func (h *profileHandler) changePass(req punqy.Request) punqy.Response {
	var request user.ChangePasswordRequest
	if err := req.ParseForm(&request); err != nil {
		return punqy.NewErrorJsonResponse(err)
	}
	if err := request.Validate(); err != nil {
		return punqy.NewValidationErrJsonResponse(err)
	}
	if err := h.profileManager.ChangePassword(req, request, req.UserValue(http.UserValueUser).(storage.User)); err != nil {
		return punqy.NewErrorJsonResponse(err)
	}
	return punqy.NewOKJsonResponse()
}
