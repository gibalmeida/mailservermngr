package port

import (
	"context"

	"github.com/gibalmeida/mailservermngr/internal/domain"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, newUser domain.User) error
	DeleteUser(ctx context.Context, username string) error
	GetUser(ctx context.Context, username string) (domain.User, error)
}

type AuthUseCase interface {
	GetTokenUsingUsernameAndPassword(ctx context.Context, username, password string) (string, error)
}
