package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/gibalmeida/mailservermngr/internal/domain"
)

type memAuthRepository struct {
	Users map[string]*domain.User
	lock  *sync.RWMutex
}

func NewMemAuthRepository() *memAuthRepository {

	return &memAuthRepository{
		Users: make(map[string]*domain.User),
		lock:  &sync.RWMutex{},
	}

}

func (r memAuthRepository) CreateUser(ctx context.Context, newUser domain.User) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	index := newUser.Username
	_, exist := r.Users[index]

	if exist {
		return errors.New("User already exist.")
	}

	r.Users[index] = &newUser

	return nil
}

func (r memAuthRepository) DeleteUser(ctx context.Context, username string) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	_, exist := r.Users[username]

	if !exist {
		return errors.New("User doesn't exist.")
	}

	delete(r.Users, username)

	return nil

}

func (r memAuthRepository) GetUser(ctx context.Context, username string) (domain.User, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	user, exist := r.Users[username]

	if !exist {
		return domain.User{}, errors.New("User not found!")
	}
	return *user, nil

}
