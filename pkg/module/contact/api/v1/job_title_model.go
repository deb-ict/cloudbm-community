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
	return &JobTitleV1{}
}

func JobTitleToListItemViewModel(model *model.JobTitle, language string, defaultLanguage string) *JobTitleListItemV1 {
	return &JobTitleListItemV1{}
}

func JobTitleFromCreateViewModel(viewModel *CreateJobTitleV1) *model.JobTitle {
	return &model.JobTitle{}
}

func JobTitleFromUpdateViewModel(viewModel *UpdateJobTitleV1) *model.JobTitle {
	return &model.JobTitle{}
}

func JobTitleTranslationToViewModel(model *model.JobTitleTranslation) *JobTitleTranslationV1 {
	return &JobTitleTranslationV1{}
}

func JobTitleTranslationFromViewModel(viewModel *JobTitleTranslationV1) *model.JobTitleTranslation {
	return &model.JobTitleTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}
