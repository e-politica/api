package repository

import (
	"github.com/e-politica/api/config"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var oauthConfig = &oauth2.Config{
	ClientID:     config.GoogleClientId,
	ClientSecret: config.GoogleClientSecret,
	RedirectURL:  config.GoogleLoginRedirectUrl,
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}

// Return random generated state and a URL to OAuth 2.0 provider's consent page.
func Login() (string, string) {
	state := uuid.NewString()
	return state, oauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
}
