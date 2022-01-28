package oauth

import (
	"context"
	punqy "github.com/punqy/core"
	model "github.com/punqy/punqy/model/storage"
	"github.com/punqy/punqy/repository/tables"
	"github.com/slmder/qbuilder"
)

type AccessTokenRepository interface {
	Find(ctx context.Context, id uint32) (model.AccessToken, error)
	Insert(ctx context.Context, entity model.AccessToken) error
	FindOneByToken(ctx context.Context, token string) (model.AccessToken, error)
}

type accessTokenRepository struct {
	punqy.Dal
}


func NewAccessToken() model.AccessToken {
	return model.AccessToken{}
}

func NewAccessTokenList() []model.AccessToken {
	return []model.AccessToken{}
}

func NewAccessTokenRepository(db punqy.Dal) AccessTokenRepository {
	return &accessTokenRepository{db}
}

func (r *accessTokenRepository) FindOneByToken(ctx context.Context, token string) (model.AccessToken, error) {
	return r.FindOneBy(ctx, qbuilder.Conditions{"token": token})
}

func (r *accessTokenRepository) Insert(ctx context.Context, entity model.AccessToken) error {
	_, err := r.InsertE(ctx, tables.OAuthAccessToken, entity)
	return r.PipeErr(err)
}

func (r *accessTokenRepository) FindBy(ctx context.Context, cond qbuilder.Conditions, pager punqy.Pagination) ([]model.AccessToken, error) {
	entities := NewAccessTokenList()
	return entities, r.Dal.FindBy(ctx, tables.OAuthAccessToken, &entities, cond, pager)
}

func (r *accessTokenRepository) FindOneBy(ctx context.Context, cond qbuilder.Conditions) (model.AccessToken, error) {
	var entity model.AccessToken
	return entity, r.Dal.FindOneBy(ctx, tables.OAuthAccessToken, &entity, cond)
}

func (r *accessTokenRepository) Find(ctx context.Context, id uint32) (model.AccessToken, error) {
	return r.FindOneBy(ctx, qbuilder.Conditions{"id": id})
}
