package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/gorilla/mux"
)

type CompanyEmailV1 struct {
	Id        string             `json:"id"`
	Type      CompanyEmailTypeV1 `json:"type"`
	Email     string             `json:"email"`
	IsDefault bool               `json:"is_default"`
}

type CompanyEmailTypeV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CompanyEmailListV1 struct {
	rest.PaginatedList
	Items []*CompanyEmailListItemV1 `json:"items"`
}

type CompanyEmailListItemV1 struct {
	Id        string             `json:"id"`
	Type      CompanyEmailTypeV1 `json:"type"`
	Email     string             `json:"email"`
	IsDefault bool               `json:"is_default"`
}

type CreateCompanyEmailV1 struct {
	TypeId    string `json:"type_id"`
	Email     string `json:"email"`
	IsDefault bool   `json:"is_default"`
}

type UpdateCompanyEmailV1 struct {
	TypeId    string `json:"type_id"`
	Email     string `json:"email"`
	IsDefault bool   `json:"is_default"`
}

func (api *apiV1) GetCompanyEmailsHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := mux.Vars(r)["companyId"]

	filter := api.parseCompanyEmailFilterV1(r)
	paging := rest.GetPaging(r)
	sort := rest.GetSorting(r)

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetCompanyEmails(ctx, companyId, (paging.PageIndex-1)*paging.PageSize, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := CompanyEmailListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*CompanyEmailListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, CompanyEmailToListItemViewModelV1(item, language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetCompanyEmailByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := mux.Vars(r)["companyId"]

	id := mux.Vars(r)["id"]
	result, err := api.service.GetCompanyEmailById(ctx, companyId, id)
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := CompanyEmailToViewModelV1(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateCompanyEmailHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := mux.Vars(r)["companyId"]

	var model *CreateCompanyEmailV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateCompanyEmail(ctx, companyId, CompanyEmailFromCreateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := CompanyEmailToViewModelV1(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateCompanyEmailHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := mux.Vars(r)["companyId"]

	id := mux.Vars(r)["id"]

	var model *UpdateCompanyEmailV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateCompanyEmail(ctx, companyId, id, CompanyEmailFromUpdateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := CompanyEmailToViewModelV1(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteCompanyEmailHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := mux.Vars(r)["companyId"]

	id := mux.Vars(r)["id"]

	err := api.service.DeleteCompanyEmail(ctx, companyId, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) parseCompanyEmailFilterV1(r *http.Request) *model.EmailFilter {
	return &model.EmailFilter{
		TypeId: r.URL.Query().Get("type"),
	}
}

func CompanyEmailToViewModelV1(model *model.Email, language string, defaultLanguage string) *CompanyEmailV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &CompanyEmailV1{
		Id: model.Id,
		Type: CompanyEmailTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		Email:     model.Email,
		IsDefault: model.IsDefault,
	}
}

func CompanyEmailToListItemViewModelV1(model *model.Email, language string, defaultLanguage string) *CompanyEmailListItemV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &CompanyEmailListItemV1{
		Id: model.Id,
		Type: CompanyEmailTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		Email:     model.Email,
		IsDefault: model.IsDefault,
	}
}

func CompanyEmailFromCreateViewModelV1(viewModel *CreateCompanyEmailV1) *model.Email {
	return &model.Email{
		Type: &model.EmailType{
			Id: viewModel.TypeId,
		},
		Email:     viewModel.Email,
		IsDefault: viewModel.IsDefault,
	}
}

func CompanyEmailFromUpdateViewModelV1(viewModel *UpdateCompanyEmailV1) *model.Email {
	return &model.Email{
		Type: &model.EmailType{
			Id: viewModel.TypeId,
		},
		Email:     viewModel.Email,
		IsDefault: viewModel.IsDefault,
	}
}
