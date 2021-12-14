package oauth

import (
	"context"
	"github.com/google/uuid"
	"github.com/slmder/qbuilder"
	model "github.com/picknpark/wshub/model/storage"
	"github.com/picknpark/wshub/pkg/storage"
	"github.com/picknpark/wshub/pkg/util"
	"github.com/picknpark/wshub/repository/schema/tables"
	"time"
)

type RefreshTokenRepository interface {
	NewRefreshToken(ctx context.Context, userId *uuid.UUID, clientId uuid.UUID, tll int) (model.RefreshToken, error)
	Find(ctx context.Context, id uint32) (model.RefreshToken, error)
	FindOneByToken(ctx context.Context, token string) (model.RefreshToken, error)
	Insert(ctx context.Context, entity *model.RefreshToken) error
	Update(ctx context.Context, entity model.RefreshToken) error
}

type refreshTokenRepository struct {
	storage.Dal
}

func NewRefreshTokenRepository(db storage.Dal) RefreshTokenRepository {
	return &refreshTokenRepository{db}
}

func (r *refreshTokenRepository) NewRefreshToken(ctx context.Context, userId *uuid.UUID, clientId uuid.UUID, tll int) (model.RefreshToken, error) {
	e := model.RefreshToken{
		Token:     util.RandomString(64),
		UserId:    userId,
		ClientId:  clientId,
		ExpiresAt: time.Now().Add(time.Duration(tll) * time.Minute),
	}
	if err := e.NewId(); err != nil {
		return e, err
	}

	if err := r.Insert(ctx, &e); err != nil {
		return model.RefreshToken{}, err
	}
	return e, nil
}

func (r *refreshTokenRepository) FindOneByToken(ctx context.Context, token string) (model.RefreshToken, error) {
	return r.FindOneBy(ctx, qbuilder.Conditions{"token": token})
}

func (r *refreshTokenRepository) Insert(ctx context.Context, entity *model.RefreshToken) error {
	sql := r.BuildInsert(tables.OAuthRefreshToken).
		Columns("id", "token", "user_id", "client_id", "expires_at", "created_at", "updated_at").
		Value(":id, :token, :user_id, :client_id, :expires_at, now(), now()").
		ToSQL()

	_, err := r.DoInsert(ctx, sql, entity)
	return r.PipeErr(err)
}

func (r *refreshTokenRepository) Update(ctx context.Context, entity model.RefreshToken) error {
	query := r.BuildUpdate(tables.OAuthRefreshToken).
		SetMap(storage.StringMap{
			"user_id":    ":user_id",
			"client_id":  ":client_id",
			"expires_at": ":expires_at",
			"updated_at": "now()",
		}).
		Where("id = :id").
		ToSQL()

	_, err := r.DoUpdate(ctx, query, entity)
	return r.PipeErr(err)
}

func (r *refreshTokenRepository) FindBy(ctx context.Context, cond qbuilder.Conditions, pager storage.Pagination) ([]model.RefreshToken, error) {
	var entities []model.RefreshToken
	return entities, r.Dal.FindBy(ctx, tables.OAuthRefreshToken, &entities, cond, pager)
}

func (r *refreshTokenRepository) FindOneBy(ctx context.Context, cond qbuilder.Conditions) (model.RefreshToken, error) {
	entity := model.RefreshToken{}
	return entity, r.Dal.FindOneBy(ctx, tables.OAuthRefreshToken, &entity, cond)
}

func (r *refreshTokenRepository) Find(ctx context.Context, id uint32) (model.RefreshToken, error) {
	return r.FindOneBy(ctx, qbuilder.Conditions{"id": id})
}
