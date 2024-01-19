package auth

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth/model"
)

type Database interface {
	Users() UserRepository
	UserTokens() UserTokenRepository
}

type UserRepository interface {
	GetUsers(ctx context.Context, offset int64, limit int64, filter *model.UserFilter, sort *core.Sort) ([]*model.User, int64, error)
	GetUserById(ctx context.Context, id string) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (string, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, user *model.User) error
}

type UserTokenRepository interface {
	CreateUserToken(ctx context.Context, user *model.User, userToken *model.UserToken) (string, error)
	DeleteUserToken(ctx context.Context, user *model.User, userToken *model.UserToken) error
}
