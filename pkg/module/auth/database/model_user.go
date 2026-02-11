package database

import (
	"time"

	"github.com/deb-ict/cloudbm-community/pkg/module/auth/model"
	"gorm.io/gorm"
)

type UserEntity struct {
	Id                 string         `gorm:"type:varchar(50);primaryKey"`
	NormalizedUsername string         `gorm:"type:varchar(100);unique;index:idx_normalized_username"`
	NormalizedEmail    string         `gorm:"type:varchar(250);unique;index:idx_normalized_email"`
	Username           string         `gorm:"type:varchar(100)"`
	PasswordHash       string         `gorm:"type:varchar(100)"`
	Email              string         `gorm:"type:varchar(250)"`
	EmailVerified      bool           `gorm:"type:boolean"`
	Phone              string         `gorm:"type:varchar(20)"`
	PhoneVerified      bool           `gorm:"type:boolean"`
	LoginFailures      int32          `gorm:"type:int"`
	IsLocked           bool           `gorm:"type:boolean"`
	IsEnabled          bool           `gorm:"type:boolean"`
	LockEnd            time.Time      `gorm:"type:datetime"`
	CreatedAt          time.Time      `gorm:"autoCreateTime"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}

func UserToDomainModel(user *UserEntity) *model.User {
	if user == nil {
		return nil
	}

	model := &model.User{
		Id:                 user.Id,
		NormalizedUsername: user.NormalizedUsername,
		NormalizedEmail:    user.NormalizedEmail,
		Username:           user.Username,
		PasswordHash:       user.PasswordHash,
		Email:              user.Email,
		EmailVerified:      user.EmailVerified,
		Phone:              user.Phone,
		PhoneVerified:      user.PhoneVerified,
		LoginFailures:      user.LoginFailures,
		IsLocked:           user.IsLocked,
		IsEnabled:          user.IsEnabled,
		LockEnd:            user.LockEnd,
	}
	return model
}

func UserFromDomainModel(model *model.User) *UserEntity {
	if model == nil {
		return nil
	}

	user := &UserEntity{
		Id:                 model.Id,
		NormalizedUsername: model.NormalizedUsername,
		NormalizedEmail:    model.NormalizedEmail,
		Username:           model.Username,
		PasswordHash:       model.PasswordHash,
		Email:              model.Email,
		EmailVerified:      model.EmailVerified,
		Phone:              model.Phone,
		PhoneVerified:      model.PhoneVerified,
		IsLocked:           model.IsLocked,
		IsEnabled:          model.IsEnabled,
		LockEnd:            model.LockEnd,
	}
	return user
}
