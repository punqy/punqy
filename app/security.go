package app

import (
	punqy "github.com/punqy/core"
)

func SecurityConfig(module security.ModuleSecurity) punqy.SecurityConfig {
	return punqy.SecurityConfig{
		Firewall: punqy.FirewallConfig{
			punqy.Area{
				Secure:  false,
				Pattern: "^/dev/",
			},
			punqy.Area{
				Secure:  false,
				Pattern: "^/api/oauth/token",
			},
			punqy.Area{
				Secure:        true,
				Pattern:       "^/api/",
				Authenticator: module.OAuthAuthenticator(),
			},
		},
	}
}
