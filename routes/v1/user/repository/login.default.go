package repository

import (
	"context"
	"errors"

	"github.com/e-politica/api/models/v1/user"
	"github.com/e-politica/api/pkg/crypto"
	"github.com/e-politica/api/pkg/database"
	"github.com/e-politica/api/pkg/session"
	"github.com/go-redis/redis/v8"
)

var (
	ErrInexistentAccount  = errors.New("requested account does not exists")
	ErrPasswordsDontMatch = errors.New("passwords do not match")
)

func LoginDefault(ctx context.Context, db *database.Db, params user.LoginDefaultParams) (sess session.Session, err error) {
	found, err := isUserRegistered(db, params.Email)
	if err != nil {
		return
	}

	if !found {
		err = ErrInexistentAccount
		return
	}

	userId, err := getUserId(db, params.Email)
	if err != nil {
		return
	}

	password, err := getDefaultAccountPassword(db, userId)
	if err != nil {
		return
	}

	if !crypto.CheckPasswordHash(params.Password, password) {
		err = ErrPasswordsDontMatch
		return
	}

	sess, err = session.GetSession(ctx, userId)
	if err == redis.Nil {
		return session.NewSession(ctx, userId)
	}

	return
}
