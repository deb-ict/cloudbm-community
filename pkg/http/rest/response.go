package rest

import (
	"encoding/json"
	"net/http"
)

const (
	ContentTypeHeaderName  string = "Content-Type"
	ContentTypeHeaderValue string = "application/json; charset=utf-8"
)

type errorInfo struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

func WriteResult(w http.ResponseWriter, data any) {
	WriteJsonResponse(w, http.StatusOK, data)
}

func WriteError(w http.ResponseWriter, statusCode int, message string) {
	WriteJsonResponse(w, statusCode, &errorInfo{
		StatusCode: statusCode,
		Message:    message,
	})
}

func WriteStatus(w http.ResponseWriter, statusCode int) {
	WriteJsonResponse(w, statusCode, nil)
}

func WriteJsonResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set(ContentTypeHeaderName, ContentTypeHeaderValue)
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
