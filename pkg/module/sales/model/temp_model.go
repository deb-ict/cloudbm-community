package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type ShippingRate struct {
	Price       decimal.Decimal
	Description string
	Min         int32
	MinType     uint8 // Hours, Days, BusinessDays, Weeks, Months
	Max         int32
	MaxType     uint8
	TaxCode     uint8 // Shipping, Non-Taxable
}

type StripeProduct struct {
	Name         string
	Description  string
	TaxCode      string
	Price        int64
	PackagePrice int64 // Package pricing model
	PackageUnits int64
}

type VolumePrice struct {
	MinUnit int64
	MaxUnit int64
	Price   int64
}

type WooProduct struct {
	Name        string
	Description string
}

type WooSimpleProduct struct {
	// General
	RegularPrice     decimal.Decimal
	SalePrice        decimal.Decimal
	SalePriceStart   time.Time
	SalePriceEnd     time.Time
	TaxStatus        uint8  // Taxable, shipping only, none
	TaxClass         uint8  // ref table?
	StrikePriceLabel string // ref table
	SalePriceLabel   string // ref table
	Unit             string // ref table
	UnitQuantity     int64  // number of unit in default product price
	GTIN             string
	MPN              string
	// Inventory
	SKU             string
	StockManagement bool
	StockStatus     int8 // InStock,OutOfStock,BackOrder
	// Shipping
	Weight                 int
	Length                 int
	Width                  int
	Height                 int
	ShippingClass          int8              // ref table
	HSCode                 string            // customs
	CountryOfManufacture   string            // ref table
	DeliveryTime           string            // ref table
	DeliveryTimePerCountry map[string]string // ref table?
	FreeShipping           bool
	// Linked products
	// Attributes
}

type WooGroupedProduct struct {
	SKU  string
	GTIN string
}

type WooDeliveryTime struct {
	Name        string
	Slug        string
	Description string
}

type WooUnit struct {
	Name        string
	Slug        string
	Description string
}

type WooPriceLabel struct {
	Name        string
	Slug        string
	Description string
}

type WooDepositType struct {
	Name        string
	Slug        string
	Description string
	Price       int64
	PackageType uint8 // none, reusable, disposable
}

/*
WooConfig
	Store address
	SellType => All/AllExcept/Specific
	SellCounty[]
	ShippingType => AllSell/All/Specific/Disabled
	ShippingCountry[]
	EnableTaxes (enable tax rates and calculations)
	Currency
	CurrencyPosition => Left/Right/LeftWithSpace/RightWithSpace
	ThousandSeparator
	DecimalSeparator
	NumberOfDecimals
	//Product:General
	WeightUnit (kg)
	DimensionUnit (cm)
	EnableReviews
	EnableRating
	//Product:Inventory
	EnabledStock
	HoldStock (reserve stock for unpaid order for x minutes, when limit reached order cancelled, 0 = disable)
	EnableLowStockNotification
	LowStockThreshold
	EnableOutOfStockNotification
	OutOfStockThreshold
	//Tax:Options
	PricesEnteredWithTax yes/no
	CalculateTaxBasedOn => ShippingAddress/BillingAddress
	ShippingTaxClass => ShippingTaxBasedOnCartItems or ref table?
	AdditionalTaxClasses []list of string
	DisplayPriceInShop => TaxIn/TaxEx
	DisplayPriceInCart => TaxIn/TaxEx
	Tax:Standard rate
		CountryCode
		StateCode
		PostalCode
		City
		Rate
		TaxName
		Compound bool
		Shipping bool
	Shipping:zones
		table
			name
			regions
			methods[]
				Free shipping
				Standard shipping
				Local pickup
	Shipping:classes
		table
			name
			slug
			description
*/
