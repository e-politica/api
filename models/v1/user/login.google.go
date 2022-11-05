package user

import "errors"

type LoginGoogle struct {
	Credential *string `form:"credential"`
	CsrfToken  *string `form:"g_csrf_token"`
}

func (l LoginGoogle) Validate(csrfCookie []byte) error {
	if l.Credential == nil {
		return errors.New("missing 'credential' field in body")
	}
	if l.CsrfToken == nil {
		return errors.New("missing 'g_csrf_token' field in body")
	}

	if *l.CsrfToken != string(csrfCookie) {
		return errors.New("failed to verify double submit 'g_csrf_token' cookie")
	}

	return nil
}
