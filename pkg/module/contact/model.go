package contact

import (
	"time"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
)

type AddressType struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"is_default"`
	IsSystem  bool   `json:"is_system"`
}

type AddressTypeList struct {
	rest.PaginatedList
	Items []AddressType `json:"items"`
}

type PhoneType struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"is_default"`
	IsSystem  bool   `json:"is_system"`
}

type PhoneTypeList struct {
	rest.PaginatedList
	Items []PhoneType `json:"items"`
}

type EmailType struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"is_default"`
	IsSystem  bool   `json:"is_system"`
}

type EmailTypeList struct {
	rest.PaginatedList
	Items []EmailType `json:"items"`
}

type UrlType struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"is_default"`
	IsSystem  bool   `json:"is_system"`
}

type UrlTypeList struct {
	rest.PaginatedList
	Items []UrlType `json:"items"`
}

type Address struct {
	Id         string `json:"id"`
	TypeId     string `json:"type_id"`
	Street     string `json:"street"`
	PostalCode string `json:"postal_code"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	IsPrimary  bool   `json:"is_primary"`
}

type AddressList struct {
	rest.PaginatedList
	Items []Address `json:"items"`
}

type Phone struct {
	Id        string `json:"id"`
	TypeId    string `json:"type_id"`
	Value     string `json:"value"`
	IsPrimary bool   `json:"is_primary"`
}

type PhoneList struct {
	rest.PaginatedList
	Items []Phone `json:"items"`
}

type Email struct {
	Id        string `json:"id"`
	TypeId    string `json:"type_id"`
	Value     string `json:"value"`
	IsPrimary bool   `json:"is_primary"`
}

type EmailList struct {
	rest.PaginatedList
	Items []Email `json:"items"`
}

type Url struct {
	Id        string `json:"id"`
	TypeId    string `json:"type_id"`
	Value     string `json:"value"`
	IsPrimary bool   `json:"is_primary"`
}

type UrlList struct {
	rest.PaginatedList
	Items []Url `json:"items"`
}

type Contact struct {
	Id         string    `json:"id"`
	ParentId   string    `json:"parent_id"`
	Title      string    `json:"title"`
	FamilyName string    `json:"family_name"`
	GivenName  string    `json:"given_name"`
	Birthdate  time.Time `json:"birthdate"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ContactList struct {
	rest.PaginatedList
	Items []Contact `json:"items"`
}

type Company struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CompanyList struct {
	rest.PaginatedList
	Items []Company `json:"items"`
}
