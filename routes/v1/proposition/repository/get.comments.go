package repository

import (
	"context"

	"github.com/e-politica/api/models/v1/proposition"
	"github.com/e-politica/api/models/v1/user"
	"github.com/e-politica/api/pkg/database"
	userrepo "github.com/e-politica/api/routes/v1/user/repository"
)

func GetComments(ctx context.Context, db *database.Db, propId string, offset, limit int) (commentsInfo []proposition.CommentInfo, err error) {
	query := `
	SELECT 
		cp_user_id,
		cp_comment
	FROM comment_proposition
	WHERE cp_proposition_id=$1
	ORDER BY cp_create_timestamp DESC
	OFFSET $2
	LIMIT $3
	`

	rows, err := db.Query(query, propId, offset, limit)
	if err != nil {
		return
	}

	var userId, comment string
	var userInfo user.PublicInfo
	for rows.Next() {
		if err = rows.Scan(&userId, &comment); err != nil {
			return
		}

		userInfo, err = userrepo.GetPublicInfo(ctx, db, userId)
		if err != nil {
			return
		}

		commentsInfo = append(commentsInfo, proposition.CommentInfo{
			Comment:     comment,
			UserPubInfo: userInfo,
		})
	}

	return
}
