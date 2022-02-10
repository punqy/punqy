package user

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"regexp"
)

type PatchProfileRequest struct {
	Username string `json:"username"`
}

func (r PatchProfileRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Username,
			validation.Required,
			is.Email,
		),
	)
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (r ChangePasswordRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.OldPassword,
			validation.Required,
			validation.Match(regexp.MustCompile("[a-zA-Z0-9!@#$%^&*()_+~]{8,}")),
		),
		validation.Field(&r.NewPassword,
			validation.Required,
			validation.Match(regexp.MustCompile("[a-zA-Z0-9!@#$%^&*()_+~]{8,}")),
			validation.By(func(value interface{}) error {
				if value == r.OldPassword {
					return fmt.Errorf("new password must differ from old")
				}
				return nil
			}),
		),
	)
}
