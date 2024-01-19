package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/contact/model"
	"github.com/gorilla/mux"
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

func (api *apiV1) GetJobTitlesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filter := api.parseJobTitleFilterV1(r)
	paging := rest.GetPaging(r)
	sort := rest.GetSorting(r)

	language := r.URL.Query().Get("language")
	if language == "" {
		language = api.service.LanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetJobTitles(ctx, (paging.PageIndex-1)*paging.PageSize, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := JobTitleListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*JobTitleListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, JobTitleToListItemViewModelV1(item, language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetJobTitleByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]
	result, err := api.service.GetJobTitleById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	response := JobTitleToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateJobTitleHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateJobTitleV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateJobTitle(ctx, JobTitleFromCreateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := JobTitleToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateJobTitleHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	var model *UpdateJobTitleV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateJobTitle(ctx, id, JobTitleFromUpdateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := JobTitleToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteJobTitleHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	err := api.service.DeleteJobTitle(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) parseJobTitleFilterV1(r *http.Request) *model.JobTitleFilter {
	return &model.JobTitleFilter{
		Language: localization.GetHttpRequestLanguage(r, api.service.LanguageProvider()),
		Name:     r.URL.Query().Get("name"),
	}
}

func JobTitleToViewModelV1(model *model.JobTitle) *JobTitleV1 {
	viewModel := &JobTitleV1{
		Id:           model.Id,
		Key:          model.Key,
		Translations: make([]*JobTitleTranslationV1, 0),
		IsSystem:     model.IsSystem,
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, JobTitleTranslationToViewModelV1(translation))
	}
	return viewModel
}

func JobTitleToListItemViewModelV1(model *model.JobTitle, language string, defaultLanguage string) *JobTitleListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &JobTitleListItemV1{
		Id:          model.Id,
		Key:         model.Key,
		Name:        translation.Name,
		Description: translation.Description,
		IsSystem:    model.IsSystem,
	}
}

func JobTitleFromCreateViewModelV1(viewModel *CreateJobTitleV1) *model.JobTitle {
	model := &model.JobTitle{
		Key:          viewModel.Key,
		Translations: make([]*model.JobTitleTranslation, 0),
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, JobTitleTranslationFromViewModelV1(translation))
	}
	return model
}

func JobTitleFromUpdateViewModelV1(viewModel *UpdateJobTitleV1) *model.JobTitle {
	model := &model.JobTitle{
		Translations: make([]*model.JobTitleTranslation, 0),
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, JobTitleTranslationFromViewModelV1(translation))
	}
	return model
}

func JobTitleTranslationToViewModelV1(model *model.JobTitleTranslation) *JobTitleTranslationV1 {
	return &JobTitleTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Description: model.Description,
	}
}

func JobTitleTranslationFromViewModelV1(viewModel *JobTitleTranslationV1) *model.JobTitleTranslation {
	return &model.JobTitleTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Description: viewModel.Description,
	}
}
