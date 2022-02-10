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

// @Summary      Get profile
// @Description  Get current (authentiacted) user's profile
// @Tags         Profile
// @Accept       json
// @Produce      json
// @Success      200 {object} punqy.JsonResponseFormat{payload=storage.User}
// @Failure      500 {object} punqy.JsonResponseFormat{payload=string}
// @Router       /profile [get]
func (h *profileHandler) get(req punqy.Request) punqy.Response {
	usr := req.UserValue(punqy.SecurityContextKey).(punqy.SecurityContext).Token.User()
	return punqy.NewJsonResponse(usr, nethttp.StatusOK, nil)
}

// @Summary      Patch profile
// @Description  Patch current (authentiacted) user's profile
// @Tags         Profile
// @Accept       json
// @Produce      json
// @Param        request body user.PatchProfileRequest true "Body"
// @Success      200 {object} punqy.JsonResponseFormat{payload=string}
// @Success      422 {object} punqy.JsonResponseFormat{payload=common.ValidationError} "validation error"
// @Failure      500 {object} punqy.JsonResponseFormat{payload=string}
// @Router       /profile [patch]
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

// @Summary      Change password
// @Description  Change password of current (authentiacted) user
// @Tags         Profile
// @Accept       json
// @Produce      json
// @Param        request body user.ChangePasswordRequest true "Body"
// @Success      200 {object} punqy.JsonResponseFormat{payload=string} "ok"
// @Success      422 {object} punqy.JsonResponseFormat{payload=common.ValidationError} "validation error"
// @Failure      500 {object} punqy.JsonResponseFormat{payload=string} "internal error"
// @Router       /profile/change-pass [post]
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
