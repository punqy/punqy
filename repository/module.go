package repository

import (
	"github.com/punqy/core"
	"github.com/punqy/punqy/repository/oauth"
)

type ModuleRepository interface {
}

type module struct {
	clientRepository oauth.ClientRepository
}

func (m module) ClientRepository() oauth.ClientRepository {
	return m.clientRepository
}

func NewModule(storage core.ModuleStorage) ModuleRepository {
	var m module
	m.clientRepository = oauth.NewClientRepository(storage.Dal())
	m.clientRepository = oauth.NewAccessTokenRepository(storage.Dal())
	m.clientRepository = oauth.NewRefreshTokenRepository(storage.Dal())
	return &m
}
