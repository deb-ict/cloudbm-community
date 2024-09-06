package v1

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/session/model"
	"github.com/gorilla/mux"
)

type SessionV1 struct {
	Id                   string            `json:"id"`
	UserId               string            `json:"userId"`
	CreatedAt            time.Time         `json:"createdAt"`
	UpdatedAt            time.Time         `json:"updatedAt"`
	ExpiresAt            time.Time         `json:"expiresAt"`
	Lifetime             time.Duration     `json:"lifetime"`
	UseSlidingExpiration bool              `json:"useSlidingExpiration"`
	Data                 map[string]string `json:"data"`
}

type SessionListV1 struct {
	rest.PaginatedList
	Items []*SessionListItemV1 `json:"items"`
}

type SessionListItemV1 struct {
	Id                   string        `json:"id"`
	UserId               string        `json:"userId"`
	CreatedAt            time.Time     `json:"createdAt"`
	UpdatedAt            time.Time     `json:"updatedAt"`
	ExpiresAt            time.Time     `json:"expiresAt"`
	Lifetime             time.Duration `json:"lifetime"`
	UseSlidingExpiration bool          `json:"useSlidingExpiration"`
}

type CreateSessionV1 struct {
	UserId               string            `json:"userId"`
	Lifetime             time.Duration     `json:"lifetime"`
	UseSlidingExpiration bool              `json:"useSlidingExpiration"`
	Data                 map[string]string `json:"data"`
}

type UpdateSessionV1 struct {
	UserId               string            `json:"userId"`
	Lifetime             time.Duration     `json:"lifetime"`
	UseSlidingExpiration bool              `json:"useSlidingExpiration"`
	Data                 map[string]string `json:"data"`
}

func (api *apiV1) GetSessionsHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filter := api.parseSessionFilterV1(r)
	paging := rest.GetPaging(r)
	sort := rest.GetSorting(r)

	result, count, err := api.service.GetSessions(ctx, (paging.PageIndex-1)*paging.PageSize, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := SessionListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*SessionListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, SessionToListItemViewModelV1(item))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetSessionByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]
	result, err := api.service.GetSessionById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	response := SessionToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateSessionHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateSessionV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateSession(ctx, SessionFromCreateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := SessionToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateSessionHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	var model *UpdateSessionV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateSession(ctx, id, SessionFromUpdateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := SessionToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteSessionHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	err := api.service.DeleteSession(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) CleanupExpiredSessionsHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := api.service.CleanupExpiredSessions(ctx)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) parseSessionFilterV1(r *http.Request) *model.SessionFilter {
	return &model.SessionFilter{}
}

func SessionToViewModelV1(model *model.Session) *SessionV1 {
	viewModel := &SessionV1{
		Id:                   model.Id,
		UserId:               model.UserId,
		CreatedAt:            model.CreatedAt,
		UpdatedAt:            model.UpdatedAt,
		ExpiresAt:            model.ExpiresAt,
		Lifetime:             model.Lifetime,
		UseSlidingExpiration: model.UseSlidingExpiration,
		Data:                 make(map[string]string),
	}
	for key, value := range model.Data {
		viewModel.Data[key] = value
	}
	return viewModel
}

func SessionToListItemViewModelV1(model *model.Session) *SessionListItemV1 {
	return &SessionListItemV1{
		Id:                   model.Id,
		UserId:               model.UserId,
		CreatedAt:            model.CreatedAt,
		UpdatedAt:            model.UpdatedAt,
		ExpiresAt:            model.ExpiresAt,
		Lifetime:             model.Lifetime,
		UseSlidingExpiration: model.UseSlidingExpiration,
	}
}

func SessionFromCreateViewModelV1(viewModel *CreateSessionV1) *model.Session {
	model := &model.Session{
		UserId:               viewModel.UserId,
		Lifetime:             viewModel.Lifetime,
		UseSlidingExpiration: viewModel.UseSlidingExpiration,
		Data:                 make(map[string]string),
	}
	for key, value := range model.Data {
		model.Data[key] = value
	}
	return model
}

func SessionFromUpdateViewModelV1(viewModel *UpdateSessionV1) *model.Session {
	model := &model.Session{
		UserId:               viewModel.UserId,
		Lifetime:             viewModel.Lifetime,
		UseSlidingExpiration: viewModel.UseSlidingExpiration,
		Data:                 make(map[string]string),
	}
	for key, value := range model.Data {
		model.Data[key] = value
	}
	return model
}
