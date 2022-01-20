package user

import (
	"context"
	"errors"
)

var (
	ErrorNotFound = errors.New("user not found")
)

type Service interface {
	GetUsers(ctx context.Context, pageIndex int, pageSize int) (UserPage, error)
	GetUserById(ctx context.Context, id string) (User, error)
}

type Repository interface {
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (svc *service) GetUsers(ctx context.Context, pageIndex int, pageSize int) (UserPage, error) {
	return UserPage{
		PageIndex: 1,
		PageSize:  25,
		Count:     2,
		Data: []User{
			{Id: "1", UserName: "admin"},
			{Id: "2", UserName: "dev"},
		},
	}, nil
}

func (svc *service) GetUserById(ctx context.Context, id string) (User, error) {
	if id == "1" {
		return User{Id: "1", UserName: "admin"}, nil
	}
	if id == "2" {
		return User{Id: "2", UserName: "dev"}, nil
	}
	return User{}, ErrorNotFound
}
