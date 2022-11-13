package user

import (
	"errors"
)

type LoginDefaultParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r LoginDefaultParams) Validate() error {
	if r.Email == "" {
		return errors.New("field 'email' must be provided")
	}
	if r.Password == "" {
		return errors.New("field 'password' must be provided")
	}

	return nil
}
