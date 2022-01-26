package customer

type PhoneType struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Primary     bool   `json:"primary"`
}

type Phone struct {
	Id          string `json:"id"`
	TypeId      string `json:"type_id"`
	Value       string `json:"value"`
	Description string `json:"description"`
	Primary     bool   `json:"primary"`
}

type EmailType struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Primary     bool   `json:"primary"`
}

type Email struct {
	Id          string `json:"id"`
	TypeId      string `json:"type_id"`
	Value       string `json:"value"`
	Description string `json:"description"`
	Primary     bool   `json:"primary"`
}

type AddressType struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Primary     bool   `json:"primary"`
}

type Address struct {
	Id         string `json:"id"`
	TypeId     string `json:"type_id"`
	Address    string `json:"address"`
	PostalCode string `json:"postal"`
	City       string `json:"city"`
	Country    string `json:"country"`
}

type Company struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Emails    []Email   `json:"emails,omitempty"`
	Phones    []Phone   `json:"phones,omitempty"`
	Addresses []Address `json:"addresses,omitempty"`
}

type Contact struct {
	Id        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Emails    []Email   `json:"emails,omitempty"`
	Phones    []Phone   `json:"phones,omitempty"`
	Addresses []Address `json:"addresses,omitempty"`
}
