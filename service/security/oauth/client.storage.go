package oauth

import (
	"context"
	"github.com/google/uuid"
	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/repository/oauth"
)

type clientStorage struct {
	clientRepo oauth.ClientRepository
}

func NewClientStorage(
	clientRepo oauth.ClientRepository,
) punqy.OAuthClientStorage {
	return &clientStorage{
		clientRepo: clientRepo,
	}
}

func (a *clientStorage) Find(ctx context.Context, id string) (punqy.OAuthClient, error) {
	var client punqy.OAuthClient
	uid, err := uuid.Parse(id)
	if err != nil {
		return client, err
	}

	client, err = a.clientRepo.Find(ctx, uid)
	if err != nil {
		return client, err
	}
	return client, nil
}

func (a *clientStorage) GetClient(ctx context.Context, id string, secret string, grantType punqy.GrantType) (punqy.OAuthClient, error) {
	var client punqy.OAuthClient
	uid, err := uuid.Parse(id)
	if err != nil {
		return client, err
	}

	client, err = a.clientRepo.FindOneByClientIdSecretAndGrantType(ctx, uid, secret, grantType)
	if err != nil {
		return client, err
	}
	return client, nil
}
