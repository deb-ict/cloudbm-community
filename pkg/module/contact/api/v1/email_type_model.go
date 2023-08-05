package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type EmailTypeV1 struct {
	Id           string                    `json:"id"`
	Key          string                    `json:"key"`
	Translations []*EmailTypeTranslationV1 `json:"translations"`
	IsDefault    bool                      `json:"is_default"`
	IsSystem     bool                      `json:"is_system"`
}

type EmailTypeTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EmailTypeListV1 struct {
	rest.PaginatedList
	Items []*EmailTypeListItemV1 `json:"items"`
}

type EmailTypeListItemV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsDefault   bool   `json:"is_default"`
	IsSystem    bool   `json:"is_system"`
}

type CreateEmailTypeV1 struct {
	Key          string                    `json:"key"`
	Translations []*EmailTypeTranslationV1 `json:"translations"`
	IsDefault    bool                      `json:"is_default"`
}

type UpdateEmailTypeV1 struct {
	Translations []*EmailTypeTranslationV1 `json:"translations"`
	IsDefault    bool                      `json:"is_default"`
}

func EmailTypeToViewModel(model *model.EmailType) *EmailTypeV1 {
	viewModel := &EmailTypeV1{
		Id:           model.Id,
		Key:          model.Key,
		Translations: make([]*EmailTypeTranslationV1, 0),
		IsDefault:    model.IsDefault,
		IsSystem:     model.IsSystem,
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, EmailTypeTranslationToViewModel(translation))
	}
	return viewModel
}

func EmailTypeToListItemViewModel(model *model.EmailType, language string, defaultLanguage string) *EmailTypeListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &EmailTypeListItemV1{
		Id:          model.Id,
		Key:         model.Key,
		Name:        translation.Name,
		Description: translation.Description,
		IsDefault:   model.IsDefault,
		IsSystem:    model.IsSystem,
	}
}

func EmailTypeFromCreateViewModel(viewModel *CreateEmailTypeV1) *model.EmailType {
	model := &model.EmailType{
		Key:          viewModel.Key,
		Translations: make([]*model.EmailTypeTranslation, 0),
		IsDefault:    viewModel.IsDefault,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, EmailTypeTranslationFromViewModel(translation))
	}
	return model
}

func EmailTypeFromUpdateViewModel(viewModel *UpdateEmailTypeV1) *model.EmailType {
	model := &model.EmailType{
		Translations: make([]*model.EmailTypeTranslation, 0),
		IsDefault:    viewModel.IsDefault,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, EmailTypeTranslationFromViewModel(translation))
	}
	return model
}

func EmailTypeTranslationToViewModel(model *model.EmailTypeTranslation) *EmailTypeTranslationV1 {
	return &EmailTypeTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Description: model.Description,
	}
}

func EmailTypeTranslationFromViewModel(viewModel *EmailTypeTranslationV1) *model.EmailTypeTranslation {
	return &model.EmailTypeTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}
