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
	return &PhoneTypeV1{}
}

func PhoneTypeToListItemViewModel(model *model.PhoneType, language string, defaultLanguage string) *PhoneTypeListItemV1 {
	return &PhoneTypeListItemV1{}
}

func PhoneTypeFromCreateViewModel(viewModel *CreatePhoneTypeV1) *model.PhoneType {
	return &model.PhoneType{}
}

func PhoneTypeFromUpdateViewModel(viewModel *UpdatePhoneTypeV1) *model.PhoneType {
	return &model.PhoneType{}
}

func PhoneTypeTranslationToViewModel(model *model.PhoneTypeTranslation) *PhoneTypeTranslationV1 {
	return &PhoneTypeTranslationV1{}
}

func PhoneTypeTranslationFromViewModel(viewModel *PhoneTypeTranslationV1) *model.PhoneTypeTranslation {
	return &model.PhoneTypeTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}
