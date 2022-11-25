package user

import "errors"

type LoginGoogleParams struct {
	Credential string `form:"credential"`
	CsrfToken  string `form:"g_csrf_token"`
}

func (l LoginGoogleParams) Validate() error {
	if l.Credential == "" {
		return errors.New("missing 'credential' field in body")
	}
	if l.CsrfToken == "" {
		return errors.New("missing 'g_csrf_token' field in body")
	}

	return nil
}
