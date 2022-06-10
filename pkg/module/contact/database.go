package contact

import (
	"context"
)

type Database interface {
	GetAddressTypeStore() AddressTypeStore
	GetPhoneTypeStore() PhoneTypeStore
	GetEmailTypeStore() EmailTypeStore
	GetUrlTypeStore() UrlTypeStore
	GetContactStore() ContactStore
	GetContactAddressStore() ContactAddressStore
	GetContactPhoneStore() ContactPhoneStore
	GetContactEmailStore() ContactEmailStore
	GetCompanyStore() CompanyStore
	GetCompanyAddressStore() CompanyAddressStore
	GetCompanyPhoneStore() CompanyPhoneStore
	GetCompanyEmailStore() CompanyEmailStore
	GetCompanyUrlStore() CompanyUrlStore
}

type AddressTypeStore interface {
	GetAddressTypes(ctx context.Context, pageIndex int, pageSize int) (*AddressTypeList, error)
	GetAddressTypeById(ctx context.Context, id string) (*AddressType, error)
	GetAddressTypeByName(ctx context.Context, name string) (*AddressType, error)
	GetDefaultAddressType(ctx context.Context) (*AddressType, error)
	CreateAddressType(ctx context.Context, addressType AddressType) (string, error)
	UpdateAddressType(ctx context.Context, id string, addressType AddressType) error
	DeleteAddressType(ctx context.Context, id string) error
}

type PhoneTypeStore interface {
	GetPhoneTypes(ctx context.Context, pageIndex int, pageSize int) (*PhoneTypeList, error)
	GetPhoneTypeById(ctx context.Context, id string) (*PhoneType, error)
	GetPhoneTypeByName(ctx context.Context, name string) (*PhoneType, error)
	GetDefaultPhoneType(ctx context.Context) (*PhoneType, error)
	CreatePhoneType(ctx context.Context, phoneType PhoneType) (string, error)
	UpdatePhoneType(ctx context.Context, id string, phoneType PhoneType) error
	DeletePhoneType(ctx context.Context, id string) error
}

type EmailTypeStore interface {
	GetEmailTypes(ctx context.Context, pageIndex int, pageSize int) (*EmailTypeList, error)
	GetEmailTypeById(ctx context.Context, id string) (*EmailType, error)
	GetEmailTypeByName(ctx context.Context, name string) (*EmailType, error)
	GetDefaultEmailType(ctx context.Context) (*EmailType, error)
	CreateEmailType(ctx context.Context, emailType EmailType) (string, error)
	UpdateEmailType(ctx context.Context, id string, emailType EmailType) error
	DeleteEmailType(ctx context.Context, id string) error
}

type UrlTypeStore interface {
	GetUrlTypes(ctx context.Context, pageIndex int, pageSize int) (*UrlTypeList, error)
	GetUrlTypeById(ctx context.Context, id string) (*UrlType, error)
	GetUrlTypeByName(ctx context.Context, name string) (*UrlType, error)
	GetDefaultUrlType(ctx context.Context) (*UrlType, error)
	CreateUrlType(ctx context.Context, urlType UrlType) (string, error)
	UpdateUrlType(ctx context.Context, id string, urlType UrlType) error
	DeleteUrlType(ctx context.Context, id string) error
}

type ContactStore interface {
	GetContacts(ctx context.Context, pageIndex int, pageSize int) (*ContactList, error)
	GetContactById(ctx context.Context, id string) (*Contact, error)
	GetContactByName(ctx context.Context, familyName string, giveName string) (*Contact, error)
	CreateContact(ctx context.Context, contact Contact) (string, error)
	UpdateContact(ctx context.Context, id string, contact Contact) error
	DeleteContact(ctx context.Context, id string) error
}

type ContactAddressStore interface {
	GetContactAddresses(ctx context.Context, contactId string, pageIndex int, pageSize int) (*AddressList, error)
	GetContactAddressById(ctx context.Context, contactId string, id string) (*Address, error)
	GetContactAddressByType(ctx context.Context, contactId string, typeId string) (*Address, error)
	GetContactPrimaryAddress(ctx context.Context, contactId string) (*Address, error)
	CreateContactAddress(ctx context.Context, contactId string, address Address) (string, error)
	UpdateContactAddress(ctx context.Context, contactId string, id string, address Address) error
	DeleteContactAddress(ctx context.Context, contactId string, id string) error
}

