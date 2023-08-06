package model

type ArticleType uint8

const (
	ArticleType_Undefined ArticleType = iota
	ArticleType_Project
	ArticleType_Product
	ArticleType_Static
)

func (t ArticleType) String() string {
	switch t {
	case ArticleType_Project:
		return "Project"
	case ArticleType_Product:
		return "Product"
	case ArticleType_Static:
		return "Static"
	default:
		return "Undefined"
	}
}
