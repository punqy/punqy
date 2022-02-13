package storage

type OAuthClient struct {
	Entity
	GrantTypes   StringList `json:"grantTypes" db:"allowed_grant_types"`
	ClientSecret string     `json:"clientSecret" db:"client_secret"`
}

func (o OAuthClient) GetID() string {
	return o.ID.String()
}
