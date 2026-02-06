package v1

import (
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/gorilla/mux"
)

const (
	RouteGetContactsV1           = "contact_api:GetContacts:v1"
	RouteGetContactByIdV1        = "contact_api:GetContactById:v1"
	RouteCreateContactV1         = "contact_api:CreateContact:v1"
	RouteUpdateContactV1         = "contact_api:UpdateContact:v1"
	RouteDeleteContactV1         = "contact_api:DeleteContact:v1"
	RouteGetContactAddressesV1   = "contact_api:GetContactAddresses:v1"
	RouteGetContactAddressByIdV1 = "contact_api:GetContactAddressById:v1"
	RouteCreateContactAddressV1  = "contact_api:CreateContactAddress:v1"
	RouteUpdateContactAddressV1  = "contact_api:UpdateContactAddress:v1"
	RouteDeleteContactAddressV1  = "contact_api:DeleteContactAddress:v1"
	RouteGetContactEmailsV1      = "contact_api:GetContactEmails:v1"
	RouteGetContactEmailByIdV1   = "contact_api:GetContactEmailById:v1"
	RouteCreateContactEmailV1    = "contact_api:CreateContactEmail:v1"
	RouteUpdateContactEmailV1    = "contact_api:UpdateContactEmail:v1"
	RouteDeleteContactEmailV1    = "contact_api:DeleteContactEmail:v1"
	RouteGetContactPhonesV1      = "contact_api:GetContactPhones:v1"
	RouteGetContactPhoneByIdV1   = "contact_api:GetContactPhoneById:v1"
	RouteCreateContactPhoneV1    = "contact_api:CreateContactPhone:v1"
	RouteUpdateContactPhoneV1    = "contact_api:UpdateContactPhone:v1"
	RouteDeleteContactPhoneV1    = "contact_api:DeleteContactPhone:v1"
	RouteGetContactUrisV1        = "contact_api:GetContactUris:v1"
	RouteGetContactUriByIdV1     = "contact_api:GetContactUriById:v1"
	RouteCreateContactUriV1      = "contact_api:CreateContactUri:v1"
	RouteUpdateContactUriV1      = "contact_api:UpdateContactUri:v1"
	RouteDeleteContactUriV1      = "contact_api:DeleteContactUri:v1"
	RouteGetCompaniesV1          = "contact_api:GetCompanies:v1"
	RouteGetCompanyByIdV1        = "contact_api:GetCompanyById:v1"
	RouteCreateCompanyV1         = "contact_api:CreateCompany:v1"
	RouteUpdateCompanyV1         = "contact_api:UpdateCompany:v1"
	RouteDeleteCompanyV1         = "contact_api:DeleteCompany:v1"
	RouteGetCompanyAddressesV1   = "contact_api:GetCompanyAddresses:v1"
	RouteGetCompanyAddressByIdV1 = "contact_api:GetCompanyAddressById:v1"
	RouteCreateCompanyAddressV1  = "contact_api:CreateCompanyAddress:v1"
	RouteUpdateCompanyAddressV1  = "contact_api:UpdateCompanyAddress:v1"
	RouteDeleteCompanyAddressV1  = "contact_api:DeleteCompanyAddress:v1"
	RouteGetCompanyEmailsV1      = "contact_api:GetCompanyEmails:v1"
	RouteGetCompanyEmailByIdV1   = "contact_api:GetCompanyEmailById:v1"
	RouteCreateCompanyEmailV1    = "contact_api:CreateCompanyEmail:v1"
	RouteUpdateCompanyEmailV1    = "contact_api:UpdateCompanyEmail:v1"
	RouteDeleteCompanyEmailV1    = "contact_api:DeleteCompanyEmail:v1"
	RouteGetCompanyPhonesV1      = "contact_api:GetCompanyPhones:v1"
	RouteGetCompanyPhoneByIdV1   = "contact_api:GetCompanyPhoneById:v1"
	RouteCreateCompanyPhoneV1    = "contact_api:CreateCompanyPhone:v1"
	RouteUpdateCompanyPhoneV1    = "contact_api:UpdateCompanyPhone:v1"
	RouteDeleteCompanyPhoneV1    = "contact_api:DeleteCompanyPhone:v1"
	RouteGetCompanyUrisV1        = "contact_api:GetCompanyUris:v1"
	RouteGetCompanyUriByIdV1     = "contact_api:GetCompanyUriById:v1"
	RouteCreateCompanyUriV1      = "contact_api:CreateCompanyUri:v1"
	RouteUpdateCompanyUriV1      = "contact_api:UpdateCompanyUri:v1"
	RouteDeleteCompanyUriV1      = "contact_api:DeleteCompanyUri:v1"
	RouteGetAddressTypesV1       = "contact_api:GetAddressTypes:v1"
	RouteGetAddressTypeByIdV1    = "contact_api:GetAddressTypeById:v1"
	RouteCreateAddressTypeV1     = "contact_api:CreateAddressType:v1"
	RouteUpdateAddressTypeV1     = "contact_api:UpdateAddressType:v1"
	RouteDeleteAddressTypeV1     = "contact_api:DeleteAddressType:v1"
	RouteGetEmailTypesV1         = "contact_api:GetEmailTypes:v1"
	RouteGetEmailTypeByIdV1      = "contact_api:GetEmailTypeById:v1"
	RouteCreateEmailTypeV1       = "contact_api:CreateEmailType:v1"
	RouteUpdateEmailTypeV1       = "contact_api:UpdateEmailType:v1"
	RouteDeleteEmailTypeV1       = "contact_api:DeleteEmailType:v1"
	RouteGetPhoneTypesV1         = "contact_api:GetPhoneTypes:v1"
	RouteGetPhoneTypeByIdV1      = "contact_api:GetPhoneTypeById:v1"
	RouteCreatePhoneTypeV1       = "contact_api:CreatePhoneType:v1"
	RouteUpdatePhoneTypeV1       = "contact_api:UpdatePhoneType:v1"
	RouteDeletePhoneTypeV1       = "contact_api:DeletePhoneType:v1"
	RouteGetUriTypesV1           = "contact_api:GetUriTypes:v1"
	RouteGetUriTypeByIdV1        = "contact_api:GetUriTypeById:v1"
	RouteCreateUriTypeV1         = "contact_api:CreateUriType:v1"
	RouteUpdateUriTypeV1         = "contact_api:UpdateUriType:v1"
	RouteDeleteUriTypeV1         = "contact_api:DeleteUriType:v1"
	RouteGetContactTitlesV1      = "contact_api:GetContactTitles:v1"
	RouteGetContactTitleByIdV1   = "contact_api:GetContactTitleById:v1"
	RouteCreateContactTitleV1    = "contact_api:CreateContactTitle:v1"
	RouteUpdateContactTitleV1    = "contact_api:UpdateContactTitle:v1"
	RouteDeleteContactTitleV1    = "contact_api:DeleteContactTitle:v1"
	RouteGetCompanyTypesV1       = "contact_api:GetCompanyTypes:v1"
	RouteGetCompanyTypeByIdV1    = "contact_api:GetCompanyTypeById:v1"
	RouteCreateCompanyTypeV1     = "contact_api:CreateCompanyType:v1"
	RouteUpdateCompanyTypeV1     = "contact_api:UpdateCompanyType:v1"
	RouteDeleteCompanyTypeV1     = "contact_api:DeleteCompanyType:v1"
	RouteGetIndustriesV1         = "contact_api:GetIndustries:v1"
	RouteGetIndustryByIdV1       = "contact_api:GetIndustryById:v1"
	RouteCreateIndustryV1        = "contact_api:CreateIndustry:v1"
	RouteUpdateIndustryV1        = "contact_api:UpdateIndustry:v1"
	RouteDeleteIndustryV1        = "contact_api:DeleteIndustry:v1"
	RouteGetJobTitlesV1          = "contact_api:GetJobTitles:v1"
	RouteGetJobTitleByIdV1       = "contact_api:GetJobTitleById:v1"
	RouteCreateJobTitleV1        = "contact_api:CreateJobTitle:v1"
	RouteUpdateJobTitleV1        = "contact_api:UpdateJobTitle:v1"
	RouteDeleteJobTitleV1        = "contact_api:DeleteJobTitle:v1"
)

