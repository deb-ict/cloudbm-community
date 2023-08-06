package model

import (
	"time"
)

type SalesInvoice struct {
	Id                  string
	CustomerType        CustomerType
	CustomerId          string
	CustomerName        string
	InvoiceNumber       string
	OrderId             string
	OrderNumber         string
	CustomerOrderNumber string
	InvoiceDate         time.Time
	DueDate             time.Time
	InvoiceAddress      *Address
	DeliveryAddress     *Address
	Lines               []*SalesInvoiceLine
}
