package usecase

import (
	"context"
	"errors"

	"github.com/gibalmeida/mailservermngr/internal/port"
	"github.com/gibalmeida/mailservermngr/pkg/jwx"
)

type authUseCase struct {
	repo             port.AuthRepository
	jwsAuthenticator *jwx.JWSAuthenticator
}

func NewAuthUseCase(repo port.AuthRepository, jwsAuth *jwx.JWSAuthenticator) port.AuthUseCase {
	return &authUseCase{
		repo:             repo,
		jwsAuthenticator: jwsAuth,
	}
}

func (u *authUseCase) GetTokenUsingUsernameAndPassword(ctx context.Context, username, password string) (string, error) {

	user, err := u.repo.GetUser(ctx, username)

	if err != nil {
		return "", err
	}

	if user.Password != password {
		return "", errors.New("Invalid credentials")
	}

	// This token allows access to API's with no scopes, and with the "mailserver:admin" claim.
	jws, err := u.jwsAuthenticator.CreateJWSWithClaims([]string{"mailserver:w"})

	if err != nil {
		return "", err
	}

	return string(jws[:]), nil // convert from a byte array to a string
}
