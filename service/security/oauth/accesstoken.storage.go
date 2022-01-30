package oauth

import (
	"context"
	punqy "github.com/punqy/core"
)

type accessTokenStorage struct {
	jwtTokenSigningKey string
	tokenTTL           int
}

func NewAccessTokenStorage(
	jwtTokenSigningKey string,
	tokenTTL int,
) punqy.OAuthAccessTokenStorage {
	return &accessTokenStorage{
		jwtTokenSigningKey: jwtTokenSigningKey,
		tokenTTL:           tokenTTL,
	}
}

func (a *accessTokenStorage) CheckCredentials(ctx context.Context, token string) (punqy.OAuthAccessToken, error) {
	return CheckTokenCredentials(ctx, token, a.jwtTokenSigningKey, TokenTypeAccess)
}

func (a *accessTokenStorage) CreateAccessToken(ctx context.Context, user punqy.UserInterface, client punqy.OAuthClient) (punqy.TokenValues, error) {
	return CreateToken(ctx, TokenTypeAccess, a.tokenTTL, a.jwtTokenSigningKey, user, client)
}
