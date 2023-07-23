package rest

import (
	"net/http"
	"strings"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/go-router"
)

func GetSorting(r *http.Request) *core.Sort {
	result := &core.Sort{}
	sortBy := router.QueryValue(r, "sortBy")
	sortOrder := router.QueryValue(r, "sortOrder")
	if sortBy != "" {
		sortField := core.SortField{Name: sortBy}
		if strings.ToLower(sortOrder) != "desc" {
			sortField.Order = core.SortAscending
		} else {
			sortField.Order = core.SortDescending
		}

		result.Fields = append(result.Fields, sortField)
	}
	return result
}
