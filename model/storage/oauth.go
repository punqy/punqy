package storage

import (
	"github.com/google/uuid"
	"time"
)

type AccessToken struct {
	Entity
	Token     string     `json:"token" db:"token"`
	UserId    *uuid.UUID `json:"userId" db:"user_id"`
	ClientId  uuid.UUID  `json:"clientId" db:"client_id"`
	ExpiresAt time.Time  `json:"expiresAt" db:"expires_at"`
}

func (t AccessToken) Expired() bool {
	return t.ExpiresAt.UTC().Unix() <= time.Now().Unix()
}

func (t AccessToken) GetUserID() *string {
	if t.UserId != nil {
		uid := t.UserId.String()
		return &uid
	}
	return nil
}

func (t AccessToken) GetClientID() string {
	return t.ClientId.String()
}

type RefreshToken struct {
	Entity
	Token     string     `json:"token" db:"token"`
	UserId    *uuid.UUID `json:"userId" db:"user_id"`
	ClientId  uuid.UUID  `json:"clientId" db:"client_id"`
	ExpiresAt time.Time  `json:"expiresAt" db:"expires_at"`
}

func (t *RefreshToken) GetUserID() *string {
	if t.UserId != nil {
		uid := t.UserId.String()
		return &uid
	}
	return nil
}

func (t *RefreshToken) GetClientID() string {
	return t.ClientId.String()
}

func (t *RefreshToken) Expired() bool {
	return t.ExpiresAt.UTC().Unix() <= time.Now().UTC().Unix()
}

func (t *RefreshToken) SetExpired() {
	t.ExpiresAt = time.Now().UTC()
}

type OAuthClient struct {
	Entity
	GrantTypes   StringList `json:"grantTypes" db:"allowed_grant_types"`
	ClientSecret string     `json:"clientSecret" db:"client_secret"`
}

func (o OAuthClient) GetID() string {
	return o.ID.String()
}
