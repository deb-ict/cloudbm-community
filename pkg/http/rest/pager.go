package rest

import (
	"net/http"
	"strconv"
)

type PageRequest struct {
	PageIndex int
	PageSize  int
}

func GetPaging(r *http.Request) PageRequest {
	var err error
	pageIndex, err := strconv.Atoi(r.URL.Query().Get("pageindex"))
	if err != nil {
		pageIndex = 1
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("pagesize"))
	if err != nil {
		pageSize = 25
	}
	return PageRequest{
		PageIndex: pageIndex,
		PageSize:  pageSize,
	}
}
