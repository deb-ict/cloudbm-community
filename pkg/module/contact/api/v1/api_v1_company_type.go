package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/gorilla/mux"
)

type CompanyTypeV1 struct {
	Id           string                      `json:"id"`
	Key          string                      `json:"key"`
	Translations []*CompanyTypeTranslationV1 `json:"translations"`
	IsSystem     bool                        `json:"is_system"`
}

type CompanyTypeTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CompanyTypeListV1 struct {
	rest.PaginatedList
	Items []*CompanyTypeListItemV1 `json:"items"`
}

type CompanyTypeListItemV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsSystem    bool   `json:"is_system"`
}

type CreateCompanyTypeV1 struct {
	Key          string                      `json:"key"`
	Translations []*CompanyTypeTranslationV1 `json:"translations"`
}

type UpdateCompanyTypeV1 struct {
	Translations []*CompanyTypeTranslationV1 `json:"translations"`
}

func (api *apiV1) GetCompanyTypesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filter := api.parseCompanyTypeFilterV1(r)
	paging := rest.GetPaging(r)
	sort := rest.GetSorting(r)

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetCompanyTypes(ctx, (paging.PageIndex-1)*paging.PageSize, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := CompanyTypeListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*CompanyTypeListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, CompanyTypeToListItemViewModelV1(item, language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetCompanyTypeByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]
	result, err := api.service.GetCompanyTypeById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	response := CompanyTypeToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateCompanyTypeHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateCompanyTypeV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateCompanyType(ctx, CompanyTypeFromCreateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := CompanyTypeToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateCompanyTypeHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	var model *UpdateCompanyTypeV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateCompanyType(ctx, id, CompanyTypeFromUpdateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := CompanyTypeToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteCompanyTypeHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	err := api.service.DeleteCompanyType(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) parseCompanyTypeFilterV1(r *http.Request) *model.CompanyTypeFilter {
	return &model.CompanyTypeFilter{
		Language: localization.GetHttpRequestLanguage(r, api.service.LanguageProvider()),
		Name:     r.URL.Query().Get("name"),
	}
}

func CompanyTypeToViewModelV1(model *model.CompanyType) *CompanyTypeV1 {
	viewModel := &CompanyTypeV1{
		Id:           model.Id,
		Key:          model.Key,
		Translations: make([]*CompanyTypeTranslationV1, 0),
		IsSystem:     model.IsSystem,
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, CompanyTypeTranslationToViewModelV1(translation))
	}
	return viewModel
}

func CompanyTypeToListItemViewModelV1(model *model.CompanyType, language string, defaultLanguage string) *CompanyTypeListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &CompanyTypeListItemV1{
		Id:          model.Id,
		Key:         model.Key,
		Name:        translation.Name,
		Description: translation.Description,
		IsSystem:    model.IsSystem,
	}
}

func CompanyTypeFromCreateViewModelV1(viewModel *CreateCompanyTypeV1) *model.CompanyType {
	model := &model.CompanyType{
		Key:          viewModel.Key,
		Translations: make([]*model.CompanyTypeTranslation, 0),
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, CompanyTypeTranslationFromViewModelV1(translation))
	}
	return model
}

func CompanyTypeFromUpdateViewModelV1(viewModel *UpdateCompanyTypeV1) *model.CompanyType {
	model := &model.CompanyType{
		Translations: make([]*model.CompanyTypeTranslation, 0),
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, CompanyTypeTranslationFromViewModelV1(translation))
	}
	return model
}

func CompanyTypeTranslationToViewModelV1(model *model.CompanyTypeTranslation) *CompanyTypeTranslationV1 {
	return &CompanyTypeTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Description: model.Description,
	}
}

func CompanyTypeTranslationFromViewModelV1(viewModel *CompanyTypeTranslationV1) *model.CompanyTypeTranslation {
	return &model.CompanyTypeTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}
