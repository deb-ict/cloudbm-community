package database

import (
	"github.com/deb-ict/cloudbm-community/pkg/module/auth"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type database struct {
	db *gorm.DB
}

func NewDatabase() (auth.Database, error) {
	//TODO: Move DSN to configuration
	dsn := "cloudbm:cloudbm@tcp(127.0.0.1:3306)/cloudbm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&UserEntity{})
	if err != nil {
		return nil, err
	}

	return &database{
		db: db,
	}, nil
}

func (d *database) Users() auth.UserRepository {
	return &userRepository{
		db: d.db,
	}
}

func (d *database) UserTokens() auth.UserTokenRepository {
	return &userTokenRepository{
		db: d.db,
	}
}
