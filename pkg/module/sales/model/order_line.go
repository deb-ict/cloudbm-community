package model

type OrderLine struct {
	Id          string
	ArticleType ArticleType
	ArticleId   string
	Description string
	Quantity    int64
}
