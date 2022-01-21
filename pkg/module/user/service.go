package user

import (
	"context"
	"errors"
)

var (
	ErrNotFound          = errors.New("user not found")
	ErrDuplicateUserName = errors.New("duplicate username")
	ErrDuplicateEmail    = errors.New("duplicate email")
)

type Service interface {
	GetUsers(ctx context.Context, pageIndex int, pageSize int) (*UserPage, error)
	GetUserById(ctx context.Context, id string) (*User, error)
	CreateUser(ctx context.Context, user User) (*User, error)
	UpdateUser(ctx context.Context, id string, user User) (*User, error)
	DeleteUser(ctx context.Context, id string) error
}

type Repository interface {
	GetUsers(ctx context.Context, pageIndex int, pageSize int) (*UserPage, error)
	GetUserById(ctx context.Context, id string) (*User, error)
	GetUserByUserName(ctx context.Context, userName string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, user User) (string, error)
	UpdateUser(ctx context.Context, id string, user User) error
	DeleteUser(ctx context.Context, id string) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (svc *service) GetUsers(ctx context.Context, pageIndex int, pageSize int) (*UserPage, error) {
	return svc.repo.GetUsers(ctx, pageIndex, pageSize)
}

func (svc *service) GetUserById(ctx context.Context, id string) (*User, error) {
	result, err := svc.repo.GetUserById(ctx, id)
	if err != nil {
		return nil, ErrNotFound
	}
	return result, nil
}

func (svc *service) CreateUser(ctx context.Context, user User) (*User, error) {
	var err error
	_, err = svc.repo.GetUserByUserName(ctx, user.UserName)
	if err == nil {
		return nil, ErrDuplicateUserName
	}
	_, err = svc.repo.GetUserByEmail(ctx, user.Email)
	if err == nil {
		return nil, ErrDuplicateEmail
	}

	newid, err := svc.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return svc.GetUserById(ctx, newid)
}

func (svc *service) UpdateUser(ctx context.Context, id string, user User) (*User, error) {
	_, err := svc.repo.GetUserById(ctx, id)
	if err != nil {
		return nil, ErrNotFound
	}

	err = svc.repo.UpdateUser(ctx, id, user)
	if err != nil {
		return nil, err
	}
	return svc.GetUserById(ctx, id)
}

func (svc *service) DeleteUser(ctx context.Context, id string) error {
	_, err := svc.repo.GetUserById(ctx, id)
	if err != nil {
		return ErrNotFound
	}
	return svc.repo.DeleteUser(ctx, id)
}
