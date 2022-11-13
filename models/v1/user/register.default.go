package user

import (
	"errors"
	"net/mail"
)

type RegisterDefaultParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Picture  string `json:"picture"`
}

func (r RegisterDefaultParams) Validate() error {
	if r.Name == "" {
		return errors.New("field 'name' must be provided")
	}
	if r.Email == "" {
		return errors.New("field 'email' must be provided")
	}
	if r.Password == "" {
		return errors.New("field 'password' must be provided")
	}
	if r.Picture == "" {
		return errors.New("field 'picture' must be provided")
	}

	if _, err := mail.ParseAddress(r.Email); err != nil {
		return errors.New("invalid email")
	}

	return nil
}
