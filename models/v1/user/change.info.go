package user

import (
	"errors"
)

type ChangeInfoParams struct {
	Name        string `json:"name"`
	Picture     string `json:"picture"`
	NewPassword string `json:"new_password"`
	Password    string `json:"password"`
}

func (r ChangeInfoParams) Validate() error {
	if r.Name == "" && r.Picture == "" && r.NewPassword == "" {
		return errors.New("must provide at least 1 field")
	}

	if r.Password == "" {
		return errors.New("field 'password' is required")
	}

	return nil
}
