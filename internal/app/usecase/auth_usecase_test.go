package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/gibalmeida/mailservermngr/internal/adapter/repository"
	"github.com/gibalmeida/mailservermngr/internal/app/usecase"
	"github.com/gibalmeida/mailservermngr/internal/domain"
	"github.com/gibalmeida/mailservermngr/internal/port"

	"github.com/gibalmeida/mailservermngr/pkg/jwx"
	"github.com/stretchr/testify/assert"
)

const PrivateKey = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIP2/emn5XnniWSLNHFBU7vRiI4xcDMTCf3ZGOJtsasg5oAoGCCqGSM49
AwEHoUQDQgAEZTd4NZBVoaB+Ts8FvF7/gw+lFuoSw1gMx4NYNuVQ8QiguZvdBCoR
SzxtSZxmvge+wnLZIESk6MP4TT0RfLjfQw==
-----END EC PRIVATE KEY-----`

func setupAuthTestCase(t *testing.T) (context.Context, port.AuthUseCase, *jwx.JWSAuthenticator, func(t *testing.T)) {
	t.Log("setup test case")

	ctx := context.TODO()
	repo := repository.NewMemAuthRepository()

	jwsAuth, err := jwx.NewJWSAuthenticator([]byte(PrivateKey))
	assert.NoError(t, err)

	uc := usecase.NewAuthUseCase(repo, jwsAuth)

	err = repo.CreateUser(ctx, domain.User{Username: "user", Password: "password"})
	assert.NoError(t, err)

	return ctx, uc, jwsAuth, func(t *testing.T) {
		t.Log("teardown test case")

		err := repo.DeleteUser(ctx, "user")
		assert.NoError(t, err)

	}
}

func TestGetToken(t *testing.T) {

	testCases := []struct {
		desc     string
		username string
		password string
		err      error
	}{
		{
			desc:     "success",
			username: "user",
			password: "password",
		},
		{
			desc:     "invalid credentials",
			username: "user",
			password: "wrongpassword",
			err:      errors.New("Invalid credentials"),
		},
	}

	ctx, uc, jwsAuth, teardownTestCase := setupAuthTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := uc.GetTokenUsingUsernameAndPassword(ctx, tc.username, tc.password)
			t.Log(result)
			if tc.err == nil {
				_, err := jwsAuth.ValidateJWS(result)
				assert.NoError(t, err)
			} else if assert.Error(t, err) {
				assert.Equal(t, tc.err, err)
			}

		})
	}

}
