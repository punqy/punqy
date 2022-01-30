package oauth

import (
	"context"
	"github.com/golang-jwt/jwt"
	punqy "github.com/punqy/core"
	"time"
)

type TokenType int

const (
	TokenTypeAccess TokenType = iota
	TokenTypeRefresh
)

type AuthToken struct {
	Token    string  `json:"token" `
	UserId   *string `json:"userId"`
	ClientId string  `json:"clientId"`
}

func (a AuthToken) GetUserID() *string {
	return a.UserId
}

func (a AuthToken) GetClientID() string {
	return a.ClientId
}

type TokenClaims struct {
	jwt.StandardClaims
	UserID   *string   `json:"user_id"`
	UserName string    `json:"user_name"`
	ClientID string    `json:"client_id"`
	Type     TokenType `json:"type"`
}

type Storage struct {
	jwtTokenSigningKey string
	tokenTTL           int
}

func CheckTokenCredentials(ctx context.Context, token string, jwtTokenSigningKey string, t TokenType) (punqy.OAuthAccessToken, error) {
	jt, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, punqy.InvalidCredentialsErr("invalid Signing Method")
		}
		claims, ok := token.Claims.(*TokenClaims)
		if !ok || claims.Type != t {
			return nil, punqy.InvalidCredentialsErr("token type invalid")
		}
		if claims.ExpiresAt <= time.Now().Unix() {
			return nil, punqy.InvalidCredentialsErr("expired token")
		}
		return []byte(jwtTokenSigningKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims := jt.Claims.(*TokenClaims)
	return AuthToken{
		Token:    token,
		UserId:   claims.UserID,
		ClientId: claims.ClientID,
	}, nil
}

func CreateToken(ctx context.Context, t TokenType, tokenTTL int, jwtTokenSigningKey string, user punqy.UserInterface, client punqy.OAuthClient) (punqy.TokenValues, error) {
	var values punqy.TokenValues
	now := time.Now()
	expires := now.Add(time.Duration(tokenTTL) * time.Minute)
	claims := TokenClaims{}
	claims.ExpiresAt = expires.Unix()
	claims.IssuedAt = now.Unix()
	if user != nil {
		uid := user.GetID()
		claims.UserID = &uid
		claims.UserName = user.GetUsername()
	}
	claims.Type = t
	claims.ClientID = client.GetID()
	claims.Issuer = "Punqy oauth"
	claims.Audience = "Punqy oauth"
	jt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := jt.SignedString([]byte(jwtTokenSigningKey))
	if err != nil {
		return values, err
	}
	values.Token = tokenString
	values.ExpiresAt = expires
	return values, nil
}

func (a *Storage) CreateRefreshToken(ctx context.Context, user punqy.UserInterface, client punqy.OAuthClient) (punqy.TokenValues, error) {
	var values punqy.TokenValues
	now := time.Now()
	expires := now.Add(time.Duration(a.tokenTTL) * time.Minute)
	claims := TokenClaims{}
	claims.ExpiresAt = expires.Unix()
	claims.IssuedAt = now.Unix()
	if user != nil {
		uid := user.GetID()
		claims.UserID = &uid
		claims.UserName = user.GetUsername()
	}
	claims.Type = TokenTypeRefresh
	claims.ClientID = client.GetID()
	claims.Issuer = "Punqy oauth"
	claims.Audience = "Punqy oauth"
	jt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := jt.SignedString([]byte(a.jwtTokenSigningKey))
	if err != nil {
		return values, err
	}
	values.Token = tokenString
	values.ExpiresAt = expires
	return values, nil
}
