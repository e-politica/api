package repository

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/e-politica/api/config"
	"github.com/e-politica/api/pkg/database"
	"github.com/e-politica/api/pkg/session"
	"github.com/jackc/pgx/v4"
)

var (
	ErrCommentDelayReached = errors.New("comment delay reached")
	ErrPropositionNotFound = errors.New("proposition not found")
)

func Comment(ctx context.Context, db *database.Db, accessToken, propId, comment string) (err error) {
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
	SELECT cp_create_timestamp
	FROM comment_proposition
	WHERE cp_user_id=$1
	ORDER BY cp_create_timestamp DESC
	`

	row, err := db.QueryRow(query, userId)
	if err != nil {
		return
	}

	var timestamp time.Time
	if err = row.Scan(&timestamp); err != nil && err != pgx.ErrNoRows {
		return
	}

	if err == nil && time.Since(timestamp) < time.Hour*config.CommentDelayHour {
		err = ErrCommentDelayReached
		return
	}

	query = `
	INSERT INTO comment_proposition (
		cp_user_id,
		cp_proposition_id,
		cp_comment
	)
	VALUES ($1, $2, $3)
	`

	_, err = db.Exec(query, userId, propId, comment)
	return
}

func propositionExists(id string) (ok bool, err error) {
	resp, err := http.Get("https://dadosabertos.camara.leg.br/api/v2/proposicoes/" + id)
	if err != nil {
		return
	}

	ok = resp.StatusCode == http.StatusOK
	return
}
