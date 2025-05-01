package model

type ProductType int

const (
	ProductTypeStandard ProductType = iota
	ProductTypeTemplate
	ProductTypeVariant
	ProductTypeBundle
	ProductTypeUnknown
)

func (m ProductType) String() string {
	switch m {
	case ProductTypeStandard:
		return "Standard"
	case ProductTypeTemplate:
		return "Template"
	case ProductTypeVariant:
		return "Variant"
	case ProductTypeBundle:
		return "Bundle"
	default:
		return "Unknown"
	}
}

func (m ProductType) IsStandard() bool {
	return m == ProductTypeStandard
}

func (m ProductType) IsTemplate() bool {
	return m == ProductTypeTemplate
}

func (m ProductType) IsVariant() bool {
	return m == ProductTypeVariant
}

func (m ProductType) IsBundle() bool {
	return m == ProductTypeBundle
}

func (m ProductType) IsValid() bool {
	return m.IsStandard() || m.IsTemplate() || m.IsVariant() || m.IsBundle()
}

func ParseProductType(value string) ProductType {
	switch value {
	case "Standard":
		return ProductTypeStandard
	case "Template":
		return ProductTypeTemplate
	case "Variant":
		return ProductTypeVariant
	case "Bundle":
		return ProductTypeBundle
	default:
		return ProductTypeUnknown
	}
}
