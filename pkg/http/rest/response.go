package rest

import (
	"encoding/json"
	"net/http"
)

const (
	ContentTypeHeaderName  string = "Content-Type"
	ContentTypeHeaderValue string = "application/json; charset=utf-8"
)

type ErrorInfo struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

type ValidationErrorInfo struct {
	ErrorInfo
	Errors map[string][]string
}

func WriteResult(w http.ResponseWriter, data any) {
	WriteJsonResponse(w, http.StatusOK, data)
}

func WriteError(w http.ResponseWriter, statusCode int, message string) {
	WriteErrorInfo(w, &ErrorInfo{
		StatusCode: statusCode,
		Message:    message,
	})
}

func WriteValidationError(w http.ResponseWriter, message string, errors map[string][]string) {
	WriteValidationErrorInfo(w, &ValidationErrorInfo{
		ErrorInfo: ErrorInfo{
			StatusCode: http.StatusBadRequest,
			Message:    message,
		},
		Errors: errors,
	})
}

func WriteErrorInfo(w http.ResponseWriter, info *ErrorInfo) {
	WriteJsonResponse(w, info.StatusCode, info)
}

func WriteValidationErrorInfo(w http.ResponseWriter, info *ValidationErrorInfo) {
	WriteJsonResponse(w, http.StatusBadRequest, info)
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
