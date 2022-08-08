package config

import "github.com/e-politica/api/pkg/env"

var (
	ServerPort = env.Get("SERVER_PORT", "8080")

	GoogleClientId         = env.Get("GOOGLE_CLIENT_ID", "")
	GoogleClientSecret     = env.Get("GOOGLE_CLIENT_SECRET", "")
	GoogleLoginRedirectUrl = env.Get("GOOGLE_LOGIN_REDIRECT_URL", "")
)
