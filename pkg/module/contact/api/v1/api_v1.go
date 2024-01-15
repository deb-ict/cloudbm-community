package v1

import (
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/gorilla/mux"
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
	r.HandleFunc("/v1/contact", api.GetContactsHandlerV1).Methods(http.MethodGet).Name("contact_api:GetContactsHandlerV1")
	r.HandleFunc("/v1/contact/{id}", api.GetContactByIdHandlerV1).Methods(http.MethodGet).Name("contact_api:GetContactByIdHandlerV1")
	r.HandleFunc("/v1/contact", api.CreateContactHandlerV1).Methods(http.MethodPost).Name("contact_api:CreateContactHandlerV1")
	r.HandleFunc("/v1/contact/{id}", api.UpdateContactHandlerV1).Methods(http.MethodPut).Name("contact_api:UpdateContactHandlerV1")
	r.HandleFunc("/v1/contact/{id}", api.DeleteContactHandlerV1).Methods(http.MethodDelete).Name("contact_api:DeleteContactHandlerV1")

	// Contact addresses
	r.HandleFunc("/v1/contact/{contactId}/address", api.GetContactAddressesHandlerV1).Methods(http.MethodGet).Name("contact_api:GetContactAddressesHandlerV1")
	r.HandleFunc("/v1/contact/{contactId}/address/{id}", api.GetContactAddressByIdHandlerV1).Methods(http.MethodGet).Name("contact_api:GetContactAddressByIdHandlerV1")
	r.HandleFunc("/v1/contact/{contactId}/address", api.CreateContactAddressHandlerV1).Methods(http.MethodPost).Name("contact_api:CreateContactAddressHandlerV1")
	r.HandleFunc("/v1/contact/{contactId}/address/{id}", api.UpdateContactAddressHandlerV1).Methods(http.MethodPut).Name("contact_api:UpdateContactAddressHandlerV1")
	r.HandleFunc("/v1/contact/{contactId}/address/{id}", api.DeleteContactAddressHandlerV1).Methods(http.MethodDelete).Name("contact_api:DeleteContactAddressHandlerV1")

	// Contact emails
	r.HandleFunc("/v1/contact/{contactId}/email", api.GetContactEmailsHandlerV1).Methods(http.MethodGet).Name("contact_api:GetContactEmailsHandlerV1")
	r.HandleFunc("/v1/contact/{contactId}/email/{id}", api.GetContactEmailByIdHandlerV1).Methods(http.MethodGet).Name("contact_api:GetContactEmailByIdHandlerV1")
	r.HandleFunc("/v1/contact/{contactId}/email", api.CreateContactEmailHandlerV1).Methods(http.MethodPost).Name("contact_api:CreateContactEmailHandlerV1")
	r.HandleFunc("/v1/contact/{contactId}/email/{id}", api.UpdateContactEmailHandlerV1).Methods(http.MethodPut).Name("contact_api:UpdateContactEmailHandlerV1")
	r.HandleFunc("/v1/contact/{contactId}/email/{id}", api.DeleteContactEmailHandlerV1).Methods(http.MethodDelete).Name("contact_api:DeleteContactEmailHandlerV1")

	// Contact phones
	r.HandleFunc("/v1/contact/{contactId}/phone", api.GetContactPhonesHandlerV1).Methods(http.MethodGet).Name("contact_api:GetContactPhonesHandlerV1")
	r.HandleFunc("/v1/contact/{contactId}/phone/{id}", api.GetContactPhoneByIdHandlerV1).Methods(http.MethodGet).Name("contact_api:GetContactPhoneByIdHandlerV1")
	r.HandleFunc("/v1/contact/{contactId}/phone", api.CreateContactPhoneHandlerV1).Methods(http.MethodPost).Name("contact_api:CreateContactPhoneHandlerV1")
	r.HandleFunc("/v1/contact/{contactId}/phone/{id}", api.UpdateContactPhoneHandlerV1).Methods(http.MethodPut).Name("contact_api:UpdateContactPhoneHandlerV1")
	r.HandleFunc("/v1/contact/{contactId}/phone/{id}", api.DeleteContactPhoneHandlerV1).Methods(http.MethodDelete).Name("contact_api:DeleteContactPhoneHandlerV1")

	// Contact uris
	r.HandleFunc("/v1/contact/{contactId}/uri", api.GetContactUrisHandlerV1).Methods(http.MethodGet).Name("contact_api:GetContactUrisHandlerV1")
	r.HandleFunc("/v1/contact/{contactId}/uri/{id}", api.GetContactUriByIdHandlerV1).Methods(http.MethodGet).Name("contact_api:GetContactUriByIdHandlerV1")
	r.HandleFunc("/v1/contact/{contactId}/uri", api.CreateContactUriHandlerV1).Methods(http.MethodPost).Name("contact_api:CreateContactUriHandlerV1")
	r.HandleFunc("/v1/contact/{contactId}/uri/{id}", api.UpdateContactUriHandlerV1).Methods(http.MethodPut).Name("contact_api:UpdateContactUriHandlerV1")
	r.HandleFunc("/v1/contact/{contactId}/uri/{id}", api.DeleteContactUriHandlerV1).Methods(http.MethodDelete).Name("contact_api:DeleteContactUriHandlerV1")

	// Companies
	r.HandleFunc("/v1/company", api.GetCompaniesHandlerV1).Methods(http.MethodGet).Name("contact_api:GetCompaniesHandlerV1")
	r.HandleFunc("/v1/company/{id}", api.GetCompanyByIdHandlerV1).Methods(http.MethodGet).Name("contact_api:GetCompanyByIdHandlerV1")
	r.HandleFunc("/v1/company", api.CreateCompanyHandlerV1).Methods(http.MethodPost).Name("contact_api:CreateCompanyHandlerV1")
	r.HandleFunc("/v1/company/{id}", api.UpdateCompanyHandlerV1).Methods(http.MethodPut).Name("contact_api:UpdateCompanyHandlerV1")
	r.HandleFunc("/v1/company/{id}", api.DeleteCompanyHandlerV1).Methods(http.MethodDelete).Name("contact_api:DeleteCompanyHandlerV1")

	// Company addresses
	r.HandleFunc("/v1/company/{companyId}/address", api.GetCompanyAddressesHandlerV1).Methods(http.MethodGet).Name("contact_api:GetCompanyAddressesHandlerV1")
	r.HandleFunc("/v1/company/{companyId}/address/{id}", api.GetCompanyAddressByIdHandlerV1).Methods(http.MethodGet).Name("contact_api:GetCompanyAddressByIdHandlerV1")
	r.HandleFunc("/v1/company/{companyId}/address", api.CreateCompanyAddressHandlerV1).Methods(http.MethodPost).Name("contact_api:CreateCompanyAddressHandlerV1")
	r.HandleFunc("/v1/company/{companyId}/address/{id}", api.UpdateCompanyAddressHandlerV1).Methods(http.MethodPut).Name("contact_api:UpdateCompanyAddressHandlerV1")
	r.HandleFunc("/v1/company/{companyId}/address/{id}", api.DeleteCompanyAddressHandlerV1).Methods(http.MethodDelete).Name("contact_api:DeleteCompanyAddressHandlerV1")

	// Company emails
	r.HandleFunc("/v1/company/{companyId}/email", api.GetCompanyEmailsHandlerV1).Methods(http.MethodGet).Name("contact_api:GetCompanyEmailsHandlerV1")
	r.HandleFunc("/v1/company/{companyId}/email/{id}", api.GetCompanyEmailByIdHandlerV1).Methods(http.MethodGet).Name("contact_api:GetCompanyEmailByIdHandlerV1")
	r.HandleFunc("/v1/company/{companyId}/email", api.CreateCompanyEmailHandlerV1).Methods(http.MethodPost).Name("contact_api:CreateCompanyEmailHandlerV1")
	r.HandleFunc("/v1/company/{companyId}/email/{id}", api.UpdateCompanyEmailHandlerV1).Methods(http.MethodPut).Name("contact_api:UpdateCompanyEmailHandlerV1")
	r.HandleFunc("/v1/company/{companyId}/email/{id}", api.DeleteCompanyEmailHandlerV1).Methods(http.MethodDelete).Name("contact_api:DeleteCompanyEmailHandlerV1")

	// Company phones
	r.HandleFunc("/v1/company/{companyId}/phone", api.GetCompanyPhonesHandlerV1).Methods(http.MethodGet).Name("contact_api:GetCompanyPhonesHandlerV1")
	r.HandleFunc("/v1/company/{companyId}/phone/{id}", api.GetCompanyByIdHandlerV1).Methods(http.MethodGet).Name("contact_api:GetCompanyByIdHandlerV1")
	r.HandleFunc("/v1/company/{companyId}/phone", api.CreateCompanyPhoneHandlerV1).Methods(http.MethodPost).Name("contact_api:CreateCompanyPhoneHandlerV1")
	r.HandleFunc("/v1/company/{companyId}/phone/{id}", api.UpdateCompanyPhoneHandlerV1).Methods(http.MethodPut).Name("contact_api:UpdateCompanyPhoneHandlerV1")
	r.HandleFunc("/v1/company/{companyId}/phone/{id}", api.DeleteCompanyPhoneHandlerV1).Methods(http.MethodDelete).Name("contact_api:DeleteCompanyPhoneHandlerV1")

	// Company uris
	r.HandleFunc("/v1/company/{companyId}/uri", api.GetCompanyUrisHandlerV1).Methods(http.MethodGet).Name("contact_api:GetCompanyUrisHandlerV1")
	r.HandleFunc("/v1/company/{companyId}/uri/{id}", api.GetCompanyUriByIdHandlerV1).Methods(http.MethodGet).Name("contact_api:GetCompanyUriByIdHandlerV1")
	r.HandleFunc("/v1/company/{companyId}/uri", api.CreateCompanyUriHandlerV1).Methods(http.MethodPost).Name("contact_api:CreateCompanyUriHandlerV1")
	r.HandleFunc("/v1/company/{companyId}/uri/{id}", api.UpdateCompanyUriHandlerV1).Methods(http.MethodPut).Name("contact_api:UpdateCompanyUriHandlerV1")
	r.HandleFunc("/v1/company/{companyId}/uri/{id}", api.DeleteCompanyUriHandlerV1).Methods(http.MethodDelete).Name("contact_api:DeleteCompanyUriHandlerV1")

	// Address types
	r.HandleFunc("/v1/addressType", api.GetAddressTypesHandlerV1).Methods(http.MethodGet).Name("contact_api:GetAddressTypesHandlerV1")
	r.HandleFunc("/v1/addressType/{id}", api.GetAddressTypeByIdHandlerV1).Methods(http.MethodGet).Name("contact_api:GetAddressTypeByIdHandlerV1")
	r.HandleFunc("/v1/addressType", api.CreateAddressTypeHandlerV1).Methods(http.MethodPost).Name("contact_api:CreateAddressTypeHandlerV1")
	r.HandleFunc("/v1/addressType/{id}", api.UpdateAddressTypeHandlerV1).Methods(http.MethodPut).Name("contact_api:UpdateAddressTypeHandlerV1")
	r.HandleFunc("/v1/addressType/{id}", api.DeleteAddressTypeHandlerV1).Methods(http.MethodDelete).Name("contact_api:DeleteAddressTypeHandlerV1")

	// Email types
	r.HandleFunc("/v1/emailType", api.GetEmailTypesHandlerV1).Methods(http.MethodGet).Name("contact_api:GetEmailTypesHandlerV1")
	r.HandleFunc("/v1/emailType/{id}", api.GetEmailTypeByIdHandlerV1).Methods(http.MethodGet).Name("contact_api:GetEmailTypeByIdHandlerV1")
	r.HandleFunc("/v1/emailType", api.CreateEmailTypeHandlerV1).Methods(http.MethodPost).Name("contact_api:CreateEmailTypeHandlerV1")
	r.HandleFunc("/v1/emailType/{id}", api.UpdateEmailTypeHandlerV1).Methods(http.MethodPut).Name("contact_api:UpdateEmailTypeHandlerV1")
	r.HandleFunc("/v1/emailType/{id}", api.DeleteEmailTypeHandlerV1).Methods(http.MethodDelete).Name("contact_api:DeleteEmailTypeHandlerV1")

	// Phone types
	r.HandleFunc("/v1/phoneType", api.GetPhoneTypesHandlerV1).Methods(http.MethodGet).Name("contact_api:GetPhoneTypesHandlerV1")
	r.HandleFunc("/v1/phoneType/{id}", api.GetPhoneTypeByIdHandlerV1).Methods(http.MethodGet).Name("contact_api:GetPhoneTypeByIdHandlerV1")
	r.HandleFunc("/v1/phoneType", api.CreatePhoneTypeHandlerV1).Methods(http.MethodPost).Name("contact_api:CreatePhoneTypeHandlerV1")
	r.HandleFunc("/v1/phoneType/{id}", api.UpdatePhoneTypeHandlerV1).Methods(http.MethodPut).Name("contact_api:UpdatePhoneTypeHandlerV1")
	r.HandleFunc("/v1/phoneType/{id}", api.DeletePhoneTypeHandlerV1).Methods(http.MethodDelete).Name("contact_api:DeletePhoneTypeHandlerV1")

	// Uri types
	r.HandleFunc("/v1/uriType", api.GetUriTypesHandlerV1).Methods(http.MethodGet).Name("contact_api:GetUriTypesHandlerV1")
	r.HandleFunc("/v1/uriType/{id}", api.GetUriTypeByIdHandlerV1).Methods(http.MethodGet).Name("contact_api:GetUriTypeByIdHandlerV1")
	r.HandleFunc("/v1/uriType", api.CreateUriTypeHandlerV1).Methods(http.MethodPost).Name("contact_api:CreateUriTypeHandlerV1")
	r.HandleFunc("/v1/uriType/{id}", api.UpdateUriTypeHandlerV1).Methods(http.MethodPut).Name("contact_api:UpdateUriTypeHandlerV1")
	r.HandleFunc("/v1/uriType/{id}", api.DeleteUriTypeHandlerV1).Methods(http.MethodDelete).Name("contact_api:DeleteUriTypeHandlerV1")

	// Contact titles
	r.HandleFunc("/v1/contactTitle", api.GetContactTitlesHandlerV1).Methods(http.MethodGet).Name("contact_api:GetContactTitlesHandlerV1")
	r.HandleFunc("/v1/contactTitle/{id}", api.GetContactTitleByIdHandlerV1).Methods(http.MethodGet).Name("contact_api:GetContactTitleByIdHandlerV1")
	r.HandleFunc("/v1/contactTitle", api.CreateContactTitleHandlerV1).Methods(http.MethodPost).Name("contact_api:CreateContactTitleHandlerV1")
	r.HandleFunc("/v1/contactTitle/{id}", api.UpdateContactTitleHandlerV1).Methods(http.MethodPut).Name("contact_api:UpdateContactTitleHandlerV1")
	r.HandleFunc("/v1/contactTitle/{id}", api.DeleteContactTitleHandlerV1).Methods(http.MethodDelete).Name("contact_api:DeleteContactTitleHandlerV1")

	// Company types
	r.HandleFunc("/v1/companyType", api.GetCompanyTypesHandlerV1).Methods(http.MethodGet).Name("contact_api:GetCompanyTypesHandlerV1")
	r.HandleFunc("/v1/companyType/{id}", api.GetCompanyTypeByIdHandlerV1).Methods(http.MethodGet).Name("contact_api:GetCompanyTypeByIdHandlerV1")
	r.HandleFunc("/v1/companyType", api.CreateCompanyTypeHandlerV1).Methods(http.MethodPost).Name("contact_api:CreateCompanyTypeHandlerV1")
	r.HandleFunc("/v1/companyType/{id}", api.UpdateCompanyTypeHandlerV1).Methods(http.MethodPut).Name("contact_api:UpdateCompanyTypeHandlerV1")
	r.HandleFunc("/v1/companyType/{id}", api.DeleteCompanyTypeHandlerV1).Methods(http.MethodDelete).Name("contact_api:DeleteCompanyTypeHandlerV1")

	// Industries
	r.HandleFunc("/v1/industry", api.GetIndustriesHandlerV1).Methods(http.MethodGet).Name("contact_api:GetIndustriesHandlerV1")
	r.HandleFunc("/v1/industry/{id}", api.GetIndustryByIdHandlerV1).Methods(http.MethodGet).Name("contact_api:GetIndustryByIdHandlerV1")
	r.HandleFunc("/v1/industry", api.CreateIndustryHandlerV1).Methods(http.MethodPost).Name("contact_api:CreateIndustryHandlerV1")
	r.HandleFunc("/v1/industry/{id}", api.UpdateIndustryHandlerV1).Methods(http.MethodPut).Name("contact_api:UpdateIndustryHandlerV1")
	r.HandleFunc("/v1/industry/{id}", api.DeleteIndustryHandlerV1).Methods(http.MethodDelete).Name("contact_api:DeleteIndustryHandlerV1")

	// Job titles
	r.HandleFunc("/v1/jobTitle", api.GetJobTitlesHandlerV1).Methods(http.MethodGet).Name("contact_api:GetJobTitlesHandlerV1")
	r.HandleFunc("/v1/jobTitle/{id}", api.GetJobTitleByIdHandlerV1).Methods(http.MethodGet).Name("contact_api:GetJobTitleByIdHandlerV1")
	r.HandleFunc("/v1/jobTitle", api.CreateJobTitleHandlerV1).Methods(http.MethodPost).Name("contact_api:CreateJobTitleHandlerV1")
	r.HandleFunc("/v1/jobTitle/{id}", api.UpdateJobTitleHandlerV1).Methods(http.MethodPut).Name("contact_api:UpdateJobTitleHandlerV1")
	r.HandleFunc("/v1/jobTitle/{id}", api.DeleteJobTitleHandlerV1).Methods(http.MethodDelete).Name("contact_api:DeleteJobTitleHandlerV1")
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
