package contact

import (
	"context"
	"errors"
)

var (
	ErrAddressTypeNotFound     = errors.New("address type not found")
	ErrAddressTypeReadOnly     = errors.New("address type is read only")
	ErrAddressTypeDuplicate    = errors.New("address type already exists")
	ErrPhoneTypeNotFound       = errors.New("phone type not found")
	ErrPhoneTypeReadOnly       = errors.New("phone type is read only")
	ErrPhoneTypeDuplicate      = errors.New("phone type already exists")
	ErrEmailTypeNotFound       = errors.New("email type not found")
	ErrEmailTypeReadOnly       = errors.New("email type is read only")
	ErrEmailTypeDuplicate      = errors.New("email type already exists")
	ErrUrlTypeNotFound         = errors.New("url type not found")
	ErrUrlTypeReadOnly         = errors.New("url type is read only")
	ErrUrlTypeDuplicate        = errors.New("url type already exists")
	ErrContactNotFound         = errors.New("contact not found")
	ErrContactDuplicateName    = errors.New("contact with name already exists")
	ErrContactAddressNotFound  = errors.New("contact address not found")
	ErrContactAddressDuplicate = errors.New("contact address already exists")
	ErrContactPhoneNotFound    = errors.New("contact phone not found")
	ErrContactPhoneDuplicate   = errors.New("contact phone alredy exists")
	ErrContactEmailNotFound    = errors.New("contact email not found")
	ErrContactEmailDuplicate   = errors.New("contact email already exists")
	ErrCompanyNotFound         = errors.New("company not found")
	ErrCompanyDuplicateName    = errors.New("company with name already exists")
	ErrCompanyAddressNotFound  = errors.New("company address not found")
	ErrCompanyAddressDuplicate = errors.New("company address already exists")
	ErrCompanyPhoneNotFound    = errors.New("company phone not found")
	ErrCompanyPhoneDuplicate   = errors.New("company phone already exists")
	ErrCompanyEmailNotFound    = errors.New("company email not found")
	ErrCompanyEmailDuplicate   = errors.New("company email already exists")
	ErrCompanyUrlNotFound      = errors.New("company url not found")
	ErrCompanyeUrlDuplicate    = errors.New("company url type already exists")
)

