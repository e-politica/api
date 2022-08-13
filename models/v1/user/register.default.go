package user

import (
	"errors"
	"net/mail"
)

type RegisterDefault struct {
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Picture  *string `json:"picture"`
}

func (r RegisterDefault) Validate() error {
	if r.Name == nil {
		return errors.New("invalid value on 'name' field")
	}
	if r.Email == nil {
		return errors.New("invalid value on 'email' field")
	}
	if r.Password == nil {
		return errors.New("invalid value on 'password' field")
	}
	if r.Picture == nil {
		return errors.New("invalid value on 'picture' field")
	}

	if _, err := mail.ParseAddress(*r.Email); err != nil {
		return errors.New("invalid email")
	}

	return nil
}
