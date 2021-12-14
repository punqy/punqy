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
	Find(ctx context.Context, id uint32) (*model.RefreshToken, error)
	FindOneByToken(ctx context.Context, token string) (punqy.OAuthRefreshToken, error)
	Insert(ctx context.Context, entity *model.RefreshToken) error
	Update(ctx context.Context, entity punqy.OAuthRefreshToken) error
	Create(ctx context.Context, entity punqy.OAuthRefreshTokenValues) error
}

type refreshTokenRepository struct {
	punqy.Dal
}

func NewRefreshTokenRepository(db punqy.Dal) RefreshTokenRepository {
	return &refreshTokenRepository{db}
}

func (r *refreshTokenRepository) FindOneByToken(ctx context.Context, token string) (punqy.OAuthRefreshToken, error) {
	return r.FindOneBy(ctx, qbuilder.Conditions{"token": token})
}

func (r *refreshTokenRepository) Create(ctx context.Context, values punqy.OAuthRefreshTokenValues) error {
	var userID *uuid.UUID
	if values.UserId != nil {
		uid, err := uuid.Parse(*values.UserId)
		if err != nil {
			return err
		}
		userID = &uid
	}
	cid, err := uuid.Parse(values.ClientId)
	if err != nil {
		return err
	}

	e := model.RefreshToken{
		Token:     values.Token,
		UserId:    userID,
		ClientId:  cid,
		ExpiresAt: values.ExpiresAt,
	}
	if err := e.NewId(); err != nil {
		return err
	}
	return r.Insert(ctx, &e)
}

func (r *refreshTokenRepository) Insert(ctx context.Context, entity *model.RefreshToken) error {
	sql := r.BuildInsert(tables.OAuthRefreshToken).
		Columns("id", "token", "user_id", "client_id", "expires_at", "created_at", "updated_at").
		Value(":id, :token, :user_id, :client_id, :expires_at, now(), now()").
		ToSQL()

	_, err := r.DoInsert(ctx, sql, entity)
	return r.PipeErr(err)
}

func (r *refreshTokenRepository) Update(ctx context.Context, entity punqy.OAuthRefreshToken) error {
	query := r.BuildUpdate(tables.OAuthRefreshToken).
		SetMap(punqy.StringMap{
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

func (r *refreshTokenRepository) FindBy(ctx context.Context, cond qbuilder.Conditions, pager punqy.Pagination) ([]model.RefreshToken, error) {
	var entities []model.RefreshToken
	return entities, r.Dal.FindBy(ctx, tables.OAuthRefreshToken, &entities, cond, pager)
}

func (r *refreshTokenRepository) FindOneBy(ctx context.Context, cond qbuilder.Conditions) (*model.RefreshToken, error) {
	entity := &model.RefreshToken{}
	return entity, r.Dal.FindOneBy(ctx, tables.OAuthRefreshToken, entity, cond)
}

func (r *refreshTokenRepository) Find(ctx context.Context, id uint32) (*model.RefreshToken, error) {
	return r.FindOneBy(ctx, qbuilder.Conditions{"id": id})
}
