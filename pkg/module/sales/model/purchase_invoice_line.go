package model

import (
	metadata "github.com/deb-ict/cloudbm-community/pkg/module/metadata/model"
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
	TaxProfile   *metadata.TaxProfile
}
