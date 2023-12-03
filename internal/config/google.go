package config

import (
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewGoogleConfig(cfg *Config) *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  fmt.Sprintf("%s/betty-api/auth/google/callback", cfg.BaseURL),
		ClientID:     cfg.Google.ClientID,
		ClientSecret: cfg.Google.ClientSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}
