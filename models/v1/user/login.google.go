package user

import "errors"

type LoginGoogleParams struct {
	Credential string `form:"credential"`
	CsrfToken  string `form:"g_csrf_token"`
}

func (l LoginGoogleParams) Validate(csrfCookie []byte) error {
	if l.Credential == "" {
		return errors.New("missing 'credential' field in body")
	}
	if l.CsrfToken == "" {
		return errors.New("missing 'g_csrf_token' field in body")
	}

	if l.CsrfToken != string(csrfCookie) {
		return errors.New("failed to verify double submit 'g_csrf_token' cookie")
	}

	return nil
}
