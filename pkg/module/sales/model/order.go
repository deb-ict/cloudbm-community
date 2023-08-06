package model

type Order struct {
	Id    string
	Lines []*OrderLine
}
