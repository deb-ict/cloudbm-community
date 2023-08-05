package contact

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type Database interface {
	ContactRepository() ContactRepository
	ContactAddressRepository() ContactAddressRepository
	ContactEmailRepository() ContactEmailRepository
	ContactPhoneRepository() ContactPhoneRepository
	ContactUriRepository() ContactUriRepository
	CompanyRepository() CompanyRepository
	CompanyAddressRepository() CompanyAddressRepository
	CompanyEmailRepository() CompanyEmailRepository
	CompanyPhoneRepository() CompanyPhoneRepository
	CompanyUriRepository() CompanyUriRepository
	AddressTypeRepository() AddressTypeRepository
	EmailTypeRepository() EmailTypeRepository
	PhoneTypeRepository() PhoneTypeRepository
	UriTypeRepository() UriTypeRepository
	ContactTitleRepository() ContactTitleRepository
	CompanyTypeRepository() CompanyTypeRepository
	IndustryRepository() IndustryRepository
	JobTitleRepository() JobTitleRepository
}

type ContactRepository interface {
	GetContacts(ctx context.Context, offset int64, limit int64, filter *model.ContactFilter, sort *core.Sort) ([]*model.Contact, int64, error)
	GetContactById(ctx context.Context, id string) (*model.Contact, error)
	CreateContact(ctx context.Context, model *model.Contact) (string, error)
	UpdateContact(ctx context.Context, model *model.Contact) error
	DeleteContact(ctx context.Context, model *model.Contact) error
}

type ContactAddressRepository interface {
	GetContactAddresses(ctx context.Context, parent *model.Contact, offset int64, limit int64, filter *model.AddressFilter, sort *core.Sort) ([]*model.Address, int64, error)
	GetContactAddressById(ctx context.Context, parent *model.Contact, id string) (*model.Address, error)
	GetContactAddressByType(ctx context.Context, parent *model.Contact, modelType *model.AddressType) (*model.Address, error)
	GetDefaultContactAddress(ctx context.Context, parent *model.Contact) (*model.Address, error)
	CreateContactAddress(ctx context.Context, parent *model.Contact, model *model.Address) (string, error)
	UpdateContactAddress(ctx context.Context, parent *model.Contact, model *model.Address) error
	DeleteContactAddress(ctx context.Context, parent *model.Contact, model *model.Address) error
}

type ContactEmailRepository interface {
	GetContactEmails(ctx context.Context, parent *model.Contact, offset int64, limit int64, filter *model.EmailFilter, sort *core.Sort) ([]*model.Email, int64, error)
	GetContactEmailById(ctx context.Context, parent *model.Contact, id string) (*model.Email, error)
	GetContactEmailByType(ctx context.Context, parent *model.Contact, modelType *model.EmailType) (*model.Email, error)
	GetDefaultContactEmail(ctx context.Context, parent *model.Contact) (*model.Email, error)
	CreateContactEmail(ctx context.Context, parent *model.Contact, model *model.Email) (string, error)
	UpdateContactEmail(ctx context.Context, parent *model.Contact, model *model.Email) error
	DeleteContactEmail(ctx context.Context, parent *model.Contact, model *model.Email) error
}

type ContactPhoneRepository interface {
	GetContactPhones(ctx context.Context, parent *model.Contact, offset int64, limit int64, filter *model.PhoneFilter, sort *core.Sort) ([]*model.Phone, int64, error)
	GetContactPhoneById(ctx context.Context, parent *model.Contact, id string) (*model.Phone, error)
	GetContactPhoneByType(ctx context.Context, parent *model.Contact, modelType *model.PhoneType) (*model.Phone, error)
	GetDefaultContactPhone(ctx context.Context, parent *model.Contact) (*model.Phone, error)
	CreateContactPhone(ctx context.Context, parent *model.Contact, model *model.Phone) (string, error)
	UpdateContactPhone(ctx context.Context, parent *model.Contact, model *model.Phone) error
	DeleteContactPhone(ctx context.Context, parent *model.Contact, model *model.Phone) error
}

type ContactUriRepository interface {
	GetContactUris(ctx context.Context, parent *model.Contact, offset int64, limit int64, filter *model.UriFilter, sort *core.Sort) ([]*model.Uri, int64, error)
	GetContactUriById(ctx context.Context, parent *model.Contact, id string) (*model.Uri, error)
	GetContactUriByType(ctx context.Context, parent *model.Contact, modelType *model.UriType) (*model.Uri, error)
	GetDefaultContactUri(ctx context.Context, parent *model.Contact) (*model.Uri, error)
	CreateContactUri(ctx context.Context, parent *model.Contact, model *model.Uri) (string, error)
	UpdateContactUri(ctx context.Context, parent *model.Contact, model *model.Uri) error
	DeleteContactUri(ctx context.Context, parent *model.Contact, model *model.Uri) error
}