type ApiV1 interface {
	RegisterRoutes(r *mux.Router)
}

type apiV1 struct {
	service contact.Service
}

func NewApiV1(service contact.Service) ApiV1 {
	return &apiV1{
		service: service,
	}
}

func (api *apiV1) RegisterRoutes(r *mux.Router) {
	// Contacts
	r.HandleFunc("/v1/contact", api.GetContactsHandlerV1).Methods(http.MethodGet).Name(RouteGetContactsV1)
	r.HandleFunc("/v1/contact/{id}", api.GetContactByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetContactByIdV1)
	r.HandleFunc("/v1/contact", api.CreateContactHandlerV1).Methods(http.MethodPost).Name(RouteCreateContactV1)
	r.HandleFunc("/v1/contact/{id}", api.UpdateContactHandlerV1).Methods(http.MethodPut).Name(RouteUpdateContactV1)
	r.HandleFunc("/v1/contact/{id}", api.DeleteContactHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteContactV1)
	// Contact addresses
	r.HandleFunc("/v1/contact/{contactId}/address", api.GetContactAddressesHandlerV1).Methods(http.MethodGet).Name(RouteGetContactAddressesV1)
	r.HandleFunc("/v1/contact/{contactId}/address/{id}", api.GetContactAddressByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetContactAddressByIdV1)
	r.HandleFunc("/v1/contact/{contactId}/address", api.CreateContactAddressHandlerV1).Methods(http.MethodPost).Name(RouteCreateContactAddressV1)
	r.HandleFunc("/v1/contact/{contactId}/address/{id}", api.UpdateContactAddressHandlerV1).Methods(http.MethodPut).Name(RouteUpdateContactAddressV1)
	r.HandleFunc("/v1/contact/{contactId}/address/{id}", api.DeleteContactAddressHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteContactAddressV1)

	// Contact emails
	r.HandleFunc("/v1/contact/{contactId}/email", api.GetContactEmailsHandlerV1).Methods(http.MethodGet).Name(RouteGetContactEmailsV1)
	r.HandleFunc("/v1/contact/{contactId}/email/{id}", api.GetContactEmailByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetContactEmailByIdV1)
	r.HandleFunc("/v1/contact/{contactId}/email", api.CreateContactEmailHandlerV1).Methods(http.MethodPost).Name(RouteCreateContactEmailV1)
	r.HandleFunc("/v1/contact/{contactId}/email/{id}", api.UpdateContactEmailHandlerV1).Methods(http.MethodPut).Name(RouteUpdateContactEmailV1)
	r.HandleFunc("/v1/contact/{contactId}/email/{id}", api.DeleteContactEmailHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteContactEmailV1)

	// Contact phones
	r.HandleFunc("/v1/contact/{contactId}/phone", api.GetContactPhonesHandlerV1).Methods(http.MethodGet).Name(RouteGetContactPhonesV1)
	r.HandleFunc("/v1/contact/{contactId}/phone/{id}", api.GetContactPhoneByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetContactPhoneByIdV1)
	r.HandleFunc("/v1/contact/{contactId}/phone", api.CreateContactPhoneHandlerV1).Methods(http.MethodPost).Name(RouteCreateContactPhoneV1)
	r.HandleFunc("/v1/contact/{contactId}/phone/{id}", api.UpdateContactPhoneHandlerV1).Methods(http.MethodPut).Name(RouteUpdateContactPhoneV1)
	r.HandleFunc("/v1/contact/{contactId}/phone/{id}", api.DeleteContactPhoneHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteContactPhoneV1)

	// Contact uris
	r.HandleFunc("/v1/contact/{contactId}/uri", api.GetContactUrisHandlerV1).Methods(http.MethodGet).Name(RouteGetContactUrisV1)
	r.HandleFunc("/v1/contact/{contactId}/uri/{id}", api.GetContactUriByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetContactUriByIdV1)
	r.HandleFunc("/v1/contact/{contactId}/uri", api.CreateContactUriHandlerV1).Methods(http.MethodPost).Name(RouteCreateContactUriV1)
	r.HandleFunc("/v1/contact/{contactId}/uri/{id}", api.UpdateContactUriHandlerV1).Methods(http.MethodPut).Name(RouteUpdateContactUriV1)
	r.HandleFunc("/v1/contact/{contactId}/uri/{id}", api.DeleteContactUriHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteContactUriV1)

	// Companies
	r.HandleFunc("/v1/company", api.GetCompaniesHandlerV1).Methods(http.MethodGet).Name(RouteGetCompaniesV1)
	r.HandleFunc("/v1/company/{id}", api.GetCompanyByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetCompanyByIdV1)
	r.HandleFunc("/v1/company", api.CreateCompanyHandlerV1).Methods(http.MethodPost).Name(RouteCreateCompanyV1)
	r.HandleFunc("/v1/company/{id}", api.UpdateCompanyHandlerV1).Methods(http.MethodPut).Name(RouteUpdateCompanyV1)
	r.HandleFunc("/v1/company/{id}", api.DeleteCompanyHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteCompanyV1)

	// Company addresses
	r.HandleFunc("/v1/company/{companyId}/address", api.GetCompanyAddressesHandlerV1).Methods(http.MethodGet).Name(RouteGetCompanyAddressesV1)
	r.HandleFunc("/v1/company/{companyId}/address/{id}", api.GetCompanyAddressByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetCompanyAddressByIdV1)
	r.HandleFunc("/v1/company/{companyId}/address", api.CreateCompanyAddressHandlerV1).Methods(http.MethodPost).Name(RouteCreateCompanyAddressV1)
	r.HandleFunc("/v1/company/{companyId}/address/{id}", api.UpdateCompanyAddressHandlerV1).Methods(http.MethodPut).Name(RouteUpdateCompanyAddressV1)
	r.HandleFunc("/v1/company/{companyId}/address/{id}", api.DeleteCompanyAddressHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteCompanyAddressV1)

	// Company emails
	r.HandleFunc("/v1/company/{companyId}/email", api.GetCompanyEmailsHandlerV1).Methods(http.MethodGet).Name(RouteGetCompanyEmailsV1)
	r.HandleFunc("/v1/company/{companyId}/email/{id}", api.GetCompanyEmailByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetCompanyEmailByIdV1)
	r.HandleFunc("/v1/company/{companyId}/email", api.CreateCompanyEmailHandlerV1).Methods(http.MethodPost).Name(RouteCreateCompanyEmailV1)
	r.HandleFunc("/v1/company/{companyId}/email/{id}", api.UpdateCompanyEmailHandlerV1).Methods(http.MethodPut).Name(RouteUpdateCompanyEmailV1)
	r.HandleFunc("/v1/company/{companyId}/email/{id}", api.DeleteCompanyEmailHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteCompanyEmailV1)

	// Company phones
	r.HandleFunc("/v1/company/{companyId}/phone", api.GetCompanyPhonesHandlerV1).Methods(http.MethodGet).Name(RouteGetCompanyPhonesV1)
	r.HandleFunc("/v1/company/{companyId}/phone/{id}", api.GetCompanyPhoneByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetCompanyPhoneByIdV1)
	r.HandleFunc("/v1/company/{companyId}/phone", api.CreateCompanyPhoneHandlerV1).Methods(http.MethodPost).Name(RouteCreateCompanyPhoneV1)
	r.HandleFunc("/v1/company/{companyId}/phone/{id}", api.UpdateCompanyPhoneHandlerV1).Methods(http.MethodPut).Name(RouteUpdateCompanyPhoneV1)
	r.HandleFunc("/v1/company/{companyId}/phone/{id}", api.DeleteCompanyPhoneHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteCompanyPhoneV1)

	// Company uris
	r.HandleFunc("/v1/company/{companyId}/uri", api.GetCompanyUrisHandlerV1).Methods(http.MethodGet).Name(RouteGetCompanyUrisV1)
	r.HandleFunc("/v1/company/{companyId}/uri/{id}", api.GetCompanyUriByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetCompanyUriByIdV1)
	r.HandleFunc("/v1/company/{companyId}/uri", api.CreateCompanyUriHandlerV1).Methods(http.MethodPost).Name(RouteCreateCompanyUriV1)
	r.HandleFunc("/v1/company/{companyId}/uri/{id}", api.UpdateCompanyUriHandlerV1).Methods(http.MethodPut).Name(RouteUpdateCompanyUriV1)
	r.HandleFunc("/v1/company/{companyId}/uri/{id}", api.DeleteCompanyUriHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteCompanyUriV1)

	// Address types
	r.HandleFunc("/v1/addressType", api.GetAddressTypesHandlerV1).Methods(http.MethodGet).Name(RouteGetAddressTypesV1)
	r.HandleFunc("/v1/addressType/{id}", api.GetAddressTypeByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetAddressTypeByIdV1)
	r.HandleFunc("/v1/addressType", api.CreateAddressTypeHandlerV1).Methods(http.MethodPost).Name(RouteCreateAddressTypeV1)
	r.HandleFunc("/v1/addressType/{id}", api.UpdateAddressTypeHandlerV1).Methods(http.MethodPut).Name(RouteUpdateAddressTypeV1)
	r.HandleFunc("/v1/addressType/{id}", api.DeleteAddressTypeHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteAddressTypeV1)
	// Email types
	r.HandleFunc("/v1/emailType", api.GetEmailTypesHandlerV1).Methods(http.MethodGet).Name(RouteGetEmailTypesV1)
	r.HandleFunc("/v1/emailType/{id}", api.GetEmailTypeByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetEmailTypeByIdV1)
	r.HandleFunc("/v1/emailType", api.CreateEmailTypeHandlerV1).Methods(http.MethodPost).Name(RouteCreateEmailTypeV1)
	r.HandleFunc("/v1/emailType/{id}", api.UpdateEmailTypeHandlerV1).Methods(http.MethodPut).Name(RouteUpdateEmailTypeV1)
	r.HandleFunc("/v1/emailType/{id}", api.DeleteEmailTypeHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteEmailTypeV1)
	// Phone types
	r.HandleFunc("/v1/phoneType", api.GetPhoneTypesHandlerV1).Methods(http.MethodGet).Name(RouteGetPhoneTypesV1)
	r.HandleFunc("/v1/phoneType/{id}", api.GetPhoneTypeByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetPhoneTypeByIdV1)
	r.HandleFunc("/v1/phoneType", api.CreatePhoneTypeHandlerV1).Methods(http.MethodPost).Name(RouteCreatePhoneTypeV1)
	r.HandleFunc("/v1/phoneType/{id}", api.UpdatePhoneTypeHandlerV1).Methods(http.MethodPut).Name(RouteUpdatePhoneTypeV1)
	r.HandleFunc("/v1/phoneType/{id}", api.DeletePhoneTypeHandlerV1).Methods(http.MethodDelete).Name(RouteDeletePhoneTypeV1)

	// Uri types
	r.HandleFunc("/v1/uriType", api.GetUriTypesHandlerV1).Methods(http.MethodGet).Name(RouteGetUriTypesV1)
	r.HandleFunc("/v1/uriType/{id}", api.GetUriTypeByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetUriTypeByIdV1)
	r.HandleFunc("/v1/uriType", api.CreateUriTypeHandlerV1).Methods(http.MethodPost).Name(RouteCreateUriTypeV1)
	r.HandleFunc("/v1/uriType/{id}", api.UpdateUriTypeHandlerV1).Methods(http.MethodPut).Name(RouteUpdateUriTypeV1)
	r.HandleFunc("/v1/uriType/{id}", api.DeleteUriTypeHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteUriTypeV1)

	// Contact titles
	r.HandleFunc("/v1/contactTitle", api.GetContactTitlesHandlerV1).Methods(http.MethodGet).Name(RouteGetContactTitlesV1)
	r.HandleFunc("/v1/contactTitle/{id}", api.GetContactTitleByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetContactTitleByIdV1)
	r.HandleFunc("/v1/contactTitle", api.CreateContactTitleHandlerV1).Methods(http.MethodPost).Name(RouteCreateContactTitleV1)
	r.HandleFunc("/v1/contactTitle/{id}", api.UpdateContactTitleHandlerV1).Methods(http.MethodPut).Name(RouteUpdateContactTitleV1)
	r.HandleFunc("/v1/contactTitle/{id}", api.DeleteContactTitleHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteContactTitleV1)

	// Company types
	r.HandleFunc("/v1/companyType", api.GetCompanyTypesHandlerV1).Methods(http.MethodGet).Name(RouteGetCompanyTypesV1)
	r.HandleFunc("/v1/companyType/{id}", api.GetCompanyTypeByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetCompanyTypeByIdV1)
	r.HandleFunc("/v1/companyType", api.CreateCompanyTypeHandlerV1).Methods(http.MethodPost).Name(RouteCreateCompanyTypeV1)
	r.HandleFunc("/v1/companyType/{id}", api.UpdateCompanyTypeHandlerV1).Methods(http.MethodPut).Name(RouteUpdateCompanyTypeV1)
	r.HandleFunc("/v1/companyType/{id}", api.DeleteCompanyTypeHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteCompanyTypeV1)

	// Industries
	r.HandleFunc("/v1/industry", api.GetIndustriesHandlerV1).Methods(http.MethodGet).Name(RouteGetIndustriesV1)
	r.HandleFunc("/v1/industry/{id}", api.GetIndustryByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetIndustryByIdV1)
	r.HandleFunc("/v1/industry", api.CreateIndustryHandlerV1).Methods(http.MethodPost).Name(RouteCreateIndustryV1)
	r.HandleFunc("/v1/industry/{id}", api.UpdateIndustryHandlerV1).Methods(http.MethodPut).Name(RouteUpdateIndustryV1)
	r.HandleFunc("/v1/industry/{id}", api.DeleteIndustryHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteIndustryV1)

	// Job titles
	r.HandleFunc("/v1/jobTitle", api.GetJobTitlesHandlerV1).Methods(http.MethodGet).Name(RouteGetJobTitlesV1)
	r.HandleFunc("/v1/jobTitle/{id}", api.GetJobTitleByIdHandlerV1).Methods(http.MethodGet).Name(RouteGetJobTitleByIdV1)
	r.HandleFunc("/v1/jobTitle", api.CreateJobTitleHandlerV1).Methods(http.MethodPost).Name(RouteCreateJobTitleV1)
	r.HandleFunc("/v1/jobTitle/{id}", api.UpdateJobTitleHandlerV1).Methods(http.MethodPut).Name(RouteUpdateJobTitleV1)
	r.HandleFunc("/v1/jobTitle/{id}", api.DeleteJobTitleHandlerV1).Methods(http.MethodDelete).Name(RouteDeleteJobTitleV1)
}

