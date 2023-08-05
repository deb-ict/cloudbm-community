package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type AddressTypeV1 struct {
	Id           string                      `json:"id"`
	Key          string                      `json:"key"`
	Translations []*AddressTypeTranslationV1 `json:"translations"`
	IsDefault    bool                        `json:"is_default"`
	IsSystem     bool                        `json:"is_system"`
}

type AddressTypeTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AddressTypeListV1 struct {
	rest.PaginatedList
	Items []*AddressTypeListItemV1 `json:"items"`
}

type AddressTypeListItemV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsDefault   bool   `json:"is_default"`
	IsSystem    bool   `json:"is_system"`
}

type CreateAddressTypeV1 struct {
	Key          string                      `json:"key"`
	Translations []*AddressTypeTranslationV1 `json:"translations"`
	IsDefault    bool                        `json:"is_default"`
}

type UpdateAddressTypeV1 struct {
	Translations []*AddressTypeTranslationV1 `json:"translations"`
	IsDefault    bool                        `json:"is_default"`
}

func AddressTypeToViewModel(model *model.AddressType) *AddressTypeV1 {
	viewModel := &AddressTypeV1{
		Id:           model.Id,
		Key:          model.Key,
		Translations: make([]*AddressTypeTranslationV1, 0),
		IsDefault:    model.IsDefault,
		IsSystem:     model.IsSystem,
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, AddressTypeTranslationToViewModel(translation))
	}
	return viewModel
}

func AddressTypeToListItemViewModel(model *model.AddressType, language string, defaultLanguage string) *AddressTypeListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &AddressTypeListItemV1{
		Id:          model.Id,
		Key:         model.Key,
		Name:        translation.Name,
		Description: translation.Description,
		IsDefault:   model.IsDefault,
		IsSystem:    model.IsSystem,
	}
}

func AddressTypeFromCreateViewModel(viewModel *CreateAddressTypeV1) *model.AddressType {
	model := &model.AddressType{
		Key:          viewModel.Key,
		Translations: make([]*model.AddressTypeTranslation, 0),
		IsDefault:    viewModel.IsDefault,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, AddressTypeTranslationFromViewModel(translation))
	}
	return model
}

func AddressTypeFromUpdateViewModel(viewModel *UpdateAddressTypeV1) *model.AddressType {
	model := &model.AddressType{
		Translations: make([]*model.AddressTypeTranslation, 0),
		IsDefault:    viewModel.IsDefault,
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, AddressTypeTranslationFromViewModel(translation))
	}
	return model
}

func AddressTypeTranslationToViewModel(model *model.AddressTypeTranslation) *AddressTypeTranslationV1 {
	return &AddressTypeTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Description: model.Description,
	}
}

func AddressTypeTranslationFromViewModel(viewModel *AddressTypeTranslationV1) *model.AddressTypeTranslation {
	return &model.AddressTypeTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}
