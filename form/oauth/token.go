package oauth

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	punqy "github.com/punqy/core"
)

type GrantAccessTokenRequest struct {
	GrantType    punqy.GrantType `json:"grantType" valid:"required,in(client_credentials|password)~invalid grant type"`
	ClientSecret string          `json:"clientSecret" valid:"required"`
	RefreshToken string          `json:"refreshToken"`
	ClientId     string          `json:"clientId" valid:"required"`
	Username     string          `json:"username"`
	Password     string          `json:"password"`
}

func (r GrantAccessTokenRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Username,
			validation.When(r.GrantType == punqy.GrantTypePassword, validation.Required),
			is.Email,
		),
		validation.Field(
			&r.Password,
			validation.When(r.GrantType == punqy.GrantTypePassword, validation.Required),
		),
		validation.Field(
			&r.ClientId,
			validation.Required,
			is.UUIDv4,
		),
		validation.Field(
			&r.ClientSecret,
			validation.Required,
		),
		validation.Field(
			&r.GrantType,
			validation.Required,
			validation.In(punqy.GrantTypeRefreshToken, punqy.GrantTypePassword, punqy.ClientCredentials),
		),
		validation.Field(
			&r.RefreshToken,
			validation.When(r.GrantType == punqy.GrantTypeRefreshToken, validation.Required),
		),
	)
}
