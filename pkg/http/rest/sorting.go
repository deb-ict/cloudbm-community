package rest

import (
	"net/http"
	"strings"

	"github.com/deb-ict/cloudbm-community/pkg/core"
)

func GetSorting(r *http.Request) *core.Sort {
	result := &core.Sort{}
	sortBy := r.URL.Query().Get("sortBy")
	sortOrder := r.URL.Query().Get("sortOrder")
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
