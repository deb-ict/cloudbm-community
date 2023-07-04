package contact

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type Service interface {
	GetContacts(ctx context.Context, offset int64, limit int64, filter *model.ContactFilter, sort *core.Sort) ([]*model.Contact, int64, error)
	GetContactById(ctx context.Context, id string) (*model.Contact, error)
	CreateContact(ctx context.Context, contact *model.Contact) (*model.Contact, error)
	UpdateContact(ctx context.Context, id string, contact *model.Contact) (*model.Contact, error)
	DeleteContact(ctx context.Context, id string) (*model.Contact, error)

	GetCompanies(ctx context.Context, offset int64, limit int64, filter *model.CompanyFilter, sort *core.Sort) ([]*model.Company, int64, error)
	GetCompanyById(ctx context.Context, id string) (*model.Company, error)
	CreateCompany(ctx context.Context, company *model.Company) (*model.Company, error)
	UpdateCompany(ctx context.Context, id string, company *model.Company) (*model.Company, error)
	DeleteCompany(ctx context.Context, id string) (*model.Company, error)

	GetAddressTypes(ctx context.Context, offset int64, limit int64, filter *model.AddressTypeFilter, sort *core.Sort) ([]*model.AddressType, int64, error)
	GetAddressTypeById(ctx context.Context, id string) (*model.AddressType, error)
	CreateAddressType(ctx context.Context, addressType *model.AddressType) (*model.AddressType, error)
	UpdateAddressType(ctx context.Context, id string, addressType *model.AddressType) (*model.AddressType, error)
	DeleteAddressType(ctx context.Context, id string) (*model.AddressType, error)

	GetEmailTypes(ctx context.Context, offset int64, limit int64, filter *model.EmailTypeFilter, sort *core.Sort) ([]*model.EmailType, int64, error)
	GetEmailTypeById(ctx context.Context, id string) (*model.EmailType, error)
	CreateEmailType(ctx context.Context, emailType *model.EmailType) (*model.EmailType, error)
	UpdateEmailType(ctx context.Context, id string, emailType *model.EmailType) (*model.EmailType, error)
	DeleteEmailType(ctx context.Context, id string) (*model.EmailType, error)

	GetPhoneTypes(ctx context.Context, offset int64, limit int64, filter *model.PhoneTypeFilter, sort *core.Sort) ([]*model.PhoneType, int64, error)
	GetPhoneTypeById(ctx context.Context, id string) (*model.PhoneType, error)
	CreatePhoneType(ctx context.Context, phoneType *model.PhoneType) (*model.PhoneType, error)
	UpdatePhoneType(ctx context.Context, id string, phoneType *model.PhoneType) (*model.PhoneType, error)
	DeletePhoneType(ctx context.Context, id string) (*model.PhoneType, error)

	GetUriTypes(ctx context.Context, offset int64, limit int64, filter *model.UriTypeFilter, sort *core.Sort) ([]*model.UriType, int64, error)
	GetUriTypeById(ctx context.Context, id string) (*model.UriType, error)
	CreateUriType(ctx context.Context, uriType *model.UriType) (*model.UriType, error)
	UpdateUriType(ctx context.Context, id string, uriType *model.UriType) (*model.UriType, error)
	DeleteUriType(ctx context.Context, id string) (*model.UriType, error)

	GetContactTitles(ctx context.Context, offset int64, limit int64, filter *model.ContactTitleFilter, sort *core.Sort) ([]*model.ContactTitle, int64, error)
	GetContactTitleById(ctx context.Context, id string) (*model.ContactTitle, error)
	CreateContactTitle(ctx context.Context, contactTitle *model.ContactTitle) (*model.ContactTitle, error)
	UpdateContactTitle(ctx context.Context, id string, contactTitle *model.ContactTitle) (*model.ContactTitle, error)
	DeleteContactTitle(ctx context.Context, id string) (*model.ContactTitle, error)

	GetCompanyTypes(ctx context.Context, offset int64, limit int64, filter *model.CompanyTypeFilter, sort *core.Sort) ([]*model.CompanyType, int64, error)
	GetCompanyTypeById(ctx context.Context, id string) (*model.CompanyType, error)
	CreateCompanyType(ctx context.Context, companyType *model.CompanyType) (*model.CompanyType, error)
	UpdateCompanyType(ctx context.Context, id string, companyType *model.CompanyType) (*model.CompanyType, error)
	DeleteCompanyType(ctx context.Context, id string) (*model.CompanyType, error)

	GetIndustryTypes(ctx context.Context, offset int64, limit int64, filter *model.IndustryFilter, sort *core.Sort) ([]*model.Industry, int64, error)
	GetIndustryById(ctx context.Context, id string) (*model.Industry, error)
	CreateIndustry(ctx context.Context, industry *model.Industry) (*model.Industry, error)
	UpdateIndustry(ctx context.Context, id string, industry *model.Industry) (*model.Industry, error)
	DeleteIndustry(ctx context.Context, id string) (*model.Industry, error)

	GetJobTitles(ctx context.Context, offset int64, limit int64, filter *model.JobTitleFilter, sort *core.Sort) ([]*model.JobTitle, int64, error)
	GetJobTitleById(ctx context.Context, id string) (*model.JobTitle, error)
	CreateJobTitle(ctx context.Context, jobTitle *model.JobTitle) (*model.JobTitle, error)
	UpdateJobTitle(ctx context.Context, id string, jobTitle *model.JobTitle) (*model.JobTitle, error)
	DeleteJobTitle(ctx context.Context, id string) (*model.JobTitle, error)
}