type Service interface {
	GetDatabase() Database
	SetDatabase(database Database)

	GetAddressTypes(ctx context.Context, pageIndex int, pageSize int) (*AddressTypeList, error)
	GetAddressTypeById(ctx context.Context, id string) (*AddressType, error)
	CreateAddressType(ctx context.Context, addressType AddressType) (*AddressType, error)
	UpdateAddressType(ctx context.Context, id string, addressType AddressType) (*AddressType, error)
	DeleteAddressType(ctx context.Context, id string) error
	ResetDefaultAddressType(ctx context.Context) error

	GetPhoneTypes(ctx context.Context, pageIndex int, pageSize int) (*PhoneTypeList, error)
	GetPhoneTypeById(ctx context.Context, id string) (*PhoneType, error)
	CreatePhoneType(ctx context.Context, phoneType PhoneType) (*PhoneType, error)
	UpdatePhoneType(ctx context.Context, id string, phoneType PhoneType) (*PhoneType, error)
	DeletePhoneType(ctx context.Context, id string) error
	ResetDefaultPhoneType(ctx context.Context) error

	GetEmailTypes(ctx context.Context, pageIndex int, pageSize int) (*EmailTypeList, error)
	GetEmailTypeById(ctx context.Context, id string) (*EmailType, error)
	CreateEmailType(ctx context.Context, emailType EmailType) (*EmailType, error)
	UpdateEmailType(ctx context.Context, id string, emailType EmailType) (*EmailType, error)
	DeleteEmailType(ctx context.Context, id string) error
	ResetDefaultEmailType(ctx context.Context) error

	GetUrlTypes(ctx context.Context, pageIndex int, pageSize int) (*UrlTypeList, error)
	GetUrlTypeById(ctx context.Context, id string) (*UrlType, error)
	CreateUrlType(ctx context.Context, urlType UrlType) (*UrlType, error)
	UpdateUrlType(ctx context.Context, id string, urlType UrlType) (*UrlType, error)
	DeleteUrlType(ctx context.Context, id string) error
	ResetDefaultUrlType(ctx context.Context) error

	GetContacts(ctx context.Context, pageIndex int, pageSize int) (*ContactList, error)
	GetContactById(ctx context.Context, id string) (*Contact, error)
	CreateContact(ctx context.Context, contact Contact) (*Contact, error)
	UpdateContact(ctx context.Context, id string, contact Contact) (*Contact, error)
	DeleteContact(ctx context.Context, id string) error

	GetContactAddresses(ctx context.Context, contactId string, pageIndex int, pageSize int) (*AddressList, error)
	GetContactAddressById(ctx context.Context, contactId string, id string) (*Address, error)
	CreateContactAddress(ctx context.Context, contactId string, address Address) (*Address, error)
	UpdateContactAddress(ctx context.Context, contactId string, id string, address Address) (*Address, error)
	DeleteContactAddress(ctx context.Context, contactId string, id string) error
	ResetPrimaryContactAddress(ctx context.Context, contactId string) error

	GetContactPhones(ctx context.Context, contactId string, pageIndex int, pageSize int) (*PhoneList, error)
	GetContactPhoneById(ctx context.Context, contactId string, id string) (*Phone, error)
	CreateContactPhone(ctx context.Context, contactId string, phone Phone) (*Phone, error)
	UpdateContactPhone(ctx context.Context, contactId string, id string, phone Phone) (*Phone, error)
	DeleteContactPhone(ctx context.Context, contactId string, id string) error
	ResetPrimaryContactPhone(ctx context.Context, contactId string) error

	GetContactEmails(ctx context.Context, contactId string, pageIndex int, pageSize int) (*EmailList, error)
	GetContactEmailById(ctx context.Context, contactId string, id string) (*Email, error)
	CreateContactEmail(ctx context.Context, contactId string, email Email) (*Email, error)
	UpdateContactEmail(ctx context.Context, contactId string, id string, email Email) (*Email, error)
	DeleteContactEmail(ctx context.Context, contactId string, id string) error
	ResetPrimaryContactEmail(ctx context.Context, contactId string) error

	GetCompanies(ctx context.Context, pageIndex int, pageSize int) (*CompanyList, error)
	GetCompanyById(ctx context.Context, id string) (*Company, error)
	CreateCompany(ctx context.Context, company Company) (*Company, error)
	UpdateCompany(ctx context.Context, id string, company Company) (*Company, error)
	DeleteCompany(ctx context.Context, id string) error

	GetCompanyAddresses(ctx context.Context, companyId string, pageIndex int, pageSize int) (*AddressList, error)
	GetCompanyAddressById(ctx context.Context, companyId string, id string) (*Address, error)
	CreateCompanyAddress(ctx context.Context, companyId string, address Address) (*Address, error)
	UpdateCompanyAddress(ctx context.Context, companyId string, id string, address Address) (*Address, error)
	DeleteCompanyAddress(ctx context.Context, companyId string, id string) error
	ResetPrimaryCompanyAddress(ctx context.Context, companyId string) error

	GetCompanyPhones(ctx context.Context, companyId string, pageIndex int, pageSize int) (*PhoneList, error)
	GetCompanyPhoneById(ctx context.Context, companyId string, id string) (*Phone, error)
	CreateCompanyPhone(ctx context.Context, companyId string, phone Phone) (*Phone, error)
	UpdateCompanyPhone(ctx context.Context, companyId string, id string, phone Phone) (*Phone, error)
	DeleteCompanyPhone(ctx context.Context, companyId string, id string) error
	ResetPrimaryCompanyPhone(ctx context.Context, companyId string) error

	GetCompanyEmails(ctx context.Context, companyId string, pageIndex int, pageSize int) (*EmailList, error)
	GetCompanyEmailById(ctx context.Context, companyId string, id string) (*Email, error)
	CreateCompanyEmail(ctx context.Context, companyId string, email Email) (*Email, error)
	UpdateCompanyEmail(ctx context.Context, companyId string, id string, email Email) (*Email, error)
	DeleteCompanyEmail(ctx context.Context, companyId string, id string) error
	ResetPrimaryCompanyEmail(ctx context.Context, companyId string) error

	GetCompanyUrls(ctx context.Context, companyId string, pageIndex int, pageSize int) (*UrlList, error)
	GetCompanyUrlById(ctx context.Context, companyId string, id string) (*Url, error)
	CreateCompanyUrl(ctx context.Context, companyId string, url Url) (*Url, error)
	UpdateCompanyUrl(ctx context.Context, companyId string, id string, url Url) (*Url, error)
	DeleteCompanyUrl(ctx context.Context, companyId string, id string) error
	ResetPrimaryCompanyUrl(ctx context.Context, companyId string) error
}

func NewService() Service {
	return &service{}
}

type service struct {
	database Database
}

func (svc *service) GetDatabase() Database {
	return svc.database
}

func (svc *service) SetDatabase(database Database) {
	svc.database = database
}
