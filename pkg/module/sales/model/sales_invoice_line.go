package model

import (
	global "github.com/deb-ict/cloudbm-community/pkg/module/global/model"
	"github.com/shopspring/decimal"
)

type SalesInvoiceLine struct {
	Id           string
	ArticleType  ArticleType
	ArticleId    string
	ArticleCode  string
	CustomerCode string
	Description  string
	Quanity      decimal.Decimal
	UnitPrice    decimal.Decimal
	TaxProfile   *global.TaxProfile
}
