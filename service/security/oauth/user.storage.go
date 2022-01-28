package oauth

import (
	"context"
	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/repository/user"
)

type UserStorage struct {
	userRepo user.Repository
	encoder  punqy.PasswordEncoder
}

func NewUserStorage(
	userRepo user.Repository,
	encoder punqy.PasswordEncoder,
) punqy.UserStorage {
	return &UserStorage{
		userRepo: userRepo,
		encoder:  encoder,
	}
}

func (a *UserStorage) CheckCredentials(ctx context.Context, username, password string) (punqy.User, error) {
	user, err := a.userRepo.FindUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if valid, _ := a.encoder.IsPasswordValid(user.GetPassword(), password); !valid {
		return user, punqy.InvalidCredentialsErr()
	}
	return user, nil
}
