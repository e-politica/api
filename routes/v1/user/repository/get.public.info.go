package repository

import (
	"context"

	"github.com/e-politica/api/models/v1/user"
	"github.com/e-politica/api/pkg/database"
)

func GetPublicInfo(ctx context.Context, db *database.Db, id string) (info user.PublicInfo, err error) {
	query := `
	SELECT 
		user_name,
		user_picture_url
	FROM user_account
	WHERE user_id=$1
	`

	row, err := db.QueryRow(query, id)
	if err != nil {
		return
	}

	err = row.Scan(&info.Name, &info.Picture)
	return
}
