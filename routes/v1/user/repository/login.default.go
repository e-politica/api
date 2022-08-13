package repository

import (
	"context"
	"errors"

	"github.com/e-politica/api/models/v1/user"
	"github.com/e-politica/api/pkg/database"
	"github.com/e-politica/api/pkg/session"
)

var (
	ErrInexistentAccount = errors.New("requested account does not exists")
)

func LoginDefault(ctx context.Context, db *database.Db, params user.RegisterDefault) (token string, err error) {
	found, err := isUserRegistered(db, *params.Email)
	if err != nil {
		return
	}

	if !found {
		err = ErrInexistentAccount
		return
	}

	userId, err := getUserId(db, *params.Email)
	if err != nil {
		return
	}

	password, err := getDefaultAccountPassword(db, userId)
	if err != nil {
		return
	}

	return session.NewSession(ctx, userId)
}
