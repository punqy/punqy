package http

import (
	punqy "github.com/punqy/core"
)

const UserValueUser = "user"

type UserValuesMiddleware interface {
	Handle(req punqy.Request, next punqy.Handler) punqy.Response
}
type userValuesMiddleware struct {
}

func NewUserValuesMiddleware() UserValuesMiddleware {
	return &userValuesMiddleware{}
}

func (u *userValuesMiddleware) Handle(req punqy.Request, next punqy.Handler) punqy.Response {
	req.SetUserValue(UserValueUser, req.UserValue(punqy.SecurityContextKey).(punqy.SecurityContext).Token.User())
	return next(req)
}



