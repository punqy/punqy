package repository

import (
	"github.com/punqy/core"
	"github.com/punqy/punqy/repository/oauth"
	"github.com/punqy/punqy/repository/user"
)

type ModuleRepository interface {
	OauthClients() oauth.ClientRepository
	Users() user.Repository
}

type module struct {
	clientRepository oauth.ClientRepository
	userRepository   user.Repository
}

func (m *module) Users() user.Repository {
	return m.userRepository
}

func (m *module) OauthClients() oauth.ClientRepository {
	return m.clientRepository
}

func NewModule(storage core.ModuleStorage) ModuleRepository {
	var m module
	m.clientRepository = oauth.NewClientRepository(storage.Dal())
	m.userRepository = user.NewUserRepository(storage.Dal())
	return &m
}
