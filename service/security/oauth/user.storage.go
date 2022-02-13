package oauth

import (
	"context"
	"github.com/google/uuid"
	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/repository/user"
)

type UserManager interface {
	punqy.UserStorage
	punqy.UserProvider
}

type userManager struct {
	userRepo user.Repository
	encoder  punqy.PasswordEncoder
}

func NewUserManager(
	userRepo user.Repository,
	encoder punqy.PasswordEncoder,
) UserManager {
	return &userManager{
		userRepo: userRepo,
		encoder:  encoder,
	}
}

func (a *userManager) CheckCredentials(ctx context.Context, username, password string) (punqy.UserInterface, error) {
	usr, err := a.userRepo.FindUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if err := a.encoder.IsPasswordValid(usr.GetPassword(), password); err != nil {
		return usr, punqy.InvalidCredentialsErr()
	}
	return usr, nil
}

func (a *userManager) FindUserByUsername(ctx context.Context, username string) (punqy.UserInterface, error) {
	return a.userRepo.FindUserByUsername(ctx, username)
}

func (a *userManager) FindUserByID(ctx context.Context, id string) (punqy.UserInterface, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return a.userRepo.FindUserByID(ctx, uid)
}
