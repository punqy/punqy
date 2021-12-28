package user

import (
	"context"
	"github.com/google/uuid"
	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/model/storage"
	"github.com/punqy/punqy/repository/tables"
	"github.com/slmder/qbuilder"
)

type Repository interface {
	Find(ctx context.Context, id uuid.UUID) (punqy.User, error)
	FindUserByID(ctx context.Context, id string) (punqy.User, error)
	FindUserByUsername(ctx context.Context, username string) (punqy.User, error)
	Insert(ctx context.Context, entity storage.User) error
}

type repository struct {
	punqy.Dal
}

func NewUserRepository(db punqy.Dal) Repository {
	return &repository{Dal: db}
}

func (r *repository) FindUserByID(ctx context.Context, id string) (punqy.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return r.Find(ctx, uid)
}

func (r *repository) FindUserByUsername(ctx context.Context, username string) (punqy.User, error) {
	return r.FindOneBy(ctx, qbuilder.Conditions{
		"username": username,
	})
}

func (r *repository) Insert(ctx context.Context, entity storage.User) error {
	_, err := r.InsertE(ctx, tables.User, entity)
	return r.PipeErr(err)
}

func (r *repository) FindBy(ctx context.Context, cond qbuilder.Conditions, pager punqy.Pagination) ([]storage.User, error) {
	var entities []storage.User
	return entities, r.Dal.FindBy(ctx, tables.User, &entities, cond, pager)
}

func (r *repository) FindOneBy(ctx context.Context, cond qbuilder.Conditions) (punqy.User, error) {
	var entity storage.User
	err := r.Dal.FindOneBy(ctx, tables.User, &entity, cond)
	return entity, r.PipeErr(err)
}

func (r *repository) Find(ctx context.Context, id uuid.UUID) (punqy.User, error) {
	return r.FindOneBy(ctx, qbuilder.Conditions{"id": id})
}
