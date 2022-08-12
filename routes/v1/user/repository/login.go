package repository

import (
	"context"
	"errors"

	"github.com/e-politica/api/config"
	"github.com/e-politica/api/pkg/database"
	"github.com/e-politica/api/pkg/session"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	goauth2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func LoginGoogle(ctx context.Context, db *database.Db, credential string) (string, error) {
	claims, err := validateGoogleCredential(ctx, credential)
	if err != nil {
		return "", err
	}

	found, err := isUserRegistered(db, claims["email"].(string))
	if err != nil {
		return "", err
	}

	if found {
		userId, err := getUserId(db, claims["email"].(string))
		if err != nil {
			return "", err
		}

		token, err := session.GetSessionToken(ctx, userId)
		if err != nil {
			if err == redis.Nil {
				return session.NewSession(ctx, userId)
			}
			return "", err
		}

		return token, err
	}

	userId, err := insertUser(
		db,
		claims["name"].(string),
		claims["email"].(string),
		claims["picture"].(string),
		"google",
	)
	if err != nil {
		return "", err
	}

	return session.NewSession(ctx, userId)
}

func validateGoogleCredential(ctx context.Context, credential string) (jwt.MapClaims, error) {
	service, err := goauth2.NewService(
		ctx,
		option.WithAPIKey(config.GoogleClientSecret),
	)
	if err != nil {
		return nil, err
	}

	_, err = service.Tokeninfo().IdToken(credential).Do()
	if err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(credential, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte{}, nil
	})
	if err != nil && !errors.Is(err, jwt.ErrInvalidKeyType) {
		return nil, err
	}

	return claims, nil
}

func isUserRegistered(db *database.Db, email string) (bool, error) {
	query := `
	SELECT COUNT(1)
	FROM user_account
	WHERE user_email=$1
	`

	row, err := db.QueryRow(query, email)
	if err != nil {
		return false, err
	}

	var count int
	err = row.Scan(&count)

	return count == 1, err
}

func getUserId(db *database.Db, email string) (userId string, err error) {
	query := `
	SELECT user_id
	FROM user_account
	WHERE user_email=$1
	`

	row, err := db.QueryRow(query, email)
	if err != nil {
		return
	}

	err = row.Scan(&userId)
	return
}

func insertUser(db *database.Db, name, email, avatarUrl, accountType string) (userId string, err error) {
	query := `
	INSERT INTO user_account (
		user_name,
		user_email,
		user_avatar_url,
		user_account_type
	)
	VALUES ($1, $2, $3, $4)
	RETURNING user_id
	`

	row, err := db.QueryRow(query, name, email, avatarUrl, accountType)
	if err != nil {
		return
	}

	err = row.Scan(&userId)
	return
}
