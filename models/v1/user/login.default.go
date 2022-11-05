package user

import (
	"errors"
)

type LoginDefault struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func (r LoginDefault) Validate() error {
	if r.Email == nil {
		return errors.New("field 'email' must be provided")
	}
	if r.Password == nil {
		return errors.New("field 'password' must be provided")
	}

	return nil
}
