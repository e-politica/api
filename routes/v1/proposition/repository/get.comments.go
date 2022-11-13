package repository

import (
	"context"

	"github.com/e-politica/api/pkg/database"
)

func GetComments(ctx context.Context, db *database.Db, propId string, offset, limit int) (comments []string, err error) {
	query := `
	SELECT cp_comment
	FROM comment_proposition
	WHERE cp_proposition_id=$1
	ORDER BY cp_create_timestamp ASC
	OFFSET $2
	LIMIT $3
	`

	rows, err := db.Query(query, propId, offset, limit)
	if err != nil {
		return
	}

	var comment string
	for rows.Next() {
		if err = rows.Scan(&comment); err != nil {
			return
		}
		comments = append(comments, comment)
	}

	return
}
