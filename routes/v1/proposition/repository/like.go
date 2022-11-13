package repository

import (
	"context"
	"errors"

	"github.com/e-politica/api/pkg/database"
	"github.com/e-politica/api/pkg/session"
)

var ErrAlreadySigned = errors.New("user already signed for this proposition")

func Like(ctx context.Context, db *database.Db, accessToken, propId string) (err error) {
	ok, err := propositionExists(propId)
	if err != nil {
		return
	}
	if !ok {
		err = ErrPropositionNotFound
		return
	}

	userId, err := session.GetUserId(ctx, accessToken)
	if err != nil {
		return
	}

	query := `
	SELECT COUNT(*)
	FROM user_proposition
	WHERE up_user_id=$1
	AND up_proposition_id=$2
	`

	row, err := db.QueryRow(query, userId, propId)
	if err != nil {
		return
	}

	var count int
	if err = row.Scan(&count); err != nil {
		return
	}

	if count != 0 {
		err = ErrAlreadySigned
		return
	}

	query = `
	INSERT INTO user_proposition (
		up_user_id,
		up_proposition_id
	)
	VALUES ($1, $2)
	`

	_, err = db.Exec(query, userId, propId)
	return
}
