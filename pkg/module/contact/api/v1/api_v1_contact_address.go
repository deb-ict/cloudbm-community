package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/gorilla/mux"
)

type ContactAddressV1 struct {
	Id         string               `json:"id"`
	Type       ContactAddressTypeV1 `json:"type"`
	Street     string               `json:"street"`
	StreetNr   string               `json:"street_nr"`
	Unit       string               `json:"unit"`
	PostalCode string               `json:"postal_code"`
	City       string               `json:"city"`
	State      string               `json:"state"`
	Country    string               `json:"country"`
	IsDefault  bool                 `json:"is_default"`
}

type ContactAddressTypeV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ContactAddressListV1 struct {
	rest.PaginatedList
	Items []*ContactAddressListItemV1 `json:"items"`
}

type ContactAddressListItemV1 struct {
	Id         string               `json:"id"`
	Type       ContactAddressTypeV1 `json:"type"`
	Street     string               `json:"street"`
	StreetNr   string               `json:"street_nr"`
	Unit       string               `json:"unit"`
	PostalCode string               `json:"postal_code"`
	City       string               `json:"city"`
	State      string               `json:"state"`
	Country    string               `json:"country"`
	IsDefault  bool                 `json:"is_default"`
}

type CreateContactAddressV1 struct {
	TypeId     string `json:"type_id"`
	Street     string `json:"street"`
	StreetNr   string `json:"street_nr"`
	Unit       string `json:"unit"`
	PostalCode string `json:"postal_code"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	IsDefault  bool   `json:"is_default"`
}

type UpdateContactAddressV1 struct {
	TypeId     string `json:"type_id"`
	Street     string `json:"street"`
	StreetNr   string `json:"street_nr"`
	Unit       string `json:"unit"`
	PostalCode string `json:"postal_code"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	IsDefault  bool   `json:"is_default"`
}

func (api *apiV1) GetContactAddressesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := mux.Vars(r)["contactId"]

	filter := api.parseContactAddressFilterV1(r)
	paging := rest.GetPaging(r)
	sort := rest.GetSorting(r)

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetContactAddresses(ctx, contactId, (paging.PageIndex-1)*paging.PageSize, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := ContactAddressListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*ContactAddressListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, ContactAddressToListItemViewModelV1(item, language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetContactAddressByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := mux.Vars(r)["contactId"]

	id := mux.Vars(r)["id"]
	result, err := api.service.GetContactAddressById(ctx, contactId, id)
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := ContactAddressToViewModelV1(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateContactAddressHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := mux.Vars(r)["contactId"]

	var model *CreateContactAddressV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateContactAddress(ctx, contactId, ContactAddressFromCreateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := ContactAddressToViewModelV1(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateContactAddressHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := mux.Vars(r)["contactId"]

	id := mux.Vars(r)["id"]

	var model *UpdateContactAddressV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateContactAddress(ctx, contactId, id, ContactAddressFromUpdateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := ContactAddressToViewModelV1(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteContactAddressHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactId := mux.Vars(r)["contactId"]

	id := mux.Vars(r)["id"]

	err := api.service.DeleteContactAddress(ctx, contactId, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) parseContactAddressFilterV1(r *http.Request) *model.AddressFilter {
	return &model.AddressFilter{
		TypeId: r.URL.Query().Get("type"),
	}
}

func ContactAddressToViewModelV1(model *model.Address, language string, defaultLanguage string) *ContactAddressV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &ContactAddressV1{
		Id: model.Id,
		Type: ContactAddressTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		Street:     model.Street,
		StreetNr:   model.StreetNr,
		Unit:       model.Unit,
		PostalCode: model.PostalCode,
		City:       model.City,
		State:      model.State,
		Country:    model.Country,
		IsDefault:  model.IsDefault,
	}
}

func ContactAddressToListItemViewModelV1(model *model.Address, language string, defaultLanguage string) *ContactAddressListItemV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &ContactAddressListItemV1{
		Id: model.Id,
		Type: ContactAddressTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		Street:     model.Street,
		StreetNr:   model.StreetNr,
		Unit:       model.Unit,
		PostalCode: model.PostalCode,
		City:       model.City,
		State:      model.State,
		Country:    model.Country,
		IsDefault:  model.IsDefault,
	}
}

func ContactAddressFromCreateViewModelV1(viewModel *CreateContactAddressV1) *model.Address {
	return &model.Address{
		Type: &model.AddressType{
			Id: viewModel.TypeId,
		},
		Street:     viewModel.Street,
		StreetNr:   viewModel.StreetNr,
		Unit:       viewModel.Unit,
		PostalCode: viewModel.PostalCode,
		City:       viewModel.City,
		State:      viewModel.State,
		Country:    viewModel.Country,
		IsDefault:  viewModel.IsDefault,
	}
}

func ContactAddressFromUpdateViewModelV1(viewModel *UpdateContactAddressV1) *model.Address {
	return &model.Address{
		Type: &model.AddressType{
			Id: viewModel.TypeId,
		},
		Street:     viewModel.Street,
		StreetNr:   viewModel.StreetNr,
		Unit:       viewModel.Unit,
		PostalCode: viewModel.PostalCode,
		City:       viewModel.City,
		State:      viewModel.State,
		Country:    viewModel.Country,
		IsDefault:  viewModel.IsDefault,
	}
}
