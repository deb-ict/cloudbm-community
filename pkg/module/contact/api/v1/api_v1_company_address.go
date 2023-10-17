package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/gorilla/mux"
)

type CompanyAddressV1 struct {
	Id         string               `json:"id"`
	Type       CompanyAddressTypeV1 `json:"type"`
	Street     string               `json:"street"`
	StreetNr   string               `json:"street_nr"`
	Unit       string               `json:"unit"`
	PostalCode string               `json:"postal_code"`
	City       string               `json:"city"`
	State      string               `json:"state"`
	Country    string               `json:"country"`
	IsDefault  bool                 `json:"is_default"`
}

type CompanyAddressTypeV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CompanyAddressListV1 struct {
	rest.PaginatedList
	Items []*CompanyAddressListItemV1 `json:"items"`
}

type CompanyAddressListItemV1 struct {
	Id         string               `json:"id"`
	Type       CompanyAddressTypeV1 `json:"type"`
	Street     string               `json:"street"`
	StreetNr   string               `json:"street_nr"`
	Unit       string               `json:"unit"`
	PostalCode string               `json:"postal_code"`
	City       string               `json:"city"`
	State      string               `json:"state"`
	Country    string               `json:"country"`
	IsDefault  bool                 `json:"is_default"`
}

type CreateCompanyAddressV1 struct {
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

type UpdateCompanyAddressV1 struct {
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

func (api *apiV1) GetCompanyAddressesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := mux.Vars(r)["companyId"]

	paging := rest.GetPaging(r)
	filter := &model.AddressFilter{}
	sort := rest.GetSorting(r)

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetCompanyAddresses(ctx, companyId, paging.PageIndex-1, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := CompanyAddressListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*CompanyAddressListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, CompanyAddressToListItemViewModel(item, language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetCompanyAddressByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := mux.Vars(r)["companyId"]

	id := mux.Vars(r)["id"]
	result, err := api.service.GetCompanyAddressById(ctx, companyId, id)
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := CompanyAddressToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateCompanyAddressHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := mux.Vars(r)["companyId"]

	var model *CreateCompanyAddressV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateCompanyAddress(ctx, companyId, CompanyAddressFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := CompanyAddressToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateCompanyAddressHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := mux.Vars(r)["companyId"]

	id := mux.Vars(r)["id"]

	var model *UpdateCompanyAddressV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateCompanyAddress(ctx, companyId, id, CompanyAddressFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := CompanyAddressToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteCompanyAddressHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := mux.Vars(r)["companyId"]

	id := mux.Vars(r)["id"]

	err := api.service.DeleteCompanyAddress(ctx, companyId, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func CompanyAddressToViewModel(model *model.Address, language string, defaultLanguage string) *CompanyAddressV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &CompanyAddressV1{
		Id: model.Id,
		Type: CompanyAddressTypeV1{
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

func CompanyAddressToListItemViewModel(model *model.Address, language string, defaultLanguage string) *CompanyAddressListItemV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &CompanyAddressListItemV1{
		Id: model.Id,
		Type: CompanyAddressTypeV1{
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

func CompanyAddressFromCreateViewModel(viewModel *CreateCompanyAddressV1) *model.Address {
	return &model.Address{
		Type: model.AddressType{
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

func CompanyAddressFromUpdateViewModel(viewModel *UpdateCompanyAddressV1) *model.Address {
	return &model.Address{
		Type: model.AddressType{
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
