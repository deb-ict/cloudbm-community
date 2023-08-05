package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type UriTypeV1 struct {
	Id           string                  `json:"id"`
	Key          string                  `json:"key"`
	Translations []*UriTypeTranslationV1 `json:"translations"`
	IsDefault    bool                    `json:"is_default"`
	IsSystem     bool                    `json:"is_system"`
}

type UriTypeTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UriTypeListV1 struct {
	rest.PaginatedList
	Items []*UriTypeListItemV1 `json:"items"`
}

type UriTypeListItemV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsDefault   bool   `json:"is_default"`
	IsSystem    bool   `json:"is_system"`
}

type CreateUriTypeV1 struct {
	Key          string                  `json:"key"`
	Translations []*UriTypeTranslationV1 `json:"translations"`
	IsDefault    bool                    `json:"is_default"`
}

type UpdateUriTypeV1 struct {
	Translations []*UriTypeTranslationV1 `json:"translations"`
	IsDefault    bool                    `json:"is_default"`
}

func UriTypeToViewModel(model *model.UriType) *UriTypeV1 {
	viewModel := &UriTypeV1{
		Id:           model.Id,
		Key:          model.Key,
		Translations: make([]*UriTypeTranslationV1, 0),
		IsDefault:    model.IsDefault,
		IsSystem:     model.IsSystem,
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, UriTypeTranslationToViewModel(translation))
	}
	return viewModel
}

func UriTypeToListItemViewModel(model *model.UriType, language string, defaultLanguage string) *UriTypeListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &UriTypeListItemV1{
		Id:          model.Id,
		Key:         model.Key,
		Name:        translation.Name,
		Description: translation.Description,
		IsDefault:   model.IsDefault,
		IsSystem:    model.IsSystem,
	}
}

func UriTypeFromCreateViewModel(viewModel *CreateUriTypeV1) *model.UriType {
	model := &model.UriType{
		Key:          viewModel.Key,
		Translations: make([]*model.UriTypeTranslation, 0),
		IsDefault:    viewModel.IsDefault,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, UriTypeTranslationFromViewModel(translation))
	}
	return model
}

func UriTypeFromUpdateViewModel(viewModel *UpdateUriTypeV1) *model.UriType {
	model := &model.UriType{
		Translations: make([]*model.UriTypeTranslation, 0),
		IsDefault:    viewModel.IsDefault,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, UriTypeTranslationFromViewModel(translation))
	}
	return model
}

func UriTypeTranslationToViewModel(model *model.UriTypeTranslation) *UriTypeTranslationV1 {
	return &UriTypeTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Description: model.Description,
	}
}

func UriTypeTranslationFromViewModel(viewModel *UriTypeTranslationV1) *model.UriTypeTranslation {
	return &model.UriTypeTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}
