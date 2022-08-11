package repository

import (
	"context"

	"github.com/e-politica/api/config"
	"github.com/e-politica/api/pkg/database"
	"github.com/e-politica/api/pkg/session"
	goauth2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func LoginGoogle(ctx context.Context, db *database.Db, credential string) (token string, err error) {
	userInfo, err := validateGoogleCredential(ctx, credential)
	if err != nil {
		return
	}

	found, err := isUserRegistered(db, userInfo.Email)
	if err != nil {
		return
	}

	if found {
		userId, err := getUserId(db, userInfo.Email)
		if err != nil {
			return "", err
		}

		return session.NewSession(ctx, userId)
	}

	// ... register user
	return
}

func validateGoogleCredential(ctx context.Context, credential string) (*goauth2.Tokeninfo, error) {
	service, err := goauth2.NewService(
		ctx,
		option.WithAPIKey(config.GoogleClientSecret),
	)
	if err != nil {
		return nil, err
	}

	return service.Tokeninfo().IdToken(credential).Do()
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

func getUserId(db *database.Db, email string) (id string, err error) {
	query := `
	SELECT user_id
	FROM user_account
	WHERE user_email=$1
	`

	row, err := db.QueryRow(query, email)
	if err != nil {
		return
	}

	err = row.Scan(&id)
	return
}
