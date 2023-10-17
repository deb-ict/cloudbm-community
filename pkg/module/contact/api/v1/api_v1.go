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

func NewApi(service contact.Service) ApiV1 {
	return &apiV1{
		service: service,
	}
}

func (api *apiV1) RegisterRoutes(r *mux.Router) {
	// Contacts
	r.HandleFunc("/v1/contact", api.GetContactsHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/contact/{id}", api.GetContactByIdHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/contact", api.CreateContactHandlerV1).Methods(http.MethodPost)
	r.HandleFunc("/v1/contact/{id}", api.UpdateContactHandlerV1).Methods(http.MethodPut)
	r.HandleFunc("/v1/contact/{id}", api.DeleteContactHandlerV1).Methods(http.MethodDelete)

	// Contact addresses
	r.HandleFunc("/v1/contact/{contactId}/address", api.GetContactAddressesHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/contact/{contactId}/address/{id}", api.GetContactAddressByIdHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/contact/{contactId}/address", api.CreateContactAddressHandlerV1).Methods(http.MethodPost)
	r.HandleFunc("/v1/contact/{contactId}/address/{id}", api.UpdateContactAddressHandlerV1).Methods(http.MethodPut)
	r.HandleFunc("/v1/contact/{contactId}/address/{id}", api.DeleteContactAddressHandlerV1).Methods(http.MethodDelete)

	// Contact emails
	r.HandleFunc("/v1/contact/{contactId}/email", api.GetContactEmailsHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/contact/{contactId}/email/{id}", api.GetContactEmailByIdHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/contact/{contactId}/email", api.CreateContactEmailHandlerV1).Methods(http.MethodPost)
	r.HandleFunc("/v1/contact/{contactId}/email/{id}", api.UpdateContactEmailHandlerV1).Methods(http.MethodPut)
	r.HandleFunc("/v1/contact/{contactId}/email/{id}", api.DeleteContactEmailHandlerV1).Methods(http.MethodDelete)

	// Contact phones
	r.HandleFunc("/v1/contact/{contactId}/phone", api.GetContactPhonesHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/contact/{contactId}/phone/{id}", api.GetContactPhoneByIdHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/contact/{contactId}/phone", api.CreateContactPhoneHandlerV1).Methods(http.MethodPost)
	r.HandleFunc("/v1/contact/{contactId}/phone/{id}", api.UpdateContactPhoneHandlerV1).Methods(http.MethodPut)
	r.HandleFunc("/v1/contact/{contactId}/phone/{id}", api.DeleteContactPhoneHandlerV1).Methods(http.MethodDelete)

	// Contact uris
	r.HandleFunc("/v1/contact/{contactId}/uri", api.GetContactUrisHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/contact/{contactId}/uri/{id}", api.GetContactUriByIdHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/contact/{contactId}/uri", api.CreateContactUriHandlerV1).Methods(http.MethodPost)
	r.HandleFunc("/v1/contact/{contactId}/uri/{id}", api.UpdateContactUriHandlerV1).Methods(http.MethodPut)
	r.HandleFunc("/v1/contact/{contactId}/uri/{id}", api.DeleteContactUriHandlerV1).Methods(http.MethodDelete)

	// Companies
	r.HandleFunc("/v1/company", api.GetCompaniesHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/company/{id}", api.GetCompanyByIdHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/company", api.CreateCompanyHandlerV1).Methods(http.MethodPost)
	r.HandleFunc("/v1/company/{id}", api.UpdateCompanyHandlerV1).Methods(http.MethodPut)
	r.HandleFunc("/v1/company/{id}", api.DeleteCompanyHandlerV1).Methods(http.MethodDelete)

	// Company addresses
	r.HandleFunc("/v1/company/{companyId}/address", api.GetCompanyAddressesHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/company/{companyId}/address/{id}", api.GetCompanyAddressByIdHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/company/{companyId}/address", api.CreateCompanyAddressHandlerV1).Methods(http.MethodPost)
	r.HandleFunc("/v1/company/{companyId}/address/{id}", api.UpdateCompanyAddressHandlerV1).Methods(http.MethodPut)
	r.HandleFunc("/v1/company/{companyId}/address/{id}", api.DeleteCompanyAddressHandlerV1).Methods(http.MethodDelete)

	// Company emails
	r.HandleFunc("/v1/company/{companyId}/email", api.GetCompanyEmailsHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/company/{companyId}/email/{id}", api.GetCompanyEmailByIdHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/company/{companyId}/email", api.CreateCompanyEmailHandlerV1).Methods(http.MethodPost)
	r.HandleFunc("/v1/company/{companyId}/email/{id}", api.UpdateCompanyEmailHandlerV1).Methods(http.MethodPut)
	r.HandleFunc("/v1/company/{companyId}/email/{id}", api.DeleteCompanyEmailHandlerV1).Methods(http.MethodDelete)

	// Company phones
	r.HandleFunc("/v1/company/{companyId}/phone", api.GetCompanyPhonesHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/company/{companyId}/phone/{id}", api.GetCompanyByIdHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/company/{companyId}/phone", api.CreateCompanyPhoneHandlerV1).Methods(http.MethodPost)
	r.HandleFunc("/v1/company/{companyId}/phone/{id}", api.UpdateCompanyPhoneHandlerV1).Methods(http.MethodPut)
	r.HandleFunc("/v1/company/{companyId}/phone/{id}", api.DeleteCompanyPhoneHandlerV1).Methods(http.MethodDelete)

	// Company uris
	r.HandleFunc("/v1/company/{companyId}/uri", api.GetCompanyUrisHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/company/{companyId}/uri/{id}", api.GetCompanyUriByIdHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/company/{companyId}/uri", api.CreateCompanyUriHandlerV1).Methods(http.MethodPost)
	r.HandleFunc("/v1/company/{companyId}/uri/{id}", api.UpdateCompanyUriHandlerV1).Methods(http.MethodPut)
	r.HandleFunc("/v1/company/{companyId}/uri/{id}", api.DeleteCompanyUriHandlerV1).Methods(http.MethodDelete)

	// Address types
	r.HandleFunc("/v1/addressType", api.GetAddressTypesHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/addressType/{id}", api.GetAddressTypeByIdHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/addressType", api.CreateAddressTypeHandlerV1).Methods(http.MethodPost)
	r.HandleFunc("/v1/addressType/{id}", api.UpdateAddressTypeHandlerV1).Methods(http.MethodPut)
	r.HandleFunc("/v1/addressType/{id}", api.DeleteAddressTypeHandlerV1).Methods(http.MethodDelete)

	// Email types
	r.HandleFunc("/v1/emailType", api.GetEmailTypesHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/emailType/{id}", api.GetEmailTypeByIdHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/emailType", api.CreateEmailTypeHandlerV1).Methods(http.MethodPost)
	r.HandleFunc("/v1/emailType/{id}", api.UpdateEmailTypeHandlerV1).Methods(http.MethodPut)
	r.HandleFunc("/v1/emailType/{id}", api.DeleteEmailTypeHandlerV1).Methods(http.MethodDelete)

	// Phone types
	r.HandleFunc("/v1/phoneType", api.GetPhoneTypesHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/phoneType/{id}", api.GetPhoneTypeByIdHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/phoneType", api.CreatePhoneTypeHandlerV1).Methods(http.MethodPost)
	r.HandleFunc("/v1/phoneType/{id}", api.UpdatePhoneTypeHandlerV1).Methods(http.MethodPut)
	r.HandleFunc("/v1/phoneType/{id}", api.DeletePhoneTypeHandlerV1).Methods(http.MethodDelete)

	// Uri types
	r.HandleFunc("/v1/uriType", api.GetUriTypesHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/uriType/{id}", api.GetUriTypeByIdHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/uriType", api.CreateUriTypeHandlerV1).Methods(http.MethodPost)
	r.HandleFunc("/v1/uriType/{id}", api.UpdateUriTypeHandlerV1).Methods(http.MethodPut)
	r.HandleFunc("/v1/uriType/{id}", api.DeleteUriTypeHandlerV1).Methods(http.MethodDelete)

	// Contact titles
	r.HandleFunc("/v1/contactTitle", api.GetContactTitlesHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/contactTitle/{id}", api.GetContactTitleByIdHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/contactTitle", api.CreateContactTitleHandlerV1).Methods(http.MethodPost)
	r.HandleFunc("/v1/contactTitle/{id}", api.UpdateContactTitleHandlerV1).Methods(http.MethodPut)
	r.HandleFunc("/v1/contactTitle/{id}", api.DeleteContactTitleHandlerV1).Methods(http.MethodDelete)

	// Company types
	r.HandleFunc("/v1/companyType", api.GetCompanyTypesHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/companyType/{id}", api.GetCompanyTypeByIdHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/companyType", api.CreateCompanyTypeHandlerV1).Methods(http.MethodPost)
	r.HandleFunc("/v1/companyType/{id}", api.UpdateCompanyTypeHandlerV1).Methods(http.MethodPut)
	r.HandleFunc("/v1/companyType/{id}", api.DeleteCompanyTypeHandlerV1).Methods(http.MethodDelete)

	// Industries
	r.HandleFunc("/v1/industry", api.GetIndustriesHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/industry/{id}", api.GetIndustryByIdHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/industry", api.CreateIndustryHandlerV1).Methods(http.MethodPost)
	r.HandleFunc("/v1/industry/{id}", api.UpdateIndustryHandlerV1).Methods(http.MethodPut)
	r.HandleFunc("/v1/industry/{id}", api.DeleteIndustryHandlerV1).Methods(http.MethodDelete)

	// Job titles
	r.HandleFunc("/v1/jobTitle", api.GetJobTitlesHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/jobTitle/{id}", api.GetJobTitleByIdHandlerV1).Methods(http.MethodGet)
	r.HandleFunc("/v1/jobTitle", api.CreateJobTitleHandlerV1).Methods(http.MethodPost)
	r.HandleFunc("/v1/jobTitle/{id}", api.UpdateJobTitleHandlerV1).Methods(http.MethodPut)
	r.HandleFunc("/v1/jobTitle/{id}", api.DeleteJobTitleHandlerV1).Methods(http.MethodDelete)
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
