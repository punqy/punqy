package app

import (
	punqy "github.com/punqy/core"
)

func SecurityConfig(oAuth punqy.OAuthAuthenticator) punqy.SecurityConfig {
	return punqy.SecurityConfig{
		Firewall: punqy.FirewallConfig{
			punqy.Area{
				Secure:  false,
				Pattern: "^/dev/",
			},
			punqy.Area{
				Secure:  false,
				Pattern: "^/api/v1/oauth/token",
			},
			punqy.Area{
				Secure:        true,
				Pattern:       "^/api/v1",
				Authenticator: oAuth,
			},
		},
	}
}