type CompanyRepository interface {
	GetCompanies(ctx context.Context, offset int64, limit int64, filter *model.CompanyFilter, sort *core.Sort) ([]*model.Company, int64, error)
	GetCompanyById(ctx context.Context, id string) (*model.Company, error)
	CreateCompany(ctx context.Context, model *model.Company) (string, error)
	UpdateCompany(ctx context.Context, model *model.Company) error
	DeleteCompany(ctx context.Context, model *model.Company) error
}

type CompanyAddressRepository interface {
	GetCompanyAddresses(ctx context.Context, parent *model.Company, offset int64, limit int64, filter *model.AddressFilter, sort *core.Sort) ([]*model.Address, int64, error)
	GetCompanyAddressById(ctx context.Context, parent *model.Company, id string) (*model.Address, error)
	GetCompanyAddressByType(ctx context.Context, parent *model.Company, modelType *model.AddressType) (*model.Address, error)
	GetDefaultCompanyAddress(ctx context.Context, parent *model.Company) (*model.Address, error)
	CreateCompanyAddress(ctx context.Context, parent *model.Company, model *model.Address) (string, error)
	UpdateCompanyAddress(ctx context.Context, parent *model.Company, model *model.Address) error
	DeleteCompanyAddress(ctx context.Context, parent *model.Company, model *model.Address) error
}

type CompanyEmailRepository interface {
	GetCompanyEmails(ctx context.Context, parent *model.Company, offset int64, limit int64, filter *model.EmailFilter, sort *core.Sort) ([]*model.Email, int64, error)
	GetCompanyEmailById(ctx context.Context, parent *model.Company, id string) (*model.Email, error)
	GetCompanyEmailByType(ctx context.Context, parent *model.Company, modelType *model.EmailType) (*model.Email, error)
	GetDefaultCompanyEmail(ctx context.Context, parent *model.Company) (*model.Email, error)
	CreateCompanyEmail(ctx context.Context, parent *model.Company, model *model.Email) (string, error)
	UpdateCompanyEmail(ctx context.Context, parent *model.Company, model *model.Email) error
	DeleteCompanyEmail(ctx context.Context, parent *model.Company, model *model.Email) error
}

type CompanyPhoneRepository interface {
	GetCompanyPhones(ctx context.Context, parent *model.Company, offset int64, limit int64, filter *model.PhoneFilter, sort *core.Sort) ([]*model.Phone, int64, error)
	GetCompanyPhoneById(ctx context.Context, parent *model.Company, id string) (*model.Phone, error)
	GetCompanyPhoneByType(ctx context.Context, parent *model.Company, modelType *model.PhoneType) (*model.Phone, error)
	GetDefaultCompanyPhone(ctx context.Context, parent *model.Company) (*model.Phone, error)
	CreateCompanyPhone(ctx context.Context, parent *model.Company, model *model.Phone) (string, error)
	UpdateCompanyPhone(ctx context.Context, parent *model.Company, model *model.Phone) error
	DeleteCompanyPhone(ctx context.Context, parent *model.Company, model *model.Phone) error
}

type CompanyUriRepository interface {
	GetCompanyUris(ctx context.Context, parent *model.Company, offset int64, limit int64, filter *model.UriFilter, sort *core.Sort) ([]*model.Uri, int64, error)
	GetCompanyUriById(ctx context.Context, parent *model.Company, id string) (*model.Uri, error)
	GetCompanyUriByType(ctx context.Context, parent *model.Company, modelType *model.UriType) (*model.Uri, error)
	GetDefaultCompanyUri(ctx context.Context, parent *model.Company) (*model.Uri, error)
	CreateCompanyUri(ctx context.Context, parent *model.Company, model *model.Uri) (string, error)
	UpdateCompanyUri(ctx context.Context, parent *model.Company, model *model.Uri) error
	DeleteCompanyUri(ctx context.Context, parent *model.Company, model *model.Uri) error
}

type AddressTypeRepository interface {
	GetAddressTypes(ctx context.Context, offset int64, limit int64, filter *model.AddressTypeFilter, sort *core.Sort) ([]*model.AddressType, int64, error)
	GetAddressTypeById(ctx context.Context, id string) (*model.AddressType, error)
	GetAddressTypeByKey(ctx context.Context, key string) (*model.AddressType, error)
	GetAddressTypeByName(ctx context.Context, language string, name string) (*model.AddressType, error)
	GetDefaultAddressType(ctx context.Context) (*model.AddressType, error)
	CreateAddressType(ctx context.Context, model *model.AddressType) (string, error)
	UpdateAddressType(ctx context.Context, model *model.AddressType) error
	DeleteAddressType(ctx context.Context, model *model.AddressType) error
}

