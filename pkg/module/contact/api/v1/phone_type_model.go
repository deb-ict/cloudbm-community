package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type PhoneTypeV1 struct {
	Id           string                    `json:"id"`
	Key          string                    `json:"key"`
	Translations []*PhoneTypeTranslationV1 `json:"translations"`
	IsDefault    bool                      `json:"is_default"`
	IsSystem     bool                      `json:"is_system"`
}

type PhoneTypeTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PhoneTypeListV1 struct {
	rest.PaginatedList
	Items []*PhoneTypeListItemV1 `json:"items"`
}

type PhoneTypeListItemV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsDefault   bool   `json:"is_default"`
	IsSystem    bool   `json:"is_system"`
}

type CreatePhoneTypeV1 struct {
	Key          string                    `json:"key"`
	Translations []*PhoneTypeTranslationV1 `json:"translations"`
	IsDefault    bool                      `json:"is_default"`
}

type UpdatePhoneTypeV1 struct {
	Translations []*PhoneTypeTranslationV1 `json:"translations"`
	IsDefault    bool                      `json:"is_default"`
}

func PhoneTypeToViewModel(model *model.PhoneType) *PhoneTypeV1 {
	viewModel := &PhoneTypeV1{
		Id:           model.Id,
		Key:          model.Key,
		Translations: make([]*PhoneTypeTranslationV1, 0),
		IsDefault:    model.IsDefault,
		IsSystem:     model.IsSystem,
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, PhoneTypeTranslationToViewModel(translation))
	}
	return viewModel
}

func PhoneTypeToListItemViewModel(model *model.PhoneType, language string, defaultLanguage string) *PhoneTypeListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &PhoneTypeListItemV1{
		Id:          model.Id,
		Key:         model.Key,
		Name:        translation.Name,
		Description: translation.Description,
		IsDefault:   model.IsDefault,
		IsSystem:    model.IsSystem,
	}
}

func PhoneTypeFromCreateViewModel(viewModel *CreatePhoneTypeV1) *model.PhoneType {
	model := &model.PhoneType{
		Key:          viewModel.Key,
		Translations: make([]*model.PhoneTypeTranslation, 0),
		IsDefault:    viewModel.IsDefault,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, PhoneTypeTranslationFromViewModel(translation))
	}
	return model
}

func PhoneTypeFromUpdateViewModel(viewModel *UpdatePhoneTypeV1) *model.PhoneType {
	model := &model.PhoneType{
		Translations: make([]*model.PhoneTypeTranslation, 0),
		IsDefault:    viewModel.IsDefault,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, PhoneTypeTranslationFromViewModel(translation))
	}
	return model
}

func PhoneTypeTranslationToViewModel(model *model.PhoneTypeTranslation) *PhoneTypeTranslationV1 {
	return &PhoneTypeTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Description: model.Description,
	}
}

func PhoneTypeTranslationFromViewModel(viewModel *PhoneTypeTranslationV1) *model.PhoneTypeTranslation {
	return &model.PhoneTypeTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}
