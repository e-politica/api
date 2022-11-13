package repository

import (
	"context"
	"errors"
	"net/http"

	"github.com/e-politica/api/pkg/database"
	"github.com/e-politica/api/pkg/session"
)

var (
	ErrAlreadySigned      = errors.New("user already signed for this politician")
	ErrPoliticianNotFound = errors.New("politician not found")
)

func Follow(ctx context.Context, db *database.Db, accessToken, politId string) (err error) {
	ok, err := politicianExists(politId)
	if err != nil {
		return
	}
	if !ok {
		err = ErrPoliticianNotFound
		return
	}

	userId, err := session.GetUserId(ctx, accessToken)
	if err != nil {
		return
	}

	query := `
	SELECT COUNT(*)
	FROM user_politician
	WHERE up_user_id=$1
	AND up_politician_id=$2
	`

	row, err := db.QueryRow(query, userId, politId)
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
	INSERT INTO user_politician (
		up_user_id,
		up_politician_id
	)
	VALUES ($1, $2)
	`

	_, err = db.Exec(query, userId, politId)
	return
}

func politicianExists(id string) (ok bool, err error) {
	resp, err := http.Get("https://dadosabertos.camara.leg.br/api/v2/deputados/" + id)
	if err != nil {
		return
	}

	ok = resp.StatusCode == http.StatusOK
	return
}
