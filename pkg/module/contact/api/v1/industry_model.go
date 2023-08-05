package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type IndustryV1 struct {
	Id           string                   `json:"id"`
	Key          string                   `json:"key"`
	Translations []*IndustryTranslationV1 `json:"translations"`
	IsSystem     bool                     `json:"is_system"`
}

type IndustryTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type IndustryListV1 struct {
	rest.PaginatedList
	Items []*IndustryListItemV1 `json:"items"`
}

type IndustryListItemV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsSystem    bool   `json:"is_system"`
}

type CreateIndustryV1 struct {
	Key          string                   `json:"key"`
	Translations []*IndustryTranslationV1 `json:"translations"`
}

type UpdateIndustryV1 struct {
	Translations []*IndustryTranslationV1 `json:"translations"`
}

func IndustryToViewModel(model *model.Industry) *IndustryV1 {
	viewModel := &IndustryV1{
		Id:           model.Id,
		Key:          model.Key,
		Translations: make([]*IndustryTranslationV1, 0),
		IsSystem:     model.IsSystem,
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, IndustryTranslationToViewModel(translation))
	}
	return viewModel
}

func IndustryToListItemViewModel(model *model.Industry, language string, defaultLanguage string) *IndustryListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &IndustryListItemV1{
		Id:          model.Id,
		Key:         model.Key,
		Name:        translation.Name,
		Description: translation.Description,
		IsSystem:    model.IsSystem,
	}
}

func IndustryFromCreateViewModel(viewModel *CreateIndustryV1) *model.Industry {
	model := &model.Industry{
		Key:          viewModel.Key,
		Translations: make([]*model.IndustryTranslation, 0),
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, IndustryTranslationFromViewModel(translation))
	}
	return model
}

func IndustryFromUpdateViewModel(viewModel *UpdateIndustryV1) *model.Industry {
	model := &model.Industry{
		Translations: make([]*model.IndustryTranslation, 0),
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, IndustryTranslationFromViewModel(translation))
	}
	return model
}

func IndustryTranslationToViewModel(model *model.IndustryTranslation) *IndustryTranslationV1 {
	return &IndustryTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Description: model.Description,
	}
}

func IndustryTranslationFromViewModel(viewModel *IndustryTranslationV1) *model.IndustryTranslation {
	return &model.IndustryTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}
