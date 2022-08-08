package repository

import (
	"context"

	goauth2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func Verify(ctx context.Context, code string) error {
	token, err := oauthConfig.Exchange(ctx, code)
	if err != nil {
		return err
	}

	tokenSrc := oauthConfig.TokenSource(ctx, token)
	oauth2Service, err := goauth2.NewService(
		ctx,
		option.WithTokenSource(tokenSrc),
	)
	if err != nil {
		return err
	}

	_, err = oauth2Service.Userinfo.V2.Me.Get().Do()
	if err != nil {
		return err
	}

	return nil
}
