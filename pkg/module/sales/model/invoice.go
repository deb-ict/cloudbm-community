package model

type Invoice struct {
	Id    string
	Lines []*InvoiceLine
}
