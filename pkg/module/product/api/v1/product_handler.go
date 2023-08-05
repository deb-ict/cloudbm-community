package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/product/model"
	"github.com/deb-ict/go-router"
)

func (api *apiV1) GetProductsHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paging := rest.GetPaging(r)
	filter := &model.ProductFilter{}
	sort := rest.GetSorting(r)

	categoryId := router.QueryValue(r, "categoryId")
	if categoryId != "" {
		filter.CategoryId = categoryId
	}

	language := router.QueryValue(r, "language")
	if language == "" {
		language = api.service.GetLanguageProvider().UserLanguage(ctx)
	}

	result, count, err := api.service.GetProducts(ctx, paging.PageIndex-1, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := ProductListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*ProductListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, ProductToListItemViewModel(item, language, api.service.GetLanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetProductByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")
	result, err := api.service.GetProductById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	response := ProductToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateProductHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateProductV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateProduct(ctx, ProductFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := ProductToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateProductHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	var model *UpdateProductV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateProduct(ctx, id, ProductFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := ProductToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteProductHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	err := api.service.DeleteProduct(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}
