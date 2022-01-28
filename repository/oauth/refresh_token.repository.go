package oauth

import (
	"context"
	"github.com/google/uuid"
	punqy "github.com/punqy/core"
	model "github.com/punqy/punqy/model/storage"
	"github.com/punqy/punqy/repository/tables"
	"github.com/slmder/qbuilder"
)

type RefreshTokenRepository interface {
	Find(ctx context.Context, id uuid.UUID) (model.RefreshToken, error)
	FindOneByToken(ctx context.Context, token string) (model.RefreshToken, error)
	Insert(ctx context.Context, entity model.RefreshToken) error
	Update(ctx context.Context, entity model.RefreshToken) error
}

type refreshTokenRepository struct {
	punqy.Dal
}

func NewRefreshTokenRepository(db punqy.Dal) RefreshTokenRepository {
	return &refreshTokenRepository{db}
}

func (r *refreshTokenRepository) FindOneByToken(ctx context.Context, token string) (model.RefreshToken, error) {
	return r.FindOneBy(ctx, qbuilder.Conditions{"token": token})
}

func (r *refreshTokenRepository) Insert(ctx context.Context, entity model.RefreshToken) error {
	_, err := r.InsertE(ctx, tables.OAuthRefreshToken, entity)
	return r.PipeErr(err)
}

func (r *refreshTokenRepository) Update(ctx context.Context, entity model.RefreshToken) error {
	_, err := r.UpdateE(ctx, tables.OAuthRefreshToken, entity)
	return r.PipeErr(err)
}

func (r *refreshTokenRepository) FindBy(ctx context.Context, cond qbuilder.Conditions, pager punqy.Pagination) ([]model.RefreshToken, error) {
	var entities []model.RefreshToken
	return entities, r.Dal.FindBy(ctx, tables.OAuthRefreshToken, &entities, cond, pager)
}

func (r *refreshTokenRepository) FindOneBy(ctx context.Context, cond qbuilder.Conditions) (model.RefreshToken, error) {
	var entity model.RefreshToken
	return entity, r.Dal.FindOneBy(ctx, tables.OAuthRefreshToken, entity, cond)
}

func (r *refreshTokenRepository) Find(ctx context.Context, id uuid.UUID) (model.RefreshToken, error) {
	return r.FindOneBy(ctx, qbuilder.Conditions{"id": id})
}
