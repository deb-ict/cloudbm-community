package service

import (
	"github.com/deb-ict/cloudbm-community/pkg/module/product"
)

type service struct {
	database product.Database
}

func NewService(opts ...product.ServiceOption) product.Service {
	svc := &service{}
	for _, opt := range opts {
		opt(svc)
	}
	svc.ensureDefaults()
	return svc
}

func (svc *service) GetDatabase() product.Database {
	return svc.database
}

func (svc *service) SetDatabase(database product.Database) error {
	svc.database = database
	return nil
}

func (svc *service) ensureDefaults() {

}
