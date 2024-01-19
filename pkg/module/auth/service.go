package auth

import (
	"context"
	"time"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth/model"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth/security"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth/util"
)

type Service interface {
	UserNormalizer() util.UserNormalizer
	PasswordHasher() security.PasswordHasher
	FeatureProvider() core.FeatureProvider

	GetUsers(ctx context.Context, offset int64, limit int64, filter *model.UserFilter, sort *core.Sort) ([]*model.User, int64, error)
	GetUserById(ctx context.Context, id string) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, model *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, id string, model *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, id string) error

	LockUser(ctx context.Context, user *model.User, duration time.Duration) (*model.User, error)
	UnlockUser(ctx context.Context, user *model.User) (*model.User, error)
	VerifyPassword(ctx context.Context, user *model.User, password string) error
}
