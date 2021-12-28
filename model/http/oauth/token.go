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

type GrantAccessTokenResponse struct {
	AccessToken           string `json:"accessToken"`
	RefreshToken          string `json:"refreshToken"`
	AccessTokenExpiresAt  int64  `json:"accessTokenExpiresAt"`
	RefreshTokenExpiresAt int64  `json:"refreshTokenExpiresAt"`
}

func FromPunqyResponse(r punqy.GrantAccessTokenResponse) GrantAccessTokenResponse {
	return GrantAccessTokenResponse(r)
}
