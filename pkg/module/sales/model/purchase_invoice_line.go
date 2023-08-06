package model

import (
	global "github.com/deb-ict/cloudbm-community/pkg/module/global/model"
	"github.com/shopspring/decimal"
)

type PurchaseInvoiceLine struct {
	Id           string
	ProductId    string
	ProductCode  string
	SupplierCode string
	Description  string
	Quanity      decimal.Decimal
	UnitPrice    decimal.Decimal
	TaxProfile   *global.TaxProfile
}
