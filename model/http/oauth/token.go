package oauth

import punqy "github.com/punqy/core"

type GrantAccessTokenRequest struct {
	GrantType    punqy.GrantType `json:"grantType" valid:"required,in(client_credentials|password)~invalid grant type"`
	ClientSecret string          `json:"clientSecret" valid:"required"`
	RefreshToken string          `json:"refreshToken"`
	ClientId     string          `json:"clientId" valid:"required"`
	Username     string          `json:"username"`
	Password     string          `json:"password"`
}

func (r GrantAccessTokenRequest) ToPunqyRequest() punqy.GrantAccessTokenRequest {
	return punqy.GrantAccessTokenRequest(r)
}
