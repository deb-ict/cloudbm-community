package v1

import (
	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
)

type JobTitleV1 struct {
	Id           string                   `json:"id"`
	Key          string                   `json:"key"`
	Translations []*JobTitleTranslationV1 `json:"translations"`
	IsSystem     bool                     `json:"is_system"`
}

type JobTitleTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type JobTitleListV1 struct {
	rest.PaginatedList
	Items []*JobTitleListItemV1 `json:"items"`
}

type JobTitleListItemV1 struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsSystem    bool   `json:"is_system"`
}

type CreateJobTitleV1 struct {
	Key          string                   `json:"key"`
	Translations []*JobTitleTranslationV1 `json:"translations"`
}

type UpdateJobTitleV1 struct {
	Translations []*JobTitleTranslationV1 `json:"translations"`
}

func JobTitleToViewModel(model *model.JobTitle) *JobTitleV1 {
	viewModel := &JobTitleV1{
		Id:           model.Id,
		Key:          model.Key,
		Translations: make([]*JobTitleTranslationV1, 0),
		IsSystem:     model.IsSystem,
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, JobTitleTranslationToViewModel(translation))
	}
	return viewModel
}

func JobTitleToListItemViewModel(model *model.JobTitle, language string, defaultLanguage string) *JobTitleListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &JobTitleListItemV1{
		Id:          model.Id,
		Key:         model.Key,
		Name:        translation.Name,
		Description: translation.Description,
		IsSystem:    model.IsSystem,
	}
}

func JobTitleFromCreateViewModel(viewModel *CreateJobTitleV1) *model.JobTitle {
	model := &model.JobTitle{
		Key:          viewModel.Key,
		Translations: make([]*model.JobTitleTranslation, 0),
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, JobTitleTranslationFromViewModel(translation))
	}
	return model
}

func JobTitleFromUpdateViewModel(viewModel *UpdateJobTitleV1) *model.JobTitle {
	model := &model.JobTitle{
		Translations: make([]*model.JobTitleTranslation, 0),
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, JobTitleTranslationFromViewModel(translation))
	}
	return model
}

func JobTitleTranslationToViewModel(model *model.JobTitleTranslation) *JobTitleTranslationV1 {
	return &JobTitleTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Description: model.Description,
	}
}

func JobTitleTranslationFromViewModel(viewModel *JobTitleTranslationV1) *model.JobTitleTranslation {
	return &model.JobTitleTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}
