package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth/model"
	"github.com/sirupsen/logrus"
)

func (svc *service) GetUsers(ctx context.Context, offset int64, limit int64, filter *model.UserFilter, sort *core.Sort) ([]*model.User, int64, error) {
	data, count, err := svc.database.Users().GetUsers(ctx, offset, limit, filter, sort)
	if err != nil {
		logrus.Errorf("Failed to get users from database: %v", err)
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetUserById(ctx context.Context, id string) (*model.User, error) {
	data, err := svc.database.Users().GetUserById(ctx, id)
	if err == nil && data == nil {
		err = auth.ErrUserNotFound
	}
	if err != nil {
		logrus.Errorf("Failed to get user with id '%s' from database: %v", id, err)
		return nil, err
	}

	return data, nil
}

func (svc *service) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	normalizedUsername := svc.userNormalizer.NormalizeUsername(username)
	data, err := svc.database.Users().GetUserByUsername(ctx, normalizedUsername)
	if err == nil && data == nil {
		err = auth.ErrUserNotFound
	}
	if err != nil {
		logrus.Errorf("Failed to get user with username '%s' from database: %v", username, err)
		return nil, err
	}

	return data, nil
}

func (svc *service) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	normalizedEmail := svc.userNormalizer.NormalizeEmail(email)
	data, err := svc.database.Users().GetUserByEmail(ctx, normalizedEmail)
	if err == nil && data == nil {
		err = auth.ErrUserNotFound
	}
	if err != nil {
		logrus.Errorf("Failed to get user with email '%s' from database: %v", email, err)
		return nil, err
	}

	return data, nil
}

func (svc *service) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	user.NormalizedUsername = svc.userNormalizer.NormalizeUsername(user.Username)
	user.NormalizedEmail = svc.userNormalizer.NormalizeEmail(user.Email)

	if err := svc.checkDuplicateUsername(ctx, user); err != nil {
		return nil, err
	}
	if err := svc.checkDuplicateEmail(ctx, user); err != nil {
		return nil, err
	}

	newId, err := svc.database.Users().CreateUser(ctx, user)
	if err != nil {
		logrus.Errorf("Failed to create user in database: %v", err)
		return nil, err
	}

	if !user.EmailVerified {
		//TODO: Send activation email
	}

	return svc.GetUserById(ctx, newId)
}

func (svc *service) UpdateUser(ctx context.Context, id string, user *model.User) (*model.User, error) {
	data, err := svc.database.Users().GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, auth.ErrUserNotFound
	}

	data.EmailVerified = user.EmailVerified
	data.Phone = user.Phone
	data.PhoneVerified = user.PhoneVerified
	data.LoginFailures = user.LoginFailures
	data.IsEnabled = user.IsEnabled
	data.IsLocked = user.IsLocked
	data.LockEnd = user.LockEnd

	return nil, core.ErrNotImplemented
}

func (svc *service) DeleteUser(ctx context.Context, id string) error {
	data, err := svc.database.Users().GetUserById(ctx, id)
	if err != nil {
		return err
	}
	if data == nil {
		return auth.ErrUserNotFound
	}

	err = svc.database.Users().DeleteUser(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (svc *service) checkDuplicateUsername(ctx context.Context, user *model.User) error {
	existing, err := svc.database.Users().GetUserByUsername(ctx, user.NormalizedUsername)
	if err != nil {
		return err
	}
	if existing != nil {
		return auth.ErrDuplicateUsername
	}
	return nil
}

func (svc *service) checkDuplicateEmail(ctx context.Context, user *model.User) error {
	existing, err := svc.database.Users().GetUserByEmail(ctx, user.NormalizedEmail)
	if err != nil {
		return err
	}
	if existing != nil {
		return auth.ErrDuplicateEmail
	}
	return nil
}
