package repository

import (
	"context"

	"github.com/e-politica/api/pkg/database"
	"github.com/e-politica/api/pkg/session"
)

func GetFollows(ctx context.Context, db *database.Db, access string, offset, limit int) (follows []string, err error) {
	userId, err := session.GetUserId(ctx, access)
	if err != nil {
		return
	}

	query := `
	SELECT up_politician_id
	FROM user_politician
	WHERE up_user_id=$1
	OFFSET $2
	LIMIT $3
	`

	rows, err := db.Query(query, userId, offset, limit)
	if err != nil {
		return
	}

	var politId string
	for rows.Next() {
		if err = rows.Scan(&politId); err != nil {
			return
		}
		follows = append(follows, politId)
	}

	return
}
