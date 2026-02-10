package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/deb-ict/go-router"
)

type CompanyPhoneV1 struct {
	Id          string             `json:"id"`
	Type        CompanyPhoneTypeV1 `json:"type"`
	PhoneNumber string             `json:"number"`
	Extension   string             `json:"extension"`
	IsDefault   bool               `json:"is_default"`
}

type CompanyPhoneTypeV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CompanyPhoneListV1 struct {
	rest.PaginatedList
	Items []*CompanyPhoneListItemV1 `json:"items"`
}

type CompanyPhoneListItemV1 struct {
	Id          string             `json:"id"`
	Type        CompanyPhoneTypeV1 `json:"type"`
	PhoneNumber string             `json:"number"`
	Extension   string             `json:"extension"`
	IsDefault   bool               `json:"is_default"`
}

type CreateCompanyPhoneV1 struct {
	TypeId      string `json:"type_id"`
	PhoneNumber string `json:"number"`
	Extension   string `json:"extension"`
	IsDefault   bool   `json:"is_default"`
}

type UpdateCompanyPhoneV1 struct {
	TypeId      string `json:"type_id"`
	PhoneNumber string `json:"number"`
	Extension   string `json:"extension"`
	IsDefault   bool   `json:"is_default"`
}

func (api *apiV1) GetCompanyPhonesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	filter := api.parseCompanyPhoneFilterV1(r)
	paging := rest.GetPaging(r)
	sort := rest.GetSorting(r)

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetCompanyPhones(ctx, companyId, (paging.PageIndex-1)*paging.PageSize, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := CompanyPhoneListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*CompanyPhoneListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, CompanyPhoneToListItemViewModelV1(item, language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetCompanyPhoneByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	id := router.Param(r, "id")
	result, err := api.service.GetCompanyPhoneById(ctx, companyId, id)
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := CompanyPhoneToViewModelV1(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateCompanyPhoneHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	var model *CreateCompanyPhoneV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateCompanyPhone(ctx, companyId, CompanyPhoneFromCreateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := CompanyPhoneToViewModelV1(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateCompanyPhoneHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	id := router.Param(r, "id")

	var model *UpdateCompanyPhoneV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateCompanyPhone(ctx, companyId, id, CompanyPhoneFromUpdateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	response := CompanyPhoneToViewModelV1(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteCompanyPhoneHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	companyId := router.Param(r, "companyId")

	id := router.Param(r, "id")

	err := api.service.DeleteCompanyPhone(ctx, companyId, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) parseCompanyPhoneFilterV1(r *http.Request) *model.PhoneFilter {
	return &model.PhoneFilter{
		TypeId: r.URL.Query().Get("type"),
	}
}

func CompanyPhoneToViewModelV1(model *model.Phone, language string, defaultLanguage string) *CompanyPhoneV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &CompanyPhoneV1{
		Id: model.Id,
		Type: CompanyPhoneTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		PhoneNumber: model.PhoneNumber,
		Extension:   model.Extension,
		IsDefault:   model.IsDefault,
	}
}

func CompanyPhoneToListItemViewModelV1(model *model.Phone, language string, defaultLanguage string) *CompanyPhoneListItemV1 {
	typeTranslation := model.Type.GetTranslation(language, defaultLanguage)
	return &CompanyPhoneListItemV1{
		Id: model.Id,
		Type: CompanyPhoneTypeV1{
			Id:          model.Type.Id,
			Key:         model.Type.Key,
			Name:        typeTranslation.Name,
			Description: typeTranslation.Description,
		},
		PhoneNumber: model.PhoneNumber,
		Extension:   model.Extension,
		IsDefault:   model.IsDefault,
	}
}

func CompanyPhoneFromCreateViewModelV1(viewModel *CreateCompanyPhoneV1) *model.Phone {
	return &model.Phone{
		Type: &model.PhoneType{
			Id: viewModel.TypeId,
		},
		PhoneNumber: viewModel.PhoneNumber,
		Extension:   viewModel.Extension,
		IsDefault:   viewModel.IsDefault,
	}
}

func CompanyPhoneFromUpdateViewModelV1(viewModel *UpdateCompanyPhoneV1) *model.Phone {
	return &model.Phone{
		Type: &model.PhoneType{
			Id: viewModel.TypeId,
		},
		PhoneNumber: viewModel.PhoneNumber,
		Extension:   viewModel.Extension,
		IsDefault:   viewModel.IsDefault,
	}
}
