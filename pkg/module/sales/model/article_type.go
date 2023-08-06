package model

type ArticleType uint8

const (
	ArticleType_Undefined ArticleType = iota
	ArticleType_Product
	ArticleType_Project
)

func (e ArticleType) String() string {
	switch e {
	case ArticleType_Product:
		return "Product"
	case ArticleType_Project:
		return "Project"
	default:
		return "Undefined"
	}
}
