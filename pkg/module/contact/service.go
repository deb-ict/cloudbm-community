package contact

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type Service interface {
	StringNormalizer() core.StringNormalizer
	FeatureProvider() core.FeatureProvider
	LanguageProvider() localization.LanguageProvider

	GetContacts(ctx context.Context, offset int64, limit int64, filter *model.ContactFilter, sort *core.Sort) ([]*model.Contact, int64, error)
	GetContactById(ctx context.Context, id string) (*model.Contact, error)
	CreateContact(ctx context.Context, model *model.Contact) (*model.Contact, error)
	UpdateContact(ctx context.Context, id string, model *model.Contact) (*model.Contact, error)
	DeleteContact(ctx context.Context, id string) error

	GetContactAddresses(ctx context.Context, contactId string, offset int64, limit int64, filter *model.AddressFilter, sort *core.Sort) ([]*model.Address, int64, error)
	GetContactAddressById(ctx context.Context, contactId string, id string) (*model.Address, error)
	CreateContactAddress(ctx context.Context, contactId string, model *model.Address) (*model.Address, error)
	UpdateContactAddress(ctx context.Context, contactId string, id string, model *model.Address) (*model.Address, error)
	DeleteContactAddress(ctx context.Context, contactId string, id string) error

	GetContactEmails(ctx context.Context, contactId string, offset int64, limit int64, filter *model.EmailFilter, sort *core.Sort) ([]*model.Email, int64, error)
	GetContactEmailById(ctx context.Context, contactId string, id string) (*model.Email, error)
	CreateContactEmail(ctx context.Context, contactId string, model *model.Email) (*model.Email, error)
	UpdateContactEmail(ctx context.Context, contactId string, id string, model *model.Email) (*model.Email, error)
	DeleteContactEmail(ctx context.Context, contactId string, id string) error

	GetContactPhones(ctx context.Context, contactId string, offset int64, limit int64, filter *model.PhoneFilter, sort *core.Sort) ([]*model.Phone, int64, error)
	GetContactPhoneById(ctx context.Context, contactId string, id string) (*model.Phone, error)
	CreateContactPhone(ctx context.Context, contactId string, model *model.Phone) (*model.Phone, error)
	UpdateContactPhone(ctx context.Context, contactId string, id string, model *model.Phone) (*model.Phone, error)
	DeleteContactPhone(ctx context.Context, contactId string, id string) error

	GetContactUris(ctx context.Context, contactId string, offset int64, limit int64, filter *model.UriFilter, sort *core.Sort) ([]*model.Uri, int64, error)
	GetContactUriById(ctx context.Context, contactId string, id string) (*model.Uri, error)
	CreateContactUri(ctx context.Context, contactId string, model *model.Uri) (*model.Uri, error)
	UpdateContactUri(ctx context.Context, contactId string, id string, model *model.Uri) (*model.Uri, error)
	DeleteContactUri(ctx context.Context, contactId string, id string) error

	GetCompanies(ctx context.Context, offset int64, limit int64, filter *model.CompanyFilter, sort *core.Sort) ([]*model.Company, int64, error)
	GetCompanyById(ctx context.Context, id string) (*model.Company, error)
	CreateCompany(ctx context.Context, model *model.Company) (*model.Company, error)
	UpdateCompany(ctx context.Context, id string, model *model.Company) (*model.Company, error)
	DeleteCompany(ctx context.Context, id string) error

	GetCompanyAddresses(ctx context.Context, companyId string, offset int64, limit int64, filter *model.AddressFilter, sort *core.Sort) ([]*model.Address, int64, error)
	GetCompanyAddressById(ctx context.Context, companyId string, id string) (*model.Address, error)
	CreateCompanyAddress(ctx context.Context, companyId string, model *model.Address) (*model.Address, error)
	UpdateCompanyAddress(ctx context.Context, companyId string, id string, model *model.Address) (*model.Address, error)
	DeleteCompanyAddress(ctx context.Context, companyId string, id string) error

	GetCompanyEmails(ctx context.Context, companyId string, offset int64, limit int64, filter *model.EmailFilter, sort *core.Sort) ([]*model.Email, int64, error)
	GetCompanyEmailById(ctx context.Context, companyId string, id string) (*model.Email, error)
	CreateCompanyEmail(ctx context.Context, companyId string, model *model.Email) (*model.Email, error)
	UpdateCompanyEmail(ctx context.Context, companyId string, id string, model *model.Email) (*model.Email, error)
	DeleteCompanyEmail(ctx context.Context, companyId string, id string) error

	GetCompanyPhones(ctx context.Context, companyId string, offset int64, limit int64, filter *model.PhoneFilter, sort *core.Sort) ([]*model.Phone, int64, error)
	GetCompanyPhoneById(ctx context.Context, companyId string, id string) (*model.Phone, error)
	CreateCompanyPhone(ctx context.Context, companyId string, model *model.Phone) (*model.Phone, error)
	UpdateCompanyPhone(ctx context.Context, companyId string, id string, model *model.Phone) (*model.Phone, error)
	DeleteCompanyPhone(ctx context.Context, companyId string, id string) error

	GetCompanyUris(ctx context.Context, companyId string, offset int64, limit int64, filter *model.UriFilter, sort *core.Sort) ([]*model.Uri, int64, error)
	GetCompanyUriById(ctx context.Context, companyId string, id string) (*model.Uri, error)
	CreateCompanyUri(ctx context.Context, companyId string, model *model.Uri) (*model.Uri, error)
	UpdateCompanyUri(ctx context.Context, companyId string, id string, model *model.Uri) (*model.Uri, error)
	DeleteCompanyUri(ctx context.Context, companyId string, id string) error

	GetAddressTypes(ctx context.Context, offset int64, limit int64, filter *model.AddressTypeFilter, sort *core.Sort) ([]*model.AddressType, int64, error)
	GetAddressTypeById(ctx context.Context, id string) (*model.AddressType, error)
	CreateAddressType(ctx context.Context, model *model.AddressType) (*model.AddressType, error)
	UpdateAddressType(ctx context.Context, id string, model *model.AddressType) (*model.AddressType, error)
	DeleteAddressType(ctx context.Context, id string) error

	GetEmailTypes(ctx context.Context, offset int64, limit int64, filter *model.EmailTypeFilter, sort *core.Sort) ([]*model.EmailType, int64, error)
	GetEmailTypeById(ctx context.Context, id string) (*model.EmailType, error)
	CreateEmailType(ctx context.Context, model *model.EmailType) (*model.EmailType, error)
	UpdateEmailType(ctx context.Context, id string, model *model.EmailType) (*model.EmailType, error)
	DeleteEmailType(ctx context.Context, id string) error

	GetPhoneTypes(ctx context.Context, offset int64, limit int64, filter *model.PhoneTypeFilter, sort *core.Sort) ([]*model.PhoneType, int64, error)
	GetPhoneTypeById(ctx context.Context, id string) (*model.PhoneType, error)
	CreatePhoneType(ctx context.Context, model *model.PhoneType) (*model.PhoneType, error)
	UpdatePhoneType(ctx context.Context, id string, model *model.PhoneType) (*model.PhoneType, error)
	DeletePhoneType(ctx context.Context, id string) error

	GetUriTypes(ctx context.Context, offset int64, limit int64, filter *model.UriTypeFilter, sort *core.Sort) ([]*model.UriType, int64, error)
	GetUriTypeById(ctx context.Context, id string) (*model.UriType, error)
	CreateUriType(ctx context.Context, model *model.UriType) (*model.UriType, error)
	UpdateUriType(ctx context.Context, id string, model *model.UriType) (*model.UriType, error)
	DeleteUriType(ctx context.Context, id string) error

	GetContactTitles(ctx context.Context, offset int64, limit int64, filter *model.ContactTitleFilter, sort *core.Sort) ([]*model.ContactTitle, int64, error)
	GetContactTitleById(ctx context.Context, id string) (*model.ContactTitle, error)
	CreateContactTitle(ctx context.Context, model *model.ContactTitle) (*model.ContactTitle, error)
	UpdateContactTitle(ctx context.Context, id string, model *model.ContactTitle) (*model.ContactTitle, error)
	DeleteContactTitle(ctx context.Context, id string) error

	GetCompanyTypes(ctx context.Context, offset int64, limit int64, filter *model.CompanyTypeFilter, sort *core.Sort) ([]*model.CompanyType, int64, error)
	GetCompanyTypeById(ctx context.Context, id string) (*model.CompanyType, error)
	CreateCompanyType(ctx context.Context, model *model.CompanyType) (*model.CompanyType, error)
	UpdateCompanyType(ctx context.Context, id string, model *model.CompanyType) (*model.CompanyType, error)
	DeleteCompanyType(ctx context.Context, id string) error

	GetIndustries(ctx context.Context, offset int64, limit int64, filter *model.IndustryFilter, sort *core.Sort) ([]*model.Industry, int64, error)
	GetIndustryById(ctx context.Context, id string) (*model.Industry, error)
	CreateIndustry(ctx context.Context, model *model.Industry) (*model.Industry, error)
	UpdateIndustry(ctx context.Context, id string, model *model.Industry) (*model.Industry, error)
	DeleteIndustry(ctx context.Context, id string) error

	GetJobTitles(ctx context.Context, offset int64, limit int64, filter *model.JobTitleFilter, sort *core.Sort) ([]*model.JobTitle, int64, error)
	GetJobTitleById(ctx context.Context, id string) (*model.JobTitle, error)
	CreateJobTitle(ctx context.Context, model *model.JobTitle) (*model.JobTitle, error)
	UpdateJobTitle(ctx context.Context, id string, model *model.JobTitle) (*model.JobTitle, error)
	DeleteJobTitle(ctx context.Context, id string) error
}
