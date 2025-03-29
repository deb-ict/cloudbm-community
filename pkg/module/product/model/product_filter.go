package model

type ProductFilter struct {
	Type             *ProductType
	Language         string
	Name             string
	CategoryId       string
	IncludeTemplates bool
	IncludeVariants  bool
}
