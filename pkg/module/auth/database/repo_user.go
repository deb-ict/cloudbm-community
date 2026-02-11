package database

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (UserEntity) TableName() string {
	return "auth_user"
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) GetUsers(ctx context.Context, offset int64, limit int64, filter *model.UserFilter, sort *core.Sort) ([]*model.User, int64, error) {
	query := gorm.G[UserEntity](r.db).Offset(int(offset)).Limit(int(limit))
	if filter != nil {
		if filter.Username != "" {
			query = query.Where("normalized_username LIKE ?", "%"+filter.Username+"%")
		}
		if filter.Email != "" {
			query = query.Where("normalized_email LIKE ?", "%"+filter.Email+"%")
		}
	}
	if sort != nil {
		for _, field := range sort.Fields {
			order := "asc"
			if field.Order == core.SortDescending {
				order = "desc"
			}
			query = query.Order(field.Name + " " + order)
		}
	}

	count, err := query.Count(ctx, "id")
	if err != nil {
		return nil, 0, err
	}

	entities, err := query.Find(ctx)
	if err != nil {
		return nil, 0, err
	}

	result := make([]*model.User, len(entities))
	for i, entity := range entities {
		result[i] = UserToDomainModel(&entity)
	}

	return result, count, nil
}

func (r *userRepository) GetUserById(ctx context.Context, id string) (*model.User, error) {
	entity, err := gorm.G[UserEntity](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}
	return UserToDomainModel(&entity), nil
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	entity, err := gorm.G[UserEntity](r.db).Where("normalized_username = ?", username).First(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return UserToDomainModel(&entity), nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	entity, err := gorm.G[UserEntity](r.db).Where("normalized_email = ?", email).First(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return UserToDomainModel(&entity), nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) (string, error) {
	user.Id = uuid.NewString()

	err := gorm.G[UserEntity](r.db).Create(ctx, UserFromDomainModel(user))
	if err != nil {
		return "", err
	}
	return user.Id, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *model.User) error {
	rows, err := gorm.G[UserEntity](r.db).Where("id = ?", user.Id).Updates(ctx, *UserFromDomainModel(user))
	if err != nil {
		return err
	}
	if rows == 0 {
		return core.ErrRecordNotChanged
	}
	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, user *model.User) error {
	rows, err := gorm.G[UserEntity](r.db).Where("id = ?", user.Id).Delete(ctx)
	if err != nil {
		return err
	}
	if rows == 0 {
		return core.ErrRecordNotDeleted
	}
	return nil
}