func (api *apiV1) handleError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	switch err {
	case contact.ErrContactNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case contact.ErrContactReadOnly:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrContactAddressNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case contact.ErrContactAddressIsDefault:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrContactEmailNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case contact.ErrContactEmailIsDefault:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrContactPhoneNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case contact.ErrContactPhoneIsDefault:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrContactUriNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case contact.ErrContactUriIsDefault:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrCompanyNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case contact.ErrCompanyReadOnly:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrCompanyAddressNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case contact.ErrCompanyAddressIsDefault:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrCompanyEmailNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case contact.ErrCompanyEmailIsDefault:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrCompanyPhoneNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case contact.ErrCompanyPhoneIsDefault:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrCompanyUriNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case contact.ErrCompanyUriIsDefault:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrAddressTypeNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case contact.ErrAddressTypeDuplicateKey:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrAddressTypeDuplicateName:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrAddressTypeReadOnly:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrAddressTypeIsDefault:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrEmailTypeNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case contact.ErrEmailTypeDuplicateKey:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrEmailTypeDuplicateName:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrEmailTypeReadOnly:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrEmailTypeIsDefault:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrPhoneTypeNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case contact.ErrPhoneTypeDuplicateKey:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrPhoneTypeDuplicateName:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrPhoneTypeReadOnly:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrPhoneTypeIsDefault:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrUriTypeNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case contact.ErrUriTypeDuplicateKey:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrUriTypeDuplicateName:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrUriTypeReadOnly:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrUriTypeIsDefault:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrContactTitleNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case contact.ErrContactTitleDuplicateKey:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrContactTitleDuplicateName:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrContactTitleReadOnly:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrJobTitleNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case contact.ErrJobTitleDuplicateKey:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrJobTitleDuplicateName:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrJobTitleReadOnly:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrCompanyTypeNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case contact.ErrCompanyTypeDuplicateKey:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrCompanyTypeDuplicateName:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrCompanyTypeReadOnly:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrIndustryNotFound:
		rest.WriteError(w, http.StatusNotFound, err.Error())
	case contact.ErrIndustryDuplicateKey:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrIndustryDuplicateName:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	case contact.ErrIndustryReadOnly:
		rest.WriteError(w, http.StatusBadRequest, err.Error())
	default:
		rest.WriteError(w, http.StatusInternalServerError, err.Error())
	}
	return true
}
