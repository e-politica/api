package repository

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"github.com/e-politica/api/config"
	"github.com/e-politica/api/pkg/database"
	"github.com/e-politica/api/pkg/session"
	"github.com/go-redis/redis/v8"
	goauth2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

var (
	ErrInvalidJwt = errors.New("invalid jwt")
)

type googleJwtPayload struct {
	Iss           string `json:"iss"`
	Nbf           int    `json:"nbf"`
	Aud           string `json:"aud"`
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Azp           string `json:"azp"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Iat           int    `json:"iat"`
	Exp           int    `json:"exp"`
	Jti           string `json:"jti"`
}

func LoginGoogle(ctx context.Context, db *database.Db, credentials string) (sess session.Session, err error) {
	payload, err := validateGoogleCredential(ctx, credentials)
	if err != nil {
		return
	}

	found, err := isUserRegistered(db, payload.Email)
	if err != nil {
		return
	}

	if found {
		return getUserSession(ctx, db, payload.Email)
	}

	userId, err := insertUser(
		db,
		payload.Name,
		payload.Email,
		payload.Picture,
		"google",
	)
	if err != nil {
		return
	}

	return session.NewSession(ctx, userId)
}

func validateGoogleCredential(ctx context.Context, credentials string) (payload googleJwtPayload, err error) {
	service, err := goauth2.NewService(
		ctx,
		option.WithAPIKey(config.GoogleClientSecret),
	)
	if err != nil {
		return
	}

	_, err = service.Tokeninfo().IdToken(credentials).Do()
	if err != nil {
		return
	}

	credSplited := strings.Split(credentials, ".")
	if len(credSplited) != 3 {
		err = ErrInvalidJwt
		return
	}

	decodedPayload, err := base64.RawURLEncoding.DecodeString(credSplited[1])
	if err != nil {
		return
	}

	err = json.Unmarshal(decodedPayload, &payload)
	return
}

func getUserSession(ctx context.Context, db *database.Db, email string) (sess session.Session, err error) {
	userId, err := getUserId(db, email)
	if err != nil {
		return
	}

	sess, err = session.GetSession(ctx, userId)
	if err == redis.Nil {
		return session.NewSession(ctx, userId)
	}

	return
}
