package user

import (
	"context"
	"github.com/google/uuid"
	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/repository/tables"
	"github.com/slmder/qbuilder"
)

type Repository interface {
	Find(ctx context.Context, id uuid.UUID) (punqy.User, error)
	FindUserByID(ctx context.Context, id string) (punqy.User, error)
}

type repository struct {
	punqy.Dal
}

func NewUserRepository(db punqy.Dal) Repository {
	return &repository{Dal: db}
}

func (r *repository) FindUserByID(ctx context.Context, id string) (punqy.User, error) {
	return r.Find(ctx, uuid.MustParse(id))
}

func (r *repository) FindBy(ctx context.Context, cond qbuilder.Conditions, pager punqy.Pagination) ([]punqy.User, error) {
	var entities []punqy.User
	return entities, r.Dal.FindBy(ctx, tables.User, &entities, cond, pager)
}

func (r *repository) FindOneBy(ctx context.Context, cond qbuilder.Conditions) (punqy.User, error) {
	var entity punqy.User
	return entity, r.Dal.FindOneBy(ctx, tables.User, &entity, cond)
}

func (r *repository) Find(ctx context.Context, id uuid.UUID) (punqy.User, error) {
	return r.FindOneBy(ctx, qbuilder.Conditions{"id": id})
}
