package oauth

import (
	"context"
	"github.com/google/uuid"
	punqy "github.com/punqy/core"
	model "github.com/punqy/punqy/model/storage"
	"github.com/punqy/punqy/repository/oauth"
	"github.com/punqy/punqy/repository/user"
	"time"
)

type AccessTokenStorage struct {
	accessTokenRepo oauth.AccessTokenRepository
	userRepo        user.Repository
}

func NewAccessTokenStorage(
	accessTokenRepo oauth.AccessTokenRepository,
	userRepo user.Repository,
) punqy.OAuthAccessTokenStorage {
	return &AccessTokenStorage{
		accessTokenRepo: accessTokenRepo,
		userRepo:        userRepo,
	}
}

func (a *AccessTokenStorage) GetAccessToken(ctx context.Context, token string) (punqy.OAuthAccessToken, error) {
	var tok punqy.OAuthAccessToken
	tok, err := a.accessTokenRepo.FindOneByToken(ctx, token)
	return tok, err
}

func (a *AccessTokenStorage) CreateAccessToken(ctx context.Context, user punqy.User, client punqy.OAuthClient, token string, expiresAt time.Time) error {
	var userID *uuid.UUID
	if _, ok := user.(model.User); ok {
		uid, err := uuid.Parse(user.GetID())
		if err != nil {
			return err
		}
		userID = &uid
	}
	cid, err := uuid.Parse(client.GetID())
	if err != nil {
		return err
	}

	e := model.AccessToken{
		Token:     token,
		UserId:    userID,
		ClientId:  cid,
		ExpiresAt: expiresAt,
	}
	if err := e.Init(); err != nil {
		return err
	}
	return a.accessTokenRepo.Insert(ctx, e)

}

func (a *AccessTokenStorage) CheckCredentials(ctx context.Context, token string) (punqy.User, error) {
	tok, err := a.accessTokenRepo.FindOneByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	if tok.UserId != nil {
		usr, err := a.userRepo.Find(ctx, *tok.UserId)
		return usr, err
	}
	return nil, nil
}
