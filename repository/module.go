package repository

import (
	"github.com/punqy/core"
	"github.com/punqy/punqy/repository/oauth"
	"github.com/punqy/punqy/repository/user"
)

type ModuleRepository interface {
	AccessTokenRepository() oauth.AccessTokenRepository
	RefreshTokenRepository() oauth.RefreshTokenRepository
	ClientRepository() oauth.ClientRepository
	UserRepository() user.Repository
}

type module struct {
	clientRepository       oauth.ClientRepository
	accessTokenRepository  oauth.AccessTokenRepository
	refreshTokenRepository oauth.RefreshTokenRepository
	userRepository         user.Repository
}

func (m *module) UserRepository() user.Repository {
	return m.userRepository
}

func (m *module) AccessTokenRepository() oauth.AccessTokenRepository {
	return m.accessTokenRepository
}

func (m *module) RefreshTokenRepository() oauth.RefreshTokenRepository {
	return m.refreshTokenRepository
}

func (m *module) ClientRepository() oauth.ClientRepository {
	return m.clientRepository
}

func NewModule(storage core.ModuleStorage) ModuleRepository {
	var m module
	m.clientRepository = oauth.NewClientRepository(storage.Dal())
	m.accessTokenRepository = oauth.NewAccessTokenRepository(storage.Dal())
	m.refreshTokenRepository = oauth.NewRefreshTokenRepository(storage.Dal())
	m.userRepository = user.NewUserRepository(storage.Dal())
	return &m
}
