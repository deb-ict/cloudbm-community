package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/logging"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth/model"
)

func (svc *service) GetUsers(ctx context.Context, offset int64, limit int64, filter *model.UserFilter, sort *core.Sort) ([]*model.User, int64, error) {
	data, count, err := svc.database.Users().GetUsers(ctx, offset, limit, filter, sort)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get users from database",
			slog.Any("error", err),
		)
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetUserById(ctx context.Context, id string) (*model.User, error) {
	data, err := svc.database.Users().GetUserById(ctx, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get user from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, auth.ErrUserNotFound
	}

	return data, nil
}

func (svc *service) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	normalizedUsername := svc.userNormalizer.NormalizeUsername(username)
	data, err := svc.database.Users().GetUserByUsername(ctx, normalizedUsername)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get user from database by username",
			slog.String("username", username),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, auth.ErrUserNotFound
	}

	return data, nil
}

func (svc *service) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	normalizedEmail := svc.userNormalizer.NormalizeEmail(email)
	data, err := svc.database.Users().GetUserByEmail(ctx, normalizedEmail)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get user from database by email",
			slog.String("email", email),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, auth.ErrUserNotFound
	}

	return data, nil
}

func (svc *service) CreateUser(ctx context.Context, model *model.User) (*model.User, error) {
	model.Normalize(svc.userNormalizer)
	model.Id = ""

	if err := svc.checkDuplicateUsername(ctx, model); err != nil {
		slog.WarnContext(ctx, "Failed to create user cause of duplicate username",
			slog.String("username", model.Username),
			slog.Any("error", err),
		)
		return nil, err
	}
	if err := svc.checkDuplicateEmail(ctx, model); err != nil {
		slog.WarnContext(ctx, "Failed to create user cause of duplicate email",
			slog.String("email", model.Email),
			slog.Any("error", err),
		)
		return nil, err
	}

	newId, err := svc.database.Users().CreateUser(ctx, model)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to create user in database",
			slog.Any("error", err),
		)
		return nil, err
	}

	return svc.GetUserById(ctx, newId)
}

func (svc *service) UpdateUser(ctx context.Context, id string, model *model.User) (*model.User, error) {
	model.Normalize(svc.userNormalizer)
	model.Id = id

	data, err := svc.database.Users().GetUserById(ctx, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get user from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}
	if data == nil {
		return nil, auth.ErrUserNotFound
	}
	data.UpdateModel(model)

	err = svc.database.Users().UpdateUser(ctx, model)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to update user in database",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return nil, err
	}

	return svc.GetUserById(ctx, id)
}

func (svc *service) DeleteUser(ctx context.Context, id string) error {
	data, err := svc.database.Users().GetUserById(ctx, id)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to get user from database by id",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}
	if data == nil {
		return auth.ErrUserNotFound
	}

	err = svc.database.Users().DeleteUser(ctx, data)
	if err != nil {
		logging.GetLoggerFromContext(ctx).ErrorContext(ctx, "Failed to delete user in database",
			slog.String("id", id),
			slog.Any("error", err),
		)
		return err
	}

	return nil
}

func (svc *service) LockUser(ctx context.Context, user *model.User, duration time.Duration) (*model.User, error) {
	user.Lock(duration)

	err := svc.database.Users().UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return svc.GetUserById(ctx, user.Id)
}

func (svc *service) UnlockUser(ctx context.Context, user *model.User) (*model.User, error) {
	user.Unlock()

	err := svc.database.Users().UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return svc.GetUserById(ctx, user.Id)
}

func (svc *service) VerifyPassword(ctx context.Context, user *model.User, password string) error {
	if user.VerifyPassword(svc.passwordHasher, password) {
		return nil
	}
	return auth.ErrPasswordNotMatch
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
