package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/global/model"
	"github.com/deb-ict/go-router"
	"github.com/shopspring/decimal"
)

type TaxProfileV1 struct {
	Id           string                     `json:"id"`
	Key          string                     `json:"key"`
	Translations []*TaxProfileTranslationV1 `json:"translations"`
	Rate         decimal.Decimal            `json:"rate"`
}

type LocalizedTaxProfileV1 struct {
	Id          string          `json:"id"`
	Key         string          `json:"key"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Rate        decimal.Decimal `json:"rate"`
}

type TaxProfileTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TaxProfileListV1 struct {
	rest.PaginatedList
	Items []*TaxProfileListItemV1 `json:"items"`
}

type TaxProfileListItemV1 struct {
	Id          string          `json:"id"`
	Key         string          `json:"key"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Rate        decimal.Decimal `json:"rate"`
}

type CreateTaxProfileV1 struct {
	Key          string                     `json:"key"`
	Translations []*TaxProfileTranslationV1 `json:"translations"`
	Rate         decimal.Decimal            `json:"rate"`
}

type UpdateTaxProfileV1 struct {
	Key          string                     `json:"key"`
	Translations []*TaxProfileTranslationV1 `json:"translations"`
	Rate         decimal.Decimal            `json:"rate"`
}

func (api *apiV1) GetTaxProfilesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paging := rest.GetPaging(r)
	filter := &model.TaxProfileFilter{}
	sort := rest.GetSorting(r)

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetTaxProfiles(ctx, paging.PageIndex-1, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := TaxProfileListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*TaxProfileListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, TaxProfileToListItemViewModel(item, language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetTaxProfileByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")
	result, err := api.service.GetTaxProfileById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	language := router.QueryValue(r, "language")
	if language == "" {
		response := TaxProfileToViewModel(result)
		rest.WriteResult(w, response)
	} else {
		response := TaxProfileToLocalizedViewModel(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
		rest.WriteResult(w, response)
	}
}

func (api *apiV1) CreateTaxProfileHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateTaxProfileV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateTaxProfile(ctx, TaxProfileFromCreateModel(model))
	if api.handleError(w, err) {
		return
	}

	response := TaxProfileToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateTaxProfileHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	var model *UpdateTaxProfileV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateTaxProfile(ctx, id, TaxProfileFromUpdateModel(model))
	if api.handleError(w, err) {
		return
	}

	response := TaxProfileToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteTaxProfileHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	err := api.service.DeleteTaxProfile(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) RegisterTaxProfileRoutes(r *router.Router) {
	r.HandleFunc(
		"/v1/taxProfile",
		api.GetTaxProfilesHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized("taxProfile.read"),
	)
	r.HandleFunc(
		"/v1/taxProfile/{id}",
		api.GetTaxProfileByIdHandlerV1,
		router.AllowedMethod(http.MethodGet),
		router.Authorized("taxProfile.read"),
	)
	r.HandleFunc(
		"/v1/taxProfile",
		api.CreateTaxProfileHandlerV1,
		router.AllowedMethod(http.MethodPost),
		router.Authorized("taxProfile.create"),
	)
	r.HandleFunc(
		"/v1/taxProfile/{id}",
		api.UpdateTaxProfileHandlerV1,
		router.AllowedMethod(http.MethodPut),
		router.Authorized("taxProfile.update"),
	)
	r.HandleFunc(
		"/v1/taxProfile/{id}",
		api.DeleteTaxProfileHandlerV1,
		router.AllowedMethod(http.MethodDelete),
		router.Authorized("taxProfile.delete"),
	)
}

func TaxProfileToViewModel(model *model.TaxProfile) *TaxProfileV1 {
	viewModel := &TaxProfileV1{
		Id:           model.Id,
		Key:          model.Key,
		Translations: make([]*TaxProfileTranslationV1, 0),
		Rate:         model.Rate,
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, TaxProfileTranslationToViewModel(translation))
	}
	return viewModel
}

func TaxProfileToLocalizedViewModel(model *model.TaxProfile, language string, defaultLanguage string) *LocalizedTaxProfileV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &LocalizedTaxProfileV1{
		Id:          model.Id,
		Key:         model.Key,
		Name:        translation.Name,
		Description: translation.Description,
		Rate:        model.Rate,
	}
}

func TaxProfileToListItemViewModel(model *model.TaxProfile, language string, defaultLanguage string) *TaxProfileListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &TaxProfileListItemV1{
		Id:          model.Id,
		Key:         model.Key,
		Name:        translation.Name,
		Description: translation.Description,
		Rate:        model.Rate,
	}
}

func TaxProfileFromCreateModel(viewModel *CreateTaxProfileV1) *model.TaxProfile {
	model := &model.TaxProfile{
		Key:          viewModel.Key,
		Translations: make([]*model.TaxProfileTranslation, 0),
		Rate:         viewModel.Rate,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, TaxProfileTranslationFromViewModel(translation))
	}
	return model
}

func TaxProfileFromUpdateModel(viewModel *UpdateTaxProfileV1) *model.TaxProfile {
	model := &model.TaxProfile{
		Key:          viewModel.Key,
		Translations: make([]*model.TaxProfileTranslation, 0),
		Rate:         viewModel.Rate,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, TaxProfileTranslationFromViewModel(translation))
	}
	return model
}

func TaxProfileTranslationToViewModel(model *model.TaxProfileTranslation) *TaxProfileTranslationV1 {
	return &TaxProfileTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Description: model.Description,
	}
}

func TaxProfileTranslationFromViewModel(viewModel *TaxProfileTranslationV1) *model.TaxProfileTranslation {
	return &model.TaxProfileTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}
