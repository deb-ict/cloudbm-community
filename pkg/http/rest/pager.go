package rest

import (
	"net/http"
	"strconv"
)

const (
	DefaultPageIndex int = 1
	DefaultPageSize  int = 25
	MinPageIndex     int = 1
	MinPageSize      int = 1
	MaxPageSize      int = 150
)

type PageRequest struct {
	PageIndex int
	PageSize  int
}

type PaginatedList struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
	ItemCount int `json:"item_count"`
}

func GetPaging(r *http.Request) PageRequest {
	var err error

	pageIndex, err := strconv.Atoi(r.URL.Query().Get("pageindex"))
	if err != nil || pageIndex < MinPageIndex {
		pageIndex = DefaultPageIndex
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("pagesize"))
	if err != nil || pageSize < MinPageSize {
		pageSize = DefaultPageSize
	}
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}

	return PageRequest{
		PageIndex: pageIndex,
		PageSize:  pageSize,
	}
}
