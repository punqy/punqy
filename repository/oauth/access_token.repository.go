package oauth

import (
	"context"
	"github.com/google/uuid"
	punqy "github.com/punqy/core"
	model "github.com/punqy/punqy/model/storage"
	"github.com/punqy/punqy/repository/tables"
	"github.com/slmder/qbuilder"
)

type AccessTokenRepository interface {
	Find(ctx context.Context, id uint32) (*model.AccessToken, error)
	Insert(ctx context.Context, entity *model.AccessToken) error
	Create(ctx context.Context, entity punqy.OAuthAccessTokenValues) error
	FindOneByToken(ctx context.Context, token string) (punqy.OAuthAccessToken, error)
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

func (r *accessTokenRepository) FindOneByToken(ctx context.Context, token string) (punqy.OAuthAccessToken, error) {
	return r.FindOneBy(ctx, qbuilder.Conditions{"token": token})
}

func (r *accessTokenRepository) Create(ctx context.Context, values punqy.OAuthAccessTokenValues) error {
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

	e := model.AccessToken{
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

func (r *accessTokenRepository) Insert(ctx context.Context, entity *model.AccessToken) error {
	sql := r.BuildInsert(tables.OAuthAccessToken).
		Columns("id", "token", "user_id", "client_id", "expires_at", "created_at", "updated_at").
		Value(":id, :token, :user_id, :client_id, :expires_at, now(), now()").
		ToSQL()

	_, err := r.DoInsert(ctx, sql, entity)
	return r.PipeErr(err)
}

func (r *accessTokenRepository) FindBy(ctx context.Context, cond qbuilder.Conditions, pager punqy.Pagination) ([]model.AccessToken, error) {
	entities := NewAccessTokenList()
	return entities, r.Dal.FindBy(ctx, tables.OAuthAccessToken, &entities, cond, pager)
}

func (r *accessTokenRepository) FindOneBy(ctx context.Context, cond qbuilder.Conditions) (*model.AccessToken, error) {
	var entity *model.AccessToken
	return entity, r.Dal.FindOneBy(ctx, tables.OAuthAccessToken, entity, cond)
}

func (r *accessTokenRepository) Find(ctx context.Context, id uint32) (*model.AccessToken, error) {
	return r.FindOneBy(ctx, qbuilder.Conditions{"id": id})
}
