package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/metadata/model"
	"github.com/gorilla/mux"
	"github.com/shopspring/decimal"
)

type TaxRateV1 struct {
	Id           string                  `json:"id"`
	Key          string                  `json:"key"`
	Translations []*TaxRateTranslationV1 `json:"translations"`
	Rate         decimal.Decimal         `json:"rate"`
	IsEnabled    bool                    `json:"is_enabled"`
}

type LocalizedTaxRateV1 struct {
	Id          string          `json:"id"`
	Key         string          `json:"key"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Rate        decimal.Decimal `json:"rate"`
	IsEnabled   bool            `json:"is_enabled"`
}

type TaxRateTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TaxRateListV1 struct {
	rest.PaginatedList
	Items []*TaxRateListItemV1 `json:"items"`
}

type TaxRateListItemV1 struct {
	Id          string          `json:"id"`
	Key         string          `json:"key"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Rate        decimal.Decimal `json:"rate"`
	IsEnabled   bool            `json:"is_enabled"`
}

type CreateTaxRateV1 struct {
	Key          string                  `json:"key"`
	Translations []*TaxRateTranslationV1 `json:"translations"`
	Rate         decimal.Decimal         `json:"rate"`
}

type UpdateTaxRateV1 struct {
	Key          string                  `json:"key"`
	Translations []*TaxRateTranslationV1 `json:"translations"`
	Rate         decimal.Decimal         `json:"rate"`
	IsEnabled    bool                    `json:"is_enabled"`
}

func (api *apiV1) GetTaxRatesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filter := api.parseTaxRateFilterV1(r)
	paging := rest.GetPaging(r)
	sort := rest.GetSorting(r)

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetTaxRates(ctx, (paging.PageIndex-1)*paging.PageSize, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := TaxRateListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*TaxRateListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, TaxRateToListItemViewModelV1(item, language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetTaxRateByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]
	result, err := api.service.GetTaxRateById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	language := r.URL.Query().Get("language")
	if language == "" {
		response := TaxRateToViewModelV1(result)
		rest.WriteResult(w, response)
	} else {
		response := TaxRateToLocalizedViewModelV1(result, language, api.service.LanguageProvider().DefaultLanguage(ctx))
		rest.WriteResult(w, response)
	}
}

func (api *apiV1) CreateTaxRateHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateTaxRateV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateTaxRate(ctx, TaxRateFromCreateModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := TaxRateToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateTaxRateHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	var model *UpdateTaxRateV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateTaxRate(ctx, id, TaxRateFromUpdateModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := TaxRateToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteTaxRateHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	err := api.service.DeleteTaxRate(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) parseTaxRateFilterV1(r *http.Request) *model.TaxRateFilter {
	return &model.TaxRateFilter{}
}

func TaxRateToViewModelV1(model *model.TaxRate) *TaxRateV1 {
	viewModel := &TaxRateV1{
		Id:           model.Id,
		Key:          model.Key,
		Translations: make([]*TaxRateTranslationV1, 0),
		Rate:         model.Rate,
		IsEnabled:    model.IsEnabled,
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, TaxRateTranslationToViewModelV1(translation))
	}
	return viewModel
}

func TaxRateToLocalizedViewModelV1(model *model.TaxRate, language string, defaultLanguage string) *LocalizedTaxRateV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &LocalizedTaxRateV1{
		Id:          model.Id,
		Key:         model.Key,
		Name:        translation.Name,
		Description: translation.Description,
		Rate:        model.Rate,
		IsEnabled:   model.IsEnabled,
	}
}

func TaxRateToListItemViewModelV1(model *model.TaxRate, language string, defaultLanguage string) *TaxRateListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &TaxRateListItemV1{
		Id:          model.Id,
		Key:         model.Key,
		Name:        translation.Name,
		Description: translation.Description,
		Rate:        model.Rate,
		IsEnabled:   model.IsEnabled,
	}
}

func TaxRateFromCreateModelV1(viewModel *CreateTaxRateV1) *model.TaxRate {
	model := &model.TaxRate{
		Key:          viewModel.Key,
		Translations: make([]*model.TaxRateTranslation, 0),
		Rate:         viewModel.Rate,
		IsEnabled:    true,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, TaxRateTranslationFromViewModelV1(translation))
	}
	return model
}

func TaxRateFromUpdateModelV1(viewModel *UpdateTaxRateV1) *model.TaxRate {
	model := &model.TaxRate{
		Key:          viewModel.Key,
		Translations: make([]*model.TaxRateTranslation, 0),
		Rate:         viewModel.Rate,
		IsEnabled:    viewModel.IsEnabled,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, TaxRateTranslationFromViewModelV1(translation))
	}
	return model
}

func TaxRateTranslationToViewModelV1(model *model.TaxRateTranslation) *TaxRateTranslationV1 {
	return &TaxRateTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Description: model.Description,
	}
}

func TaxRateTranslationFromViewModelV1(viewModel *TaxRateTranslationV1) *model.TaxRateTranslation {
	return &model.TaxRateTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}
