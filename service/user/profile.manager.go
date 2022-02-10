package user

import (
	"context"
	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/model/http/user"
	"github.com/punqy/punqy/model/storage"
	userrepo "github.com/punqy/punqy/repository/user"
)

type ProfileManager interface {
	PatchProfile(context.Context, user.PatchProfileRequest, storage.User) error
	ChangePassword(ctx context.Context, req user.ChangePasswordRequest, usr storage.User) error
}

type profileManager struct {
	userRepo userrepo.Repository
	encoder  punqy.PasswordEncoder
}

func NewProfileManager(
	userRepo userrepo.Repository,
	encoder punqy.PasswordEncoder,
) ProfileManager {
	return &profileManager{
		userRepo: userRepo,
		encoder:  encoder,
	}
}

func (p *profileManager) PatchProfile(ctx context.Context, req user.PatchProfileRequest, usr storage.User) error {
	usr.Username = req.Username
	return p.userRepo.Update(ctx, usr)
}

func (p *profileManager) ChangePassword(ctx context.Context, req user.ChangePasswordRequest, usr storage.User) error {
	if err := p.encoder.IsPasswordValid(usr.Password, req.OldPassword); err != nil {
		return punqy.BadRequestErr("old password doesn't match")
	}
	passwordHash, err := p.encoder.EncodePassword(req.NewPassword, nil)
	if err != nil {
		return err
	}
	usr.Password = passwordHash
	return p.userRepo.Update(ctx, usr)
}
