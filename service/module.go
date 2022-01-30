package service

import (
	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/app/config"
	"github.com/punqy/punqy/repository"
	"github.com/punqy/punqy/service/security/oauth"
)

type ModuleService interface {
	AccessTokenStorage() punqy.OAuthAccessTokenStorage
	RefreshTokenStorage() punqy.OAuthRefreshTokenStorage
	ClientStorage() punqy.OAuthClientStorage
	UserManager() oauth.UserManager
}

type module struct {
	accessTokenStorage  punqy.OAuthAccessTokenStorage
	refreshTokenStorage punqy.OAuthRefreshTokenStorage
	userManager         oauth.UserManager
	clientStorage       punqy.OAuthClientStorage
}

func (m *module) ClientStorage() punqy.OAuthClientStorage {
	return m.clientStorage
}

func (m *module) UserManager() oauth.UserManager {
	return m.userManager
}

func (m *module) AccessTokenStorage() punqy.OAuthAccessTokenStorage {
	return m.accessTokenStorage
}

func (m *module) RefreshTokenStorage() punqy.OAuthRefreshTokenStorage {
	return m.refreshTokenStorage
}

func NewModule(repo repository.ModuleRepository, encoder punqy.PasswordEncoder, config config.ModuleConfig) ModuleService {
	var m module
	m.accessTokenStorage = oauth.NewAccessTokenStorage(config.Config().JWTAccessTokenSigningKey, config.Config().OauthAccessTokenTTL)
	m.refreshTokenStorage = oauth.NewRefreshTokenStorage(config.Config().JWTRefreshTokenSigningKey, config.Config().OauthRefreshTokenTTL)
	m.userManager = oauth.NewUserManager(repo.UserRepository(), encoder)
	m.clientStorage = oauth.NewClientStorage(repo.ClientRepository())
	return &m
}
