package v1

import (
	"encoding/json"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth/model"
	"github.com/gorilla/mux"
)

type UserV1 struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Phone         string `json:"phone"`
	PhoneVerified bool   `json:"phone_verified"`
	IsLocked      bool   `json:"is_locked"`
	IsEnabled     bool   `json:"is_enabled"`
}

type UserListV1 struct {
	rest.PaginatedList
	Items []*UserListItemV1 `json:"items"`
}

type UserListItemV1 struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Phone         string `json:"phone"`
	PhoneVerified bool   `json:"phone_verified"`
	IsLocked      bool   `json:"is_locked"`
	IsEnabled     bool   `json:"is_enabled"`
}

type CreateUserV1 struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Phone         string `json:"phone"`
	PhoneVerified bool   `json:"phone_verified"`
	IsEnabled     bool   `json:"is_enabled"`
}

type UpdateUserV1 struct {
	EmailVerified bool   `json:"email_verified"`
	Phone         string `json:"phone"`
	PhoneVerified bool   `json:"phone_verified"`
	IsEnabled     bool   `json:"is_enabled"`
}

func (api *apiV1) GetUsersHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filter := api.parseUserFilterV1(r)
	paging := rest.GetPaging(r)
	sort := rest.GetSorting(r)

	result, count, err := api.service.GetUsers(ctx, (paging.PageIndex-1)*paging.PageSize, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := UserListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*UserListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, UserToListItemViewModelV1(item))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetUserByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]
	result, err := api.service.GetUserById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	response := UserToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateUserHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateUserV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateUser(ctx, UserFromCreateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := UserToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateUserHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	var model *UpdateUserV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateUser(ctx, id, UserFromUpdateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := UserToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteUserHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	err := api.service.DeleteUser(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) parseUserFilterV1(r *http.Request) *model.UserFilter {
	filter := &model.UserFilter{
		Username: r.URL.Query().Get("username"),
		Email:    r.URL.Query().Get("email"),
	}
	return filter
}

func UserToViewModel(model *model.User) *UserV1 {
	return &UserV1{
		Id:            model.Id,
		Username:      model.Username,
		Email:         model.Email,
		EmailVerified: model.EmailVerified,
		Phone:         model.Phone,
		PhoneVerified: model.PhoneVerified,
		IsLocked:      model.IsLocked,
		IsEnabled:     model.IsEnabled,
	}
}

func UserToListItemViewModelV1(model *model.User) *UserListItemV1 {
	return &UserListItemV1{
		Id:            model.Id,
		Username:      model.Username,
		Email:         model.Email,
		EmailVerified: model.EmailVerified,
		Phone:         model.Phone,
		PhoneVerified: model.PhoneVerified,
		IsLocked:      model.IsLocked,
		IsEnabled:     model.IsEnabled,
	}
}

func UserFromCreateViewModelV1(viewModel *CreateUserV1) *model.User {
	return &model.User{
		Username:      viewModel.Username,
		Email:         viewModel.Email,
		EmailVerified: viewModel.EmailVerified,
		Phone:         viewModel.Phone,
		PhoneVerified: viewModel.PhoneVerified,
		IsEnabled:     viewModel.IsEnabled,
	}
}

func UserFromUpdateViewModelV1(viewModel *UpdateUserV1) *model.User {
	return &model.User{
		EmailVerified: viewModel.EmailVerified,
		Phone:         viewModel.Phone,
		PhoneVerified: viewModel.PhoneVerified,
		IsEnabled:     viewModel.IsEnabled,
	}
}
