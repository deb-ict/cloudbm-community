package contact

import (
	"net/http"

	"github.com/gorilla/mux"
)

type ApiHandler interface {
	GetService() Service
	RegisterRoutes(r *mux.Router)

	GetAddressTypes(w http.ResponseWriter, r *http.Request)
	GetAddressTypeById(w http.ResponseWriter, r *http.Request)
	CreateAddressType(w http.ResponseWriter, r *http.Request)
	UpdateAddressType(w http.ResponseWriter, r *http.Request)
	DeleteAddressType(w http.ResponseWriter, r *http.Request)

	GetPhoneTypes(w http.ResponseWriter, r *http.Request)
	GetPhoneTypeById(w http.ResponseWriter, r *http.Request)
	CreatePhoneType(w http.ResponseWriter, r *http.Request)
	UpdatePhoneType(w http.ResponseWriter, r *http.Request)
	DeletePhoneType(w http.ResponseWriter, r *http.Request)

	GetEmailTypes(w http.ResponseWriter, r *http.Request)
	GetEmailTypeById(w http.ResponseWriter, r *http.Request)
	CreateEmailType(w http.ResponseWriter, r *http.Request)
	UpdateEmailType(w http.ResponseWriter, r *http.Request)
	DeleteEmailType(w http.ResponseWriter, r *http.Request)

	GetUrlTypes(w http.ResponseWriter, r *http.Request)
	GetUrlTypeById(w http.ResponseWriter, r *http.Request)
	CreateUrlType(w http.ResponseWriter, r *http.Request)
	UpdateUrlType(w http.ResponseWriter, r *http.Request)
	DeleteUrlType(w http.ResponseWriter, r *http.Request)

	GetContacts(w http.ResponseWriter, r *http.Request)
	GetContactById(w http.ResponseWriter, r *http.Request)
	CreateContact(w http.ResponseWriter, r *http.Request)
	UpdateContact(w http.ResponseWriter, r *http.Request)
	DeleteContact(w http.ResponseWriter, r *http.Request)

	GetContactAddresses(w http.ResponseWriter, r *http.Request)
	GetContactAddressById(w http.ResponseWriter, r *http.Request)
	CreateContactAddress(w http.ResponseWriter, r *http.Request)
	UpdateContactAddress(w http.ResponseWriter, r *http.Request)
	DeleteContactAddress(w http.ResponseWriter, r *http.Request)

	GetContactPhones(w http.ResponseWriter, r *http.Request)
	GetContactPhoneById(w http.ResponseWriter, r *http.Request)
	CreateContactPhone(w http.ResponseWriter, r *http.Request)
	UpdateContactPhone(w http.ResponseWriter, r *http.Request)
	DeleteContactPhone(w http.ResponseWriter, r *http.Request)

	GetContactEmails(w http.ResponseWriter, r *http.Request)
	GetContactEmailById(w http.ResponseWriter, r *http.Request)
	CreateContactEmail(w http.ResponseWriter, r *http.Request)
	UpdateContactEmail(w http.ResponseWriter, r *http.Request)
	DeleteContactEmail(w http.ResponseWriter, r *http.Request)

	GetCompanies(w http.ResponseWriter, r *http.Request)
	GetCompanyById(w http.ResponseWriter, r *http.Request)
	CreateCompany(w http.ResponseWriter, r *http.Request)
	UpdateCompany(w http.ResponseWriter, r *http.Request)
	DeleteCompany(w http.ResponseWriter, r *http.Request)

	GetCompanyAddresses(w http.ResponseWriter, r *http.Request)
	GetCompanyAddressById(w http.ResponseWriter, r *http.Request)
	CreateCompanyAddress(w http.ResponseWriter, r *http.Request)
	UpdateCompanyAddress(w http.ResponseWriter, r *http.Request)
	DeleteCompanyAddress(w http.ResponseWriter, r *http.Request)

	GetCompanyPhones(w http.ResponseWriter, r *http.Request)
	GetCompanyPhoneById(w http.ResponseWriter, r *http.Request)
	CreateCompanyPhone(w http.ResponseWriter, r *http.Request)
	UpdateCompanyPhone(w http.ResponseWriter, r *http.Request)
	DeleteCompanyPhone(w http.ResponseWriter, r *http.Request)

	GetCompanyEmails(w http.ResponseWriter, r *http.Request)
	GetCompanyEmailById(w http.ResponseWriter, r *http.Request)
	CreateCompanyEmail(w http.ResponseWriter, r *http.Request)
	UpdateCompanyEmail(w http.ResponseWriter, r *http.Request)
	DeleteCompanyEmail(w http.ResponseWriter, r *http.Request)

	GetCompanyUrls(w http.ResponseWriter, r *http.Request)
	GetCompanyUrlById(w http.ResponseWriter, r *http.Request)
	CreateCompanyUrl(w http.ResponseWriter, r *http.Request)
	UpdateCompanyUrl(w http.ResponseWriter, r *http.Request)
	DeleteCompanyUrl(w http.ResponseWriter, r *http.Request)
}

