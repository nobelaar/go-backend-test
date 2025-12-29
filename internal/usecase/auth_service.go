package usecase

import (
	"context"
	"errors"
	"server/internal/domain"
	"server/internal/repository"
)

type AuthService struct {
	users repository.UserRepository
	hash  PasswordHasher
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}

func NewAuthService(users repository.UserRepository, hash PasswordHasher) *AuthService {
	return &AuthService{
		users: users,
		hash:  hash,
	}
}

func (a *AuthService) Register(ctx context.Context, username, password string) error {
	existing, err := a.users.FindByUsername(username)
	if err != nil && !errors.Is(err, domain.ErrUserNotFound) {
		return err
	}

	if existing != nil {
		return domain.ErrUserAlreadyExists
	}

	hash, err := a.hash.Hash(password)
	if err != nil {
		return err
	}

	user := &domain.User{
		Username: username,
		Password: hash,
	}

	return a.users.Create(user)
}

func (a *AuthService) Login(ctx context.Context, username, password string) error {
	user, err := a.users.FindByUsername(username)
	if errors.Is(err, domain.ErrUserNotFound) {
		return domain.ErrInvalidCredentials
	}
	if err != nil {
		return err
	}

	if err := a.hash.Compare(user.Password, password); err != nil {
		return domain.ErrInvalidCredentials
	}

	return nil
}
