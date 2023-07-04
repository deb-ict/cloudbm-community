package auth

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth/model"
)

type Service interface {
	GetUsers(ctx context.Context, offset int64, limit int64, filter *model.UserFilter, sort *core.Sort) ([]*model.User, int64, error)
	GetUserById(ctx context.Context, id string) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, id string, user *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, id string) error

	GenerateActivationToken(ctx context.Context, user *model.User) (string, error)
	GeneratePasswordResetToken(ctx context.Context, user *model.User) (string, error)

	ActivateUser(ctx context.Context, user *model.User, token string) error
	ResetPassword(ctx context.Context, user *model.User, token string, password string) error
	ChangePassword(ctx context.Context, user *model.User, oldPassword string, newPassword string) error
	ValidatePassword(ctx context.Context, user *model.User, password string) bool
}