func newApiHandler(svc Service) ApiHandler {
	api := &apiHandler{
		service: svc,
	}
	return api
}

type apiHandler struct {
	service Service
}

func (api *apiHandler) GetService() Service {
	return api.service
}

func (api *apiHandler) RegisterRoutes(r *mux.Router) {
	api.registerV1Routes(r)
}

func (api *apiHandler) registerV1Routes(r *mux.Router) {
	r.HandleFunc("/v1/address_type", api.GetAddressTypes).Methods(http.MethodGet)
	r.HandleFunc("/v1/address_type/{id}", api.GetAddressTypeById).Methods(http.MethodGet)
	r.HandleFunc("/v1/address_type", api.CreateAddressType).Methods(http.MethodPost)
	r.HandleFunc("/v1/address_type/{id}", api.UpdateAddressType).Methods(http.MethodPut)
	r.HandleFunc("/v1/address_type/{id}", api.DeleteAddressType).Methods(http.MethodDelete)

	r.HandleFunc("/v1/phone_type", api.GetPhoneTypes).Methods(http.MethodGet)
	r.HandleFunc("/v1/phone_type/{id}", api.GetPhoneTypeById).Methods(http.MethodGet)
	r.HandleFunc("/v1/phone_type", api.CreatePhoneType).Methods(http.MethodPost)
	r.HandleFunc("/v1/phone_type/{id}", api.UpdatePhoneType).Methods(http.MethodPut)
	r.HandleFunc("/v1/phone_type/{id}", api.DeletePhoneType).Methods(http.MethodDelete)

	r.HandleFunc("/v1/email_type", api.GetEmailTypes).Methods(http.MethodGet)
	r.HandleFunc("/v1/email_type/{id}", api.GetEmailTypeById).Methods(http.MethodGet)
	r.HandleFunc("/v1/email_type", api.CreateEmailType).Methods(http.MethodPost)
	r.HandleFunc("/v1/email_type/{id}", api.UpdateEmailType).Methods(http.MethodPut)
	r.HandleFunc("/v1/email_type/{id}", api.DeleteEmailType).Methods(http.MethodDelete)

	r.HandleFunc("/v1/url_type", api.GetUrlTypes).Methods(http.MethodGet)
	r.HandleFunc("/v1/url_type/{id}", api.GetUrlTypeById).Methods(http.MethodGet)
	r.HandleFunc("/v1/url_type", api.CreateUrlType).Methods(http.MethodPost)
	r.HandleFunc("/v1/url_type/{id}", api.UpdateUrlType).Methods(http.MethodPut)
	r.HandleFunc("/v1/url_type/{id}", api.DeleteUrlType).Methods(http.MethodDelete)

	r.HandleFunc("/v1/contact", api.GetContacts).Methods(http.MethodGet)
	r.HandleFunc("/v1/contact/{id}", api.GetContactById).Methods(http.MethodGet)
	r.HandleFunc("/v1/contact", api.CreateContact).Methods(http.MethodPost)
	r.HandleFunc("/v1/contact/{id}", api.UpdateContact).Methods(http.MethodPut)
	r.HandleFunc("/v1/contact/{id}", api.DeleteContact).Methods(http.MethodDelete)

	r.HandleFunc("/v1/contact/{contactId}/address", api.GetContactAddresses).Methods(http.MethodGet)
	r.HandleFunc("/v1/contact/{contactId}/address/{id}", api.GetContactAddressById).Methods(http.MethodGet)
	r.HandleFunc("/v1/contact/{contactId}/address", api.CreateContactAddress).Methods(http.MethodPost)
	r.HandleFunc("/v1/contact/{contactId}/address/{id}", api.UpdateContactAddress).Methods(http.MethodPut)
	r.HandleFunc("/v1/contact/{contactId}/address/{id}", api.DeleteContactAddress).Methods(http.MethodDelete)

	r.HandleFunc("/v1/contact/{contactId}/phone", api.GetContactPhones).Methods(http.MethodGet)
	r.HandleFunc("/v1/contact/{contactId}/phone/{id}", api.GetContactPhoneById).Methods(http.MethodGet)
	r.HandleFunc("/v1/contact/{contactId}/phone", api.CreateContactPhone).Methods(http.MethodPost)
	r.HandleFunc("/v1/contact/{contactId}/phone/{id}", api.UpdateContactPhone).Methods(http.MethodPut)
	r.HandleFunc("/v1/contact/{contactId}/phone/{id}", api.DeleteContactPhone).Methods(http.MethodDelete)

	r.HandleFunc("/v1/contact/{contactId}/email", api.GetContactEmails).Methods(http.MethodGet)
	r.HandleFunc("/v1/contact/{contactId}/email/{id}", api.GetContactEmailById).Methods(http.MethodGet)
	r.HandleFunc("/v1/contact/{contactId}/email", api.CreateContactEmail).Methods(http.MethodPost)
	r.HandleFunc("/v1/contact/{contactId}/email/{id}", api.UpdateContactEmail).Methods(http.MethodPut)
	r.HandleFunc("/v1/contact/{contactId}/email/{id}", api.DeleteContactEmail).Methods(http.MethodDelete)

	r.HandleFunc("/v1/company", api.GetCompanies).Methods(http.MethodGet)
	r.HandleFunc("/v1/company/{id}", api.GetCompanyById).Methods(http.MethodGet)
	r.HandleFunc("/v1/company", api.CreateCompany).Methods(http.MethodPost)
	r.HandleFunc("/v1/company/{id}", api.UpdateCompany).Methods(http.MethodPut)
	r.HandleFunc("/v1/company/{id}", api.DeleteCompany).Methods(http.MethodDelete)

	r.HandleFunc("/v1/company/{companyId}/address", api.GetCompanyAddresses).Methods(http.MethodGet)
	r.HandleFunc("/v1/company/{companyId}/address/{id}", api.GetCompanyAddressById).Methods(http.MethodGet)
	r.HandleFunc("/v1/company/{companyId}/address", api.CreateCompanyAddress).Methods(http.MethodPost)
	r.HandleFunc("/v1/company/{companyId}/address/{id}", api.UpdateCompanyAddress).Methods(http.MethodPut)
	r.HandleFunc("/v1/company/{companyId}/address/{id}", api.DeleteCompanyAddress).Methods(http.MethodDelete)

	r.HandleFunc("/v1/company/{companyId}/phone", api.GetCompanyPhones).Methods(http.MethodGet)
	r.HandleFunc("/v1/company/{companyId}/phone/{id}", api.GetCompanyPhoneById).Methods(http.MethodGet)
	r.HandleFunc("/v1/company/{companyId}/phone", api.CreateCompanyPhone).Methods(http.MethodPost)
	r.HandleFunc("/v1/company/{companyId}/phone/{id}", api.UpdateCompanyPhone).Methods(http.MethodPut)
	r.HandleFunc("/v1/company/{companyId}/phone/{id}", api.DeleteCompanyPhone).Methods(http.MethodDelete)

	r.HandleFunc("/v1/company/{companyId}/email", api.GetCompanyEmails).Methods(http.MethodGet)
	r.HandleFunc("/v1/company/{companyId}/email/{id}", api.GetCompanyEmailById).Methods(http.MethodGet)
	r.HandleFunc("/v1/company/{companyId}/email", api.CreateCompanyEmail).Methods(http.MethodPost)
	r.HandleFunc("/v1/company/{companyId}/email/{id}", api.UpdateCompanyEmail).Methods(http.MethodPut)
	r.HandleFunc("/v1/company/{companyId}/email/{id}", api.DeleteCompanyEmail).Methods(http.MethodDelete)

	r.HandleFunc("/v1/company/{companyId}/url", api.GetCompanyUrls).Methods(http.MethodGet)
	r.HandleFunc("/v1/company/{companyId}/url/{id}", api.GetCompanyUrlById).Methods(http.MethodGet)
	r.HandleFunc("/v1/company/{companyId}/url", api.CreateCompanyUrl).Methods(http.MethodPost)
	r.HandleFunc("/v1/company/{companyId}/url/{id}", api.UpdateCompanyUrl).Methods(http.MethodPut)
	r.HandleFunc("/v1/company/{companyId}/url/{id}", api.DeleteCompanyUrl).Methods(http.MethodDelete)
}