type ContactPhoneStore interface {
	GetContactPhones(ctx context.Context, contactId string, pageIndex int, pageSize int) (*PhoneList, error)
	GetContactPhoneById(ctx context.Context, contactId string, id string) (*Phone, error)
	GetContactPhoneByType(ctx context.Context, contactId string, typeId string) (*Phone, error)
	GetContactPrimaryPhone(ctx context.Context, contactId string) (*Phone, error)
	CreateContactPhone(ctx context.Context, contactId string, phone Phone) (string, error)
	UpdateContactPhone(ctx context.Context, contactId string, id string, phone Phone) error
	DeleteContactPhone(ctx context.Context, contactId string, id string) error
}

type ContactEmailStore interface {
	GetContactEmails(ctx context.Context, contactId string, pageIndex int, pageSize int) (*EmailList, error)
	GetContactEmailById(ctx context.Context, contactId string, id string) (*Email, error)
	GetContactEmailByType(ctx context.Context, contactId string, typeId string) (*Email, error)
	GetContactPrimaryEmail(ctx context.Context, contactId string) (*Email, error)
	CreateContactEmail(ctx context.Context, contactId string, email Email) (string, error)
	UpdateContactEmail(ctx context.Context, contactId string, id string, email Email) error
	DeleteContactEmail(ctx context.Context, contactId string, id string) error
}

type CompanyStore interface {
	GetCompanies(ctx context.Context, pageIndex int, pageSize int) (*CompanyList, error)
	GetCompanyById(ctx context.Context, id string) (*Company, error)
	GetCompanyByName(ctx context.Context, name string) (*Company, error)
	CreateCompany(ctx context.Context, company Company) (string, error)
	UpdateCompany(ctx context.Context, id string, company Company) error
	DeleteCompany(ctx context.Context, id string) error
}

type CompanyAddressStore interface {
	GetCompanyAddresses(ctx context.Context, companyId string, pageIndex int, pageSize int) (*AddressList, error)
	GetCompanyAddressById(ctx context.Context, companyId string, id string) (*Address, error)
	GetCompanyAddressByType(ctx context.Context, companyId string, typeId string) (*Address, error)
	GetCompanyPrimaryAddress(ctx context.Context, companyId string) (*Address, error)
	CreateCompanyAddress(ctx context.Context, companyId string, address Address) (string, error)
	UpdateCompanyAddress(ctx context.Context, companyId string, id string, address Address) error
	DeleteCompanyAddress(ctx context.Context, companyId string, id string) error
}

type CompanyPhoneStore interface {
	GetCompanyPhones(ctx context.Context, companyId string, pageIndex int, pageSize int) (*PhoneList, error)
	GetCompanyPhoneById(ctx context.Context, companyId string, id string) (*Phone, error)
	GetCompanyPhoneByType(ctx context.Context, companyId string, typeId string) (*Phone, error)
	GetCompanyPrimaryPhone(ctx context.Context, companyId string) (*Phone, error)
	CreateCompanyPhone(ctx context.Context, companyId string, phone Phone) (string, error)
	UpdateCompanyPhone(ctx context.Context, companyId string, id string, phone Phone) error
	DeleteCompanyPhone(ctx context.Context, companyId string, id string) error
}

type CompanyEmailStore interface {
	GetCompanyEmails(ctx context.Context, companyId string, pageIndex int, pageSize int) (*EmailList, error)
	GetCompanyEmailById(ctx context.Context, companyId string, id string) (*Email, error)
	GetCompanyEmailByType(ctx context.Context, companyId string, typeId string) (*Email, error)
	GetCompanyPrimaryEmail(ctx context.Context, companyId string) (*Email, error)
	CreateCompanyEmail(ctx context.Context, companyId string, email Email) (string, error)
	UpdateCompanyEmail(ctx context.Context, companyId string, id string, email Email) error
	DeleteCompanyEmail(ctx context.Context, companyId string, id string) error
}

type CompanyUrlStore interface {
	GetCompanyUrls(ctx context.Context, companyId string, pageIndex int, pageSize int) (*UrlList, error)
	GetCompanyUrlById(ctx context.Context, companyId string, id string) (*Url, error)
	GetCompanyUrlByType(ctx context.Context, companyId string, typeId string) (*Url, error)
	GetCompanyPrimaryUrl(ctx context.Context, companyId string) (*Url, error)
	CreateCompanyUrl(ctx context.Context, companyId string, url Url) (string, error)
	UpdateCompanyUrl(ctx context.Context, companyId string, id string, url Url) error
	DeleteCompanyUrl(ctx context.Context, companyId string, id string) error
}
