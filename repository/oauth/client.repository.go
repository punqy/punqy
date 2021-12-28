package oauth

import (
	"context"
	"fmt"
	punqy "github.com/punqy/core"
	model "github.com/punqy/punqy/model/storage"
	"github.com/punqy/punqy/repository/tables"
	"github.com/slmder/qbuilder"
)

type ClientRepository interface {
	Find(ctx context.Context, id string) (punqy.OAuthClient, error)
	FindOneByClientIdSecretAndGrantType(ctx context.Context, cID, sec string, gt punqy.GrantType) (punqy.OAuthClient, error)
	Insert(ctx context.Context, entity *model.OAuthClient) error
	NewOauthClient(ctx context.Context) (model.OAuthClient, error)
}

type clientRepository struct {
	punqy.Dal
}

func NewClientRepository(db punqy.Dal) ClientRepository {
	return &clientRepository{Dal: db}
}

func (r *clientRepository) NewOauthClient(ctx context.Context) (model.OAuthClient, error) {
	e := model.OAuthClient{
		GrantTypes:   model.StringList{punqy.GrantTypeRefreshToken.String(), punqy.ClientCredentials.String(), punqy.GrantTypePassword.String()},
		ClientSecret: fmt.Sprintf("cs_%s", punqy.RandomString(64)),
	}
	if err := e.Init(); err != nil {
		return e, err
	}

	if err := r.Insert(ctx, &e); err != nil {
		return model.OAuthClient{}, err
	}
	return e, nil
}

func (r *clientRepository) FindOneByClientIdSecretAndGrantType(ctx context.Context, cID, sec string, gt punqy.GrantType) (punqy.OAuthClient, error) {
	var entity model.OAuthClient
	query := r.SelectE(entity).
		From(tables.OAuthClient).
		Where("id = $1 AND client_secret = $2").
		AndWhere("$3 IN (select jsonb_array_elements_text(allowed_grant_types))").
		Limit(1).
		ToSQL()
	err := r.DoSelectOne(ctx, &entity, query, cID, sec, gt)
	return entity, r.PipeErr(err)
}

func (r *clientRepository) Insert(ctx context.Context, entity *model.OAuthClient) error {
	_, err := r.InsertE(ctx, tables.OAuthClient, entity)
	return r.PipeErr(err)
}

func (r *clientRepository) FindBy(ctx context.Context, cond qbuilder.Conditions, pager punqy.Pagination) ([]model.OAuthClient, error) {
	var entities []model.OAuthClient
	return entities, r.Dal.FindBy(ctx, tables.OAuthClient, &entities, cond, pager)
}

func (r *clientRepository) FindOneBy(ctx context.Context, cond qbuilder.Conditions) (model.OAuthClient, error) {
	entity := model.OAuthClient{}
	return entity, r.Dal.FindOneBy(ctx, tables.OAuthClient, &entity, cond)
}

func (r *clientRepository) Find(ctx context.Context, id string) (punqy.OAuthClient, error) {
	return r.FindOneBy(ctx, qbuilder.Conditions{"id": id})
}
