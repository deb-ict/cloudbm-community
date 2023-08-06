package model

import (
	"time"
)

type SalesOrder struct {
	Id                  string
	CustomerType        CustomerType
	CustomerId          string
	CustomerName        string
	OrderNumber         string
	CustomerOrderNumber string
	OrderDate           time.Time
	DueDate             time.Time
	InvoiceAddress      *Address
	DeliveryAddress     *Address
	Lines               []*SalesOrderLine
}
