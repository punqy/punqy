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

type RefreshTokenStorage struct {
	tokenRepo oauth.RefreshTokenRepository
	userRepo  user.Repository
}

func NewRefreshTokenStorage(
	tokenRepo oauth.RefreshTokenRepository,
	userRepo user.Repository,
) punqy.OAuthRefreshTokenStorage {
	return &RefreshTokenStorage{
		tokenRepo: tokenRepo,
		userRepo:  userRepo,
	}
}

func (a *RefreshTokenStorage) CreateRefreshToken(ctx context.Context, user punqy.User, client punqy.OAuthClient, token string, expiresAt time.Time) error {
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

	e := model.RefreshToken{
		Token:     token,
		UserId:    userID,
		ClientId:  cid,
		ExpiresAt: expiresAt,
	}
	if err := e.Init(); err != nil {
		return err
	}
	return a.tokenRepo.Insert(ctx, e)

}

func (a *RefreshTokenStorage) CheckCredentials(ctx context.Context, token string) (punqy.User, error) {
	tok, err := a.tokenRepo.FindOneByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	if tok.UserId != nil {
		usr, err := a.userRepo.Find(ctx, *tok.UserId)
		return usr, err
	}
	return nil, nil
}
