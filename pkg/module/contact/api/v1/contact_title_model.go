package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type ContactTitleV1 struct {
	Id           string                       `json:"id"`
	Key          string                       `json:"key"`
	Translations []*ContactTitleTranslationV1 `json:"translations"`
	IsSystem     bool                         `json:"is_system"`
}

type ContactTitleTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ContactTitleListV1 struct {
	rest.PaginatedList
	Items []*ContactTitleListItemV1 `json:"items"`
}

type ContactTitleListItemV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsSystem    bool   `json:"is_system"`
}

type CreateContactTitleV1 struct {
	Key          string                       `json:"key"`
	Translations []*ContactTitleTranslationV1 `json:"translations"`
}

type UpdateContactTitleV1 struct {
	Translations []*ContactTitleTranslationV1 `json:"translations"`
}

func ContactTitleToViewModel(model *model.ContactTitle) *ContactTitleV1 {
	viewModel := &ContactTitleV1{
		Id:           model.Id,
		Key:          model.Key,
		Translations: make([]*ContactTitleTranslationV1, 0),
		IsSystem:     model.IsSystem,
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, ContactTitleTranslationToViewModel(translation))
	}
	return viewModel
}

func ContactTitleToListItemViewModel(model *model.ContactTitle, language string, defaultLanguage string) *ContactTitleListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &ContactTitleListItemV1{
		Id:          model.Id,
		Key:         model.Key,
		Name:        translation.Name,
		Description: translation.Description,
		IsSystem:    model.IsSystem,
	}
}

func ContactTitleFromCreateViewModel(viewModel *CreateContactTitleV1) *model.ContactTitle {
	model := &model.ContactTitle{
		Key:          viewModel.Key,
		Translations: make([]*model.ContactTitleTranslation, 0),
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, ContactTitleTranslationFromViewModel(translation))
	}
	return model
}

func ContactTitleFromUpdateViewModel(viewModel *UpdateContactTitleV1) *model.ContactTitle {
	model := &model.ContactTitle{
		Translations: make([]*model.ContactTitleTranslation, 0),
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, ContactTitleTranslationFromViewModel(translation))
	}
	return model
}

func ContactTitleTranslationToViewModel(model *model.ContactTitleTranslation) *ContactTitleTranslationV1 {
	return &ContactTitleTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Description: model.Description,
	}
}

func ContactTitleTranslationFromViewModel(viewModel *ContactTitleTranslationV1) *model.ContactTitleTranslation {
	return &model.ContactTitleTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}
