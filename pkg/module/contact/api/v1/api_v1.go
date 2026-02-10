package v1

import (
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact"
	"github.com/deb-ict/go-router"
	"github.com/deb-ict/go-router/authorization"
)

const (
	PolicyContactReadV1    = "contact_api:ReadContact:v1"
	PolicyContactCreateV1  = "contact_api:CreateContact:v1"
	PolicyContactUpdateV1  = "contact_api:UpdateContact:v1"
	PolicyContactDeleteV1  = "contact_api:DeleteContact:v1"
	PolicyCompanyReadV1    = "contact_api:ReadCompany:v1"
	PolicyCompanyCreateV1  = "contact_api:CreateCompany:v1"
	PolicyCompanyUpdateV1  = "contact_api:UpdateCompany:v1"
	PolicyCompanyDeleteV1  = "contact_api:DeleteCompany:v1"
	PolicyMetadataReadV1   = "contact_api:ReadMetadata:v1"
	PolicyMetadataCreateV1 = "contact_api:CreateMetadata:v1"
	PolicyMetadataUpdateV1 = "contact_api:UpdateMetadata:v1"
	PolicyMetadataDeleteV1 = "contact_api:DeleteMetadata:v1"
)

type ApiV1 interface {
	RegisterAuthorizationPolicies(middleware *authorization.Middleware)
	RegisterRoutes(r *router.Router)
}

type apiV1 struct {
	service contact.Service
}

func NewApiV1(service contact.Service) ApiV1 {
	return &apiV1{
		service: service,
	}
}

