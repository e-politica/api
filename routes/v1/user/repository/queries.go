package repository

import (
	"context"

	"github.com/e-politica/api/pkg/database"
	"github.com/jackc/pgx/v4"
)

func isUserRegistered(db *database.Db, email string) (bool, error) {
	query := `
	SELECT COUNT(1)
	FROM user_account
	WHERE user_email=$1
	`

	row, err := db.QueryRow(query, email)
	if err != nil {
		return false, err
	}

	var count int
	err = row.Scan(&count)

	return count == 1, err
}

func getUserId(db *database.Db, email string) (userId string, err error) {
	query := `
	SELECT user_id
	FROM user_account
	WHERE user_email=$1
	`

	row, err := db.QueryRow(query, email)
	if err != nil {
		return
	}

	err = row.Scan(&userId)
	return
}

func getDefaultAccountPassword(db *database.Db, userId string) (password string, err error) {
	query := `
	SELECT df_account_password
	FROM default_account
	WHERE df_account_user_id=$1
	`

	row, err := db.QueryRow(query, userId)
	if err != nil {
		return
	}

	err = row.Scan(&password)
	return
}

func insertUser(db *database.Db, name, email, pictureUrl, accountType string) (userId string, err error) {
	query := `
	INSERT INTO user_account (
		user_name,
		user_email,
		user_picture_url,
		user_account_type
	)
	VALUES ($1, $2, $3, $4)
	RETURNING user_id
	`

	row, err := db.QueryRow(query, name, email, pictureUrl, accountType)
	if err != nil {
		return
	}

	err = row.Scan(&userId)
	return
}

func insertUserTx(ctx context.Context, tx pgx.Tx, name, email, pictureUrl, accountType string) (userId string, err error) {
	query := `
	INSERT INTO user_account (
		user_name,
		user_email,
		user_picture_url,
		user_account_type
	)
	VALUES ($1, $2, $3, $4)
	RETURNING user_id
	`

	row := tx.QueryRow(ctx, query, name, email, pictureUrl, accountType)

	err = row.Scan(&userId)
	return
}

func insertDefaultAccount(ctx context.Context, tx pgx.Tx, userId, password string) error {
	query := `
	INSERT INTO default_account (
		df_account_user_id,
		df_account_password
	)
	VALUES ($1, $2)
	`

	_, err := tx.Exec(ctx, query, userId, password)
	return err
}
