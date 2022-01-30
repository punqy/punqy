package oauth

import (
	"context"
	punqy "github.com/punqy/core"
)

type refreshTokenStorage struct {
	jwtTokenSigningKey string
	tokenTTL           int
}

func NewRefreshTokenStorage(
	jwtTokenSigningKey string,
	tokenTTL int,
) punqy.OAuthRefreshTokenStorage {
	return &refreshTokenStorage{
		jwtTokenSigningKey: jwtTokenSigningKey,
		tokenTTL:           tokenTTL,
	}
}

func (a *refreshTokenStorage) CheckCredentials(ctx context.Context, token string) (punqy.OAuthRefreshToken, error) {
	return CheckTokenCredentials(ctx, token, a.jwtTokenSigningKey, TokenTypeRefresh)
}

func (a *refreshTokenStorage) CreateRefreshToken(ctx context.Context, user punqy.UserInterface, client punqy.OAuthClient) (punqy.TokenValues, error) {
	return CreateToken(ctx, TokenTypeRefresh, a.tokenTTL, a.jwtTokenSigningKey, user, client)
}