type EmailTypeRepository interface {
	GetEmailTypes(ctx context.Context, offset int64, limit int64, filter *model.EmailTypeFilter, sort *core.Sort) ([]*model.EmailType, int64, error)
	GetEmailTypeById(ctx context.Context, id string) (*model.EmailType, error)
	GetEmailTypeByKey(ctx context.Context, key string) (*model.EmailType, error)
	GetEmailTypeByName(ctx context.Context, language string, name string) (*model.EmailType, error)
	GetDefaultEmailType(ctx context.Context) (*model.EmailType, error)
	CreateEmailType(ctx context.Context, model *model.EmailType) (string, error)
	UpdateEmailType(ctx context.Context, model *model.EmailType) error
	DeleteEmailType(ctx context.Context, model *model.EmailType) error
}

type PhoneTypeRepository interface {
	GetPhoneTypes(ctx context.Context, offset int64, limit int64, filter *model.PhoneTypeFilter, sort *core.Sort) ([]*model.PhoneType, int64, error)
	GetPhoneTypeById(ctx context.Context, id string) (*model.PhoneType, error)
	GetPhoneTypeByKey(ctx context.Context, key string) (*model.PhoneType, error)
	GetPhoneTypeByName(ctx context.Context, language string, name string) (*model.PhoneType, error)
	GetDefaultPhoneType(ctx context.Context) (*model.PhoneType, error)
	CreatePhoneType(ctx context.Context, model *model.PhoneType) (string, error)
	UpdatePhoneType(ctx context.Context, model *model.PhoneType) error
	DeletePhoneType(ctx context.Context, model *model.PhoneType) error
}

type UriTypeRepository interface {
	GetUriTypes(ctx context.Context, offset int64, limit int64, filter *model.UriTypeFilter, sort *core.Sort) ([]*model.UriType, int64, error)
	GetUriTypeById(ctx context.Context, id string) (*model.UriType, error)
	GetUriTypeByKey(ctx context.Context, key string) (*model.UriType, error)
	GetUriTypeByName(ctx context.Context, language string, name string) (*model.UriType, error)
	GetDefaultUriType(ctx context.Context) (*model.UriType, error)
	CreateUriType(ctx context.Context, model *model.UriType) (string, error)
	UpdateUriType(ctx context.Context, model *model.UriType) error
	DeleteUriType(ctx context.Context, model *model.UriType) error
}

type ContactTitleRepository interface {
	GetContactTitles(ctx context.Context, offset int64, limit int64, filter *model.ContactTitleFilter, sort *core.Sort) ([]*model.ContactTitle, int64, error)
	GetContactTitleById(ctx context.Context, id string) (*model.ContactTitle, error)
	GetContactTitleByKey(ctx context.Context, key string) (*model.ContactTitle, error)
	GetContactTitleByName(ctx context.Context, language string, name string) (*model.ContactTitle, error)
	CreateContactTitle(ctx context.Context, model *model.ContactTitle) (string, error)
	UpdateContactTitle(ctx context.Context, model *model.ContactTitle) error
	DeleteContactTitle(ctx context.Context, model *model.ContactTitle) error
}

type CompanyTypeRepository interface {
	GetCompanyTypes(ctx context.Context, offset int64, limit int64, filter *model.CompanyTypeFilter, sort *core.Sort) ([]*model.CompanyType, int64, error)
	GetCompanyTypeById(ctx context.Context, id string) (*model.CompanyType, error)
	GetCompanyTypeByKey(ctx context.Context, key string) (*model.CompanyType, error)
	GetCompanyTypeByName(ctx context.Context, language string, name string) (*model.CompanyType, error)
	CreateCompanyType(ctx context.Context, model *model.CompanyType) (string, error)
	UpdateCompanyType(ctx context.Context, model *model.CompanyType) error
	DeleteCompanyType(ctx context.Context, model *model.CompanyType) error
}

type IndustryRepository interface {
	GetIndustries(ctx context.Context, offset int64, limit int64, filter *model.IndustryFilter, sort *core.Sort) ([]*model.Industry, int64, error)
	GetIndustryById(ctx context.Context, id string) (*model.Industry, error)
	GetIndustryByKey(ctx context.Context, key string) (*model.Industry, error)
	GetIndustryByName(ctx context.Context, language string, name string) (*model.Industry, error)
	CreateIndustry(ctx context.Context, model *model.Industry) (string, error)
	UpdateIndustry(ctx context.Context, model *model.Industry) error
	DeleteIndustry(ctx context.Context, model *model.Industry) error
}

type JobTitleRepository interface {
	GetJobTitles(ctx context.Context, offset int64, limit int64, filter *model.JobTitleFilter, sort *core.Sort) ([]*model.JobTitle, int64, error)
	GetJobTitleById(ctx context.Context, id string) (*model.JobTitle, error)
	GetJobTitleByKey(ctx context.Context, key string) (*model.JobTitle, error)
	GetJobTitleByName(ctx context.Context, language string, name string) (*model.JobTitle, error)
	CreateJobTitle(ctx context.Context, model *model.JobTitle) (string, error)
	UpdateJobTitle(ctx context.Context, model *model.JobTitle) error
	DeleteJobTitle(ctx context.Context, model *model.JobTitle) error
}
