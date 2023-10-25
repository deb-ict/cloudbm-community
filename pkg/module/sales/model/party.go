package model

type Party struct {
	Id               string
	CompanyId        string
	ContactId        string
	RecipientCompany string
	RecipientName    string
	AddressLine1     string
	AddressLine2     string
	PostalCode       string
	City             string
	State            string
	Country          string
	VatNumber        string
}
