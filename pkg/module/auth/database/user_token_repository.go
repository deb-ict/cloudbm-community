package database

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/module/auth/model"
	"gorm.io/gorm"
)

type userTokenRepository struct {
	db *gorm.DB
}

func (r *userTokenRepository) CreateUserToken(ctx context.Context, user *model.User, userToken *model.UserToken) (string, error) {
	// Implementation here
	return "", nil
}

func (r *userTokenRepository) DeleteUserToken(ctx context.Context, user *model.User, userToken *model.UserToken) error {
	// Implementation here
	return nil
}
