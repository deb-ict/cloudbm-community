package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/gorilla/mux"
)

type CompanyUriV1 struct {
	Id        string           `json:"id"`
	Type      CompanyUriTypeV1 `json:"type"`
	Uri       string           `json:"uri"`
	IsDefault bool             `json:"is_default"`
}

type CompanyUriTypeV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CompanyUriListV1 struct {
	rest.PaginatedList
	Items []*CompanyUriListItemV1 `json:"items"`
}

type CompanyUriListItemV1 struct {
	Id        string           `json:"id"`
	Type      CompanyUriTypeV1 `json:"type"`
	Uri       string           `json:"uri"`
	IsDefault bool             `json:"is_default"`
}

type CreateCompanyUriV1 struct {
	TypeId    string `json:"type_id"`
	Uri       string `json:"uri"`
	IsDefault bool   `json:"is_default"`
}

type UpdateCompanyUriV1 struct {
	TypeId    string `json:"type_id"`
	Uri       string `json:"uri"`
	IsDefault bool   `json:"is_default"`
}

func (api *apiV1) GetCompanyUrisHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := mux.Vars(r)["companyId"]

	paging := rest.GetPaging(r)
	filter := &model.UriFilter{}
	sort := rest.GetSorting(r)

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetCompanyUris(ctx, companyId, paging.PageIndex-1, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := CompanyUriListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*CompanyUriListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, CompanyUriToListItemViewModel(item, language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetCompanyUriByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := mux.Vars(r)["companyId"]

	id := mux.Vars(r)["id"]
	result, err := api.service.GetCompanyUriById(ctx, companyId, id)
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := CompanyUriToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateCompanyUriHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := mux.Vars(r)["companyId"]

	var model *CreateCompanyUriV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateCompanyUri(ctx, companyId, CompanyUriFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := CompanyUriToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateCompanyUriHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := mux.Vars(r)["companyId"]

	id := mux.Vars(r)["id"]

	var model *UpdateCompanyUriV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateCompanyUri(ctx, companyId, id, CompanyUriFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := CompanyUriToViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteCompanyUriHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := mux.Vars(r)["companyId"]

	id := mux.Vars(r)["id"]

	err := api.service.DeleteCompanyUri(ctx, companyId, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func CompanyUriToViewModel(model *model.Uri, language string, defaultLanguage string) *CompanyUriV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &CompanyUriV1{
		Id: model.Id,
		Type: CompanyUriTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		Uri:       model.Uri,
		IsDefault: model.IsDefault,
	}
}

func CompanyUriToListItemViewModel(model *model.Uri, language string, defaultLanguage string) *CompanyUriListItemV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &CompanyUriListItemV1{
		Id: model.Id,
		Type: CompanyUriTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		Uri:       model.Uri,
		IsDefault: model.IsDefault,
	}
}

func CompanyUriFromCreateViewModel(viewModel *CreateCompanyUriV1) *model.Uri {
	return &model.Uri{
		Type: model.UriType{
			Id: viewModel.TypeId,
		},
		Uri:       viewModel.Uri,
		IsDefault: viewModel.IsDefault,
	}
}

func CompanyUriFromUpdateViewModel(viewModel *UpdateCompanyUriV1) *model.Uri {
	return &model.Uri{
		Type: model.UriType{
			Id: viewModel.TypeId,
		},
		Uri:       viewModel.Uri,
		IsDefault: viewModel.IsDefault,
	}
}
