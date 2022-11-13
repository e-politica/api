package repository

import (
	"context"
	"errors"
	"time"

	"github.com/e-politica/api/models/v1/user"
	"github.com/e-politica/api/pkg/crypto"
	"github.com/e-politica/api/pkg/database"
	"github.com/e-politica/api/pkg/session"
)

var ErrNotDefaultAccount = errors.New("not a default account")

func ChangeInfo(ctx context.Context, db *database.Db, access string, params user.ChangeInfoParams) (err error) {
	userId, err := session.GetUserId(ctx, access)
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

	query := `
	UPDATE user_account
	SET	
		user_name=(CASE WHEN LENGTH($2) = 0 THEN user_name ELSE $2 END),
		user_picture_url=(CASE WHEN LENGTH($3) = 0 THEN user_picture_url ELSE $3 END),
		user_update_timestamp=$4
	WHERE user_id=$1
	`

	_, err = db.Exec(query, userId, params.Name, params.Picture, time.Now())
	if err != nil {
		return
	}

	if params.NewPassword == "" {
		return
	}

	query = `
	SELECT user_account_type
	FROM user_account
	WHERE user_id=$1
	`

	row, err := db.QueryRow(query, userId)
	if err != nil {
		return
	}

	var accountType string
	if err = row.Scan(&accountType); err != nil {
		return
	}

	if accountType != "default" {
		err = ErrNotDefaultAccount
		return
	}

	newPassword, err := crypto.HashPassword(params.NewPassword)
	if err != nil {
		return
	}

	query = `
	UPDATE default_account
	SET 
		df_account_password=$2,
		df_account_update_timestamp=$3
	WHERE df_account_user_id=$1
	`

	_, err = db.Exec(query, userId, newPassword, time.Now())
	return
}
