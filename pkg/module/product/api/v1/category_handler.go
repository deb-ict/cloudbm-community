package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
	"github.com/deb-ict/go-router"
)

func (api *apiV1) GetCateogiesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paging := rest.GetPaging(r)
	filter := &model.CategoryFilter{}
	sort := rest.GetSorting(r)

	parentId := router.QueryValue(r, "parentId")
	if parentId != "" {
		filter.ParentId = parentId
	}

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.GetLanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetCategories(ctx, paging.PageIndex-1, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := CategoryListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*CategoryListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, CategoryToListItemViewModel(item, language, api.service.GetLanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetCategoryByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")
	result, err := api.service.GetCategoryById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	response := CategoryToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateCategoryHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateCategoryV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateCategory(ctx, CategoryFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := CategoryToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateCategoryHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	var model *UpdateCategoryV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateCategory(ctx, id, CategoryFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := CategoryToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteCategoryHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	err := api.service.DeleteCategory(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}
