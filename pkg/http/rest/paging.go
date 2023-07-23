package rest

import (
	"net/http"
	"strconv"

	"github.com/deb-ict/go-router"
)

const (
	DefaultPageIndex int64 = 1
	DefaultPageSize  int64 = 25
	MinPageIndex     int64 = 1
	MinPageSize      int64 = 1
	MaxPageSize      int64 = 150
)

type PageRequest struct {
	PageIndex int64
	PageSize  int64
}

type PaginatedList struct {
	PageIndex int64 `json:"pageIndex"`
	PageSize  int64 `json:"pageSize"`
	ItemCount int64 `json:"count"`
}

func GetPaging(r *http.Request) PageRequest {
	var err error

	pageIndex, err := strconv.ParseInt(router.QueryValue(r, "pageIndex"), 10, 64)
	if err != nil || pageIndex < MinPageIndex {
		pageIndex = DefaultPageIndex
	}

	pageSize, err := strconv.ParseInt(router.QueryValue(r, "pageSize"), 10, 64)
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
