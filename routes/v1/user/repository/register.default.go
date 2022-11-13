package repository

import (
	"context"
	"errors"

	"github.com/e-politica/api/models/v1/user"
	"github.com/e-politica/api/pkg/crypto"
	"github.com/e-politica/api/pkg/database"
	"github.com/e-politica/api/pkg/session"
)

var (
	ErrExistentAccount = errors.New("the requested account already exists")
)

func RegisterDefault(ctx context.Context, db *database.Db, params user.RegisterDefaultParams) (sess session.Session, err error) {
	found, err := isUserRegistered(db, params.Email)
	if err != nil {
		return
	}

	if found {
		err = ErrExistentAccount
		return
	}

	tx, err := db.Conn.Begin(*db.Ctx)
	if err != nil {
		return
	}
	defer tx.Rollback(*db.Ctx)

	userId, err := insertUserTx(
		*db.Ctx,
		tx,
		params.Name,
		params.Email,
		params.Picture,
		"default",
	)
	if err != nil {
		return
	}

	password, err := crypto.HashPassword(params.Password)
	if err != nil {
		return
	}

	err = insertDefaultAccountTx(
		*db.Ctx,
		tx,
		userId,
		password,
	)
	if err != nil {
		return
	}

	tx.Commit(*db.Ctx)

	return session.NewSession(ctx, userId)
}