func (api *apiV1) RegisterAuthorizationPolicies(middleware *authorization.Middleware) {
	middleware.SetPolicy(authorization.NewPolicy(PolicyContactReadV1,
		authorization.NewScopeRequirement("contact.read"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyContactCreateV1,
		authorization.NewScopeRequirement("contact.create"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyContactUpdateV1,
		authorization.NewScopeRequirement("contact.update"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyContactDeleteV1,
		authorization.NewScopeRequirement("contact.delete"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyCompanyReadV1,
		authorization.NewScopeRequirement("company.read"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyCompanyCreateV1,
		authorization.NewScopeRequirement("company.create"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyCompanyUpdateV1,
		authorization.NewScopeRequirement("company.update"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyCompanyDeleteV1,
		authorization.NewScopeRequirement("company.delete"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyMetadataReadV1,
		authorization.NewScopeRequirement("contact.metadata.read"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyMetadataCreateV1,
		authorization.NewScopeRequirement("contact.metadata.create"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyMetadataUpdateV1,
		authorization.NewScopeRequirement("contact.metadata.update"),
	))
	middleware.SetPolicy(authorization.NewPolicy(PolicyMetadataDeleteV1,
		authorization.NewScopeRequirement("contact.metadata.delete"),
	))
}

func (api *apiV1) RegisterRoutes(r *router.Router) {
	// Contacts
	r.HandleFunc("/v1/contact", api.GetContactsHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyContactReadV1),
	)
	r.HandleFunc("/v1/contact/{id}", api.GetContactByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyContactReadV1),
	)
	r.HandleFunc("/v1/contact", api.CreateContactHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyContactCreateV1),
	)
	r.HandleFunc("/v1/contact/{id}", api.UpdateContactHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyContactUpdateV1),
	)
	r.HandleFunc("/v1/contact/{id}", api.DeleteContactHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyContactDeleteV1),
	)

	// Contact addresses
	r.HandleFunc("/v1/contact/{contactId}/address", api.GetContactAddressesHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyContactReadV1),
	)
	r.HandleFunc("/v1/contact/{contactId}/address/{id}", api.GetContactAddressByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyContactReadV1),
	)
	r.HandleFunc("/v1/contact/{contactId}/address", api.CreateContactAddressHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyContactCreateV1),
	)
	r.HandleFunc("/v1/contact/{contactId}/address/{id}", api.UpdateContactAddressHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyContactUpdateV1),
	)
	r.HandleFunc("/v1/contact/{contactId}/address/{id}", api.DeleteContactAddressHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyContactDeleteV1),
	)

	// Contact emails
	r.HandleFunc("/v1/contact/{contactId}/email", api.GetContactEmailsHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyContactReadV1),
	)
	r.HandleFunc("/v1/contact/{contactId}/email/{id}", api.GetContactEmailByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyContactReadV1),
	)
	r.HandleFunc("/v1/contact/{contactId}/email", api.CreateContactEmailHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyContactCreateV1),
	)
	r.HandleFunc("/v1/contact/{contactId}/email/{id}", api.UpdateContactEmailHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyContactUpdateV1),
	)
	r.HandleFunc("/v1/contact/{contactId}/email/{id}", api.DeleteContactEmailHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyContactDeleteV1),
	)

	// Contact phones
	r.HandleFunc("/v1/contact/{contactId}/phone", api.GetContactPhonesHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyContactReadV1),
	)
	r.HandleFunc("/v1/contact/{contactId}/phone/{id}", api.GetContactPhoneByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyContactReadV1),
	)
	r.HandleFunc("/v1/contact/{contactId}/phone", api.CreateContactPhoneHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyContactCreateV1),
	)
	r.HandleFunc("/v1/contact/{contactId}/phone/{id}", api.UpdateContactPhoneHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyContactUpdateV1),
	)
	r.HandleFunc("/v1/contact/{contactId}/phone/{id}", api.DeleteContactPhoneHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyContactDeleteV1),
	)

	// Contact uris
	r.HandleFunc("/v1/contact/{contactId}/uri", api.GetContactUrisHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyContactReadV1),
	)
	r.HandleFunc("/v1/contact/{contactId}/uri/{id}", api.GetContactUriByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyContactReadV1),
	)
	r.HandleFunc("/v1/contact/{contactId}/uri", api.CreateContactUriHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyContactCreateV1),
	)
	r.HandleFunc("/v1/contact/{contactId}/uri/{id}", api.UpdateContactUriHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyContactUpdateV1),
	)
	r.HandleFunc("/v1/contact/{contactId}/uri/{id}", api.DeleteContactUriHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyContactDeleteV1),
	)

	// Companies
	r.HandleFunc("/v1/company", api.GetCompaniesHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyCompanyReadV1),
	)
	r.HandleFunc("/v1/company/{id}", api.GetCompanyByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyCompanyReadV1),
	)
	r.HandleFunc("/v1/company", api.CreateCompanyHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyCompanyCreateV1),
	)
	r.HandleFunc("/v1/company/{id}", api.UpdateCompanyHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyCompanyUpdateV1),
	)
	r.HandleFunc("/v1/company/{id}", api.DeleteCompanyHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyCompanyDeleteV1),
	)

	// Company addresses
	r.HandleFunc("/v1/company/{companyId}/address", api.GetCompanyAddressesHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyCompanyReadV1),
	)
	r.HandleFunc("/v1/company/{companyId}/address/{id}", api.GetCompanyAddressByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyCompanyReadV1),
	)
	r.HandleFunc("/v1/company/{companyId}/address", api.CreateCompanyAddressHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyCompanyCreateV1),
	)
	r.HandleFunc("/v1/company/{companyId}/address/{id}", api.UpdateCompanyAddressHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyCompanyUpdateV1),
	)
	r.HandleFunc("/v1/company/{companyId}/address/{id}", api.DeleteCompanyAddressHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyCompanyDeleteV1),
	)

	// Company emails
	r.HandleFunc("/v1/company/{companyId}/email", api.GetCompanyEmailsHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyCompanyReadV1),
	)
	r.HandleFunc("/v1/company/{companyId}/email/{id}", api.GetCompanyEmailByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyCompanyReadV1),
	)
	r.HandleFunc("/v1/company/{companyId}/email", api.CreateCompanyEmailHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyCompanyCreateV1),
	)
	r.HandleFunc("/v1/company/{companyId}/email/{id}", api.UpdateCompanyEmailHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyCompanyUpdateV1),
	)
	r.HandleFunc("/v1/company/{companyId}/email/{id}", api.DeleteCompanyEmailHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyCompanyDeleteV1),
	)

	// Company phones
	r.HandleFunc("/v1/company/{companyId}/phone", api.GetCompanyPhonesHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyCompanyReadV1),
	)
	r.HandleFunc("/v1/company/{companyId}/phone/{id}", api.GetCompanyPhoneByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyCompanyReadV1),
	)
	r.HandleFunc("/v1/company/{companyId}/phone", api.CreateCompanyPhoneHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyCompanyCreateV1),
	)
	r.HandleFunc("/v1/company/{companyId}/phone/{id}", api.UpdateCompanyPhoneHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyCompanyUpdateV1),
	)
	r.HandleFunc("/v1/company/{companyId}/phone/{id}", api.DeleteCompanyPhoneHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyCompanyDeleteV1),
	)

	// Company uris
	r.HandleFunc("/v1/company/{companyId}/uri", api.GetCompanyUrisHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyCompanyReadV1),
	)
	r.HandleFunc("/v1/company/{companyId}/uri/{id}", api.GetCompanyUriByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyCompanyReadV1),
	)
	r.HandleFunc("/v1/company/{companyId}/uri", api.CreateCompanyUriHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyCompanyCreateV1),
	)
	r.HandleFunc("/v1/company/{companyId}/uri/{id}", api.UpdateCompanyUriHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyCompanyUpdateV1),
	)
	r.HandleFunc("/v1/company/{companyId}/uri/{id}", api.DeleteCompanyUriHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyCompanyDeleteV1),
	)

	// Address types
	r.HandleFunc("/v1/addressType", api.GetAddressTypesHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyMetadataReadV1),
	)
	r.HandleFunc("/v1/addressType/{id}", api.GetAddressTypeByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyMetadataReadV1),
	)
	r.HandleFunc("/v1/emailType", api.CreateEmailTypeHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyMetadataCreateV1),
	)
	r.HandleFunc("/v1/emailType/{id}", api.GetEmailTypeByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyMetadataReadV1),
	)
	r.HandleFunc("/v1/emailType/{id}", api.UpdateEmailTypeHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyMetadataUpdateV1),
	)
	r.HandleFunc("/v1/emailType/{id}", api.DeleteEmailTypeHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyMetadataDeleteV1),
	)

	// Phone types
	r.HandleFunc("/v1/phoneType", api.GetPhoneTypesHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyMetadataReadV1),
	)
	r.HandleFunc("/v1/phoneType/{id}", api.GetPhoneTypeByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyMetadataReadV1),
	)
	r.HandleFunc("/v1/phoneType", api.CreatePhoneTypeHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyMetadataCreateV1),
	)
	r.HandleFunc("/v1/phoneType/{id}", api.UpdatePhoneTypeHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyMetadataUpdateV1),
	)
	r.HandleFunc("/v1/phoneType/{id}", api.DeletePhoneTypeHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyMetadataDeleteV1),
	)

	// Uri types
	r.HandleFunc("/v1/uriType", api.GetUriTypesHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyMetadataReadV1),
	)
	r.HandleFunc("/v1/uriType/{id}", api.GetUriTypeByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyMetadataReadV1),
	)
	r.HandleFunc("/v1/uriType", api.CreateUriTypeHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyMetadataCreateV1),
	)
	r.HandleFunc("/v1/uriType/{id}", api.UpdateUriTypeHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyMetadataUpdateV1),
	)
	r.HandleFunc("/v1/uriType/{id}", api.DeleteUriTypeHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyMetadataDeleteV1),
	)

	// Contact titles
	r.HandleFunc("/v1/contactTitle", api.GetContactTitlesHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyMetadataReadV1),
	)
	r.HandleFunc("/v1/contactTitle/{id}", api.GetContactTitleByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyMetadataReadV1),
	)
	r.HandleFunc("/v1/contactTitle", api.CreateContactTitleHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyMetadataCreateV1),
	)
	r.HandleFunc("/v1/contactTitle/{id}", api.UpdateContactTitleHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyMetadataUpdateV1),
	)
	r.HandleFunc("/v1/contactTitle/{id}", api.DeleteContactTitleHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyMetadataDeleteV1),
	)

	// Company types
	r.HandleFunc("/v1/companyType", api.GetCompanyTypesHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyMetadataReadV1),
	)
	r.HandleFunc("/v1/companyType/{id}", api.GetCompanyTypeByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyMetadataReadV1),
	)
	r.HandleFunc("/v1/companyType", api.CreateCompanyTypeHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyMetadataCreateV1),
	)
	r.HandleFunc("/v1/companyType/{id}", api.UpdateCompanyTypeHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyMetadataUpdateV1),
	)
	r.HandleFunc("/v1/companyType/{id}", api.DeleteCompanyTypeHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyMetadataDeleteV1),
	)

	// Industries
	r.HandleFunc("/v1/industry", api.GetIndustriesHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyMetadataReadV1),
	)
	r.HandleFunc("/v1/industry/{id}", api.GetIndustryByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyMetadataReadV1),
	)
	r.HandleFunc("/v1/industry", api.CreateIndustryHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyMetadataCreateV1),
	)
	r.HandleFunc("/v1/industry/{id}", api.UpdateIndustryHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyMetadataUpdateV1),
	)
	r.HandleFunc("/v1/industry/{id}", api.DeleteIndustryHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyMetadataDeleteV1),
	)

	// Job titles
	r.HandleFunc("/v1/jobTitle", api.GetJobTitlesHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyMetadataReadV1),
	)
	r.HandleFunc("/v1/jobTitle/{id}", api.GetJobTitleByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized(PolicyMetadataReadV1),
	)
	r.HandleFunc("/v1/jobTitle", api.CreateJobTitleHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized(PolicyMetadataCreateV1),
	)
	r.HandleFunc("/v1/jobTitle/{id}", api.UpdateJobTitleHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized(PolicyMetadataUpdateV1),
	)
	r.HandleFunc("/v1/jobTitle/{id}", api.DeleteJobTitleHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized(PolicyMetadataDeleteV1),
	)
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
