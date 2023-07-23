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

func WriteResult(w http.ResponseWriter, data interface{}) {
	w.Header().Set(ContentTypeHeaderName, ContentTypeHeaderValue)
	json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, statusCode int, message string) {
	info := &errorInfo{
		StatusCode: statusCode,
		Message:    message,
	}
	w.Header().Set(ContentTypeHeaderName, ContentTypeHeaderValue)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(info)
}

func WriteStatus(w http.ResponseWriter, statusCode int) {
	w.Header().Set(ContentTypeHeaderName, ContentTypeHeaderValue)
	w.WriteHeader(statusCode)
}
