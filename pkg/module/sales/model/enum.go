package model

type (
	ArticleType   uint8
	AllowanceType uint8
	ChargeType    uint8
)

const (
	ArticleType_Undefined ArticleType = iota
	ArticleType_Product
	ArticleType_Project
	ArticleType_Service
)

const (
	AllowanceType_Undefined AllowanceType = iota
	AllowanceType_Fixed
	AllowanceType_Factor
)

const (
	ChargeType_Undefined ChargeType = iota
	ChargeType_Fixed
	ChargeType_Factor
)

func (t ArticleType) String() string {
	switch t {
	case ArticleType_Product:
		return "Product"
	case ArticleType_Project:
		return "Project"
	case ArticleType_Service:
		return "Service"
	default:
		return "Undefined"
	}
}

func (e ChargeType) String() string {
	switch e {
	case ChargeType_Fixed:
		return "Fixed"
	case ChargeType_Factor:
		return "Factor"
	default:
		return "Undefined"
	}
}

func (e AllowanceType) String() string {
	switch e {
	case AllowanceType_Fixed:
		return "Fixed"
	case AllowanceType_Factor:
		return "Factor"
	default:
		return "Undefined"
	}
}
