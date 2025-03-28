package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
	"github.com/gorilla/mux"
)

type AttributeValueV1 struct {
	Id           string                         `json:"id"`
	Translations []*AttributeValueTranslationV1 `json:"translations"`
	Value        string                         `json:"value"`
	IsEnabled    bool                           `json:"is_enabled"`
}

type AttributeValueTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

type AttributeValueListV1 struct {
	rest.PaginatedList
	Items []*AttributeValueListItemV1 `json:"items"`
}

type AttributeValueListItemV1 struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Value       string `json:"value"`
	IsEnabled   bool   `json:"is_enabled"`
}

type CreateAttributeValueV1 struct {
	Translations []*AttributeValueTranslationV1 `json:"translations"`
	Value        string                         `json:"value"`
	IsEnabled    bool                           `json:"is_enabled"`
}

type UpdateAttributeValueV1 struct {
	Translations []*AttributeValueTranslationV1 `json:"translations"`
	Value        string                         `json:"value"`
	IsEnabled    bool                           `json:"is_enabled"`
}

func (api *apiV1) GetAttributeValuesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	attributeId := mux.Vars(r)["attributeId"]
	filter := api.parseAttributeValueFilterV1(r)
	paging := rest.GetPaging(r)
	sort := rest.GetSorting(r)

	result, count, err := api.service.GetAttributeValues(ctx, attributeId, (paging.PageIndex-1)*paging.PageSize, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := AttributeValueListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*AttributeValueListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, AttributeValueToListItemViewModelV1(item, filter.Language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetAttributeValueByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	attributeId := mux.Vars(r)["attributeId"]
	id := mux.Vars(r)["id"]
	result, err := api.service.GetAttributeValueById(ctx, attributeId, id)
	if api.handleError(w, err) {
		return
	}

	response := AttributeValueToViewModelV1(result)
	rest.WriteResult(w, response)

}

func (api *apiV1) CreateAttributeValueHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	attributeId := mux.Vars(r)["attributeId"]

	var model *CreateAttributeValueV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateAttributeValue(ctx, attributeId, AttributeValueFromCreateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := AttributeValueToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateAttributeValueHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	attributeId := mux.Vars(r)["attributeId"]
	id := mux.Vars(r)["id"]

	var model *UpdateAttributeValueV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateAttributeValue(ctx, attributeId, id, AttributeValueFromUpdateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := AttributeValueToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteAttributeValueHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	attributeId := mux.Vars(r)["attributeId"]
	id := mux.Vars(r)["id"]

	err := api.service.DeleteAttributeValue(ctx, attributeId, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) parseAttributeValueFilterV1(r *http.Request) *model.AttributeValueFilter {
	filter := &model.AttributeValueFilter{}

	filter.Language = r.URL.Query().Get("language")
	if filter.Language == "" {
		filter.Language = api.service.LanguageProvider().UserLanguage(r.Context())
	}
	filter.Name = r.URL.Query().Get("name")

	return filter
}

func AttributeValueToViewModelV1(model *model.AttributeValue) *AttributeValueV1 {
	viewModel := &AttributeValueV1{
		Id:           model.Id,
		Translations: make([]*AttributeValueTranslationV1, 0),
		Value:        model.Value,
		IsEnabled:    model.IsEnabled,
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, AttributeValueTranslationToViewModelV1(translation))
	}
	return viewModel
}

func AttributeValueToListItemViewModelV1(model *model.AttributeValue, language string, defaultLanguage string) *AttributeValueListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &AttributeValueListItemV1{
		Id:          model.Id,
		Name:        translation.Name,
		Slug:        translation.Slug,
		Description: translation.Description,
		Value:       model.Value,
		IsEnabled:   model.IsEnabled,
	}
}

func AttributeValueFromCreateViewModelV1(viewModel *CreateAttributeValueV1) *model.AttributeValue {
	model := &model.AttributeValue{
		Translations: make([]*model.AttributeValueTranslation, 0),
		Value:        viewModel.Value,
		IsEnabled:    viewModel.IsEnabled,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, AttributeValueTranslationFromViewModelV1(translation))
	}
	return model
}

func AttributeValueFromUpdateViewModelV1(viewModel *UpdateAttributeValueV1) *model.AttributeValue {
	model := &model.AttributeValue{
		Translations: make([]*model.AttributeValueTranslation, 0),
		Value:        viewModel.Value,
		IsEnabled:    viewModel.IsEnabled,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, AttributeValueTranslationFromViewModelV1(translation))
	}
	return model
}

func AttributeValueTranslationToViewModelV1(model *model.AttributeValueTranslation) *AttributeValueTranslationV1 {
	return &AttributeValueTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Slug:        model.Slug,
		Description: model.Description,
	}
}

func AttributeValueTranslationFromViewModelV1(viewModel *AttributeValueTranslationV1) *model.AttributeValueTranslation {
	return &model.AttributeValueTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Slug:        viewModel.Slug,
		Description: viewModel.Description,
	}
}
