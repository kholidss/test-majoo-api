package utils

import (
	"math"
	"net/http"
	"sort"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ErrorResponse is the response that represents an error.
type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// Error is required by the error interface.
func (e ErrorResponse) Error() string {
	return e.Message
}

// StatusCode is required by routing.HTTPError interface.
func (e ErrorResponse) StatusCode() int {
	return e.Status
}

// InternalServerError creates a new error response representing an internal server error (HTTP 500)
func InternalServerError(msg string) *ErrorResponse {
	if msg == "" {
		msg = "We encountered an error while processing your request."
	}
	return &ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: msg,
	}
}

// NotFound creates a new error response representing a resource-not-found error (HTTP 404)
func NotFound(msg string) *ErrorResponse {
	if msg == "" {
		msg = "The requested resource was not found."
	}
	return &ErrorResponse{
		Status:  http.StatusNotFound,
		Message: msg,
	}
}

// Unauthorized creates a new error response representing an authentication/authorization failure (HTTP 401)
func Unauthorized(msg string) *ErrorResponse {
	if msg == "" {
		msg = "You are not authenticated to perform the requested action."
	}
	return &ErrorResponse{
		Status:  http.StatusUnauthorized,
		Message: msg,
	}
}

// Forbidden creates a new error response representing an authorization failure (HTTP 403)
func Forbidden(msg string) *ErrorResponse {
	if msg == "" {
		msg = "You are not authorized to perform the requested action."
	}
	return &ErrorResponse{
		Status:  http.StatusForbidden,
		Message: msg,
	}
}

// BadRequest creates a new error response representing a bad request (HTTP 400)
func BadRequest(msg string) *ErrorResponse {
	if msg == "" {
		msg = "Your request is in a bad format."
	}
	return &ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: msg,
	}
}

func UnprocessableEntity(msg string) *ErrorResponse {
	if msg == "" {
		msg = "Error in transaction operation"
	}
	return &ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: msg,
	}
}

type invalidField struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// InvalidInput creates a new error response representing a data validation error (HTTP 400).
func InvalidInput(errs validation.Errors) *ErrorResponse {
	var details []invalidField
	var fields []string
	for field := range errs {
		fields = append(fields, field)
	}
	sort.Strings(fields)
	for _, field := range fields {
		details = append(details, invalidField{
			Field: field,
			Error: errs[field].Error(),
		})
	}

	return &ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: "There is some problem with the data you submitted.",
		Details: details,
	}
}

type SuccessResponse struct {
	Message string
}

func SuccessMessage(msg string) *SuccessResponse {
	if msg == "" {
		msg = "Successful transaction."
	}

	return &SuccessResponse{
		Message: msg,
	}
}

type PaginateMeta struct {
	CurrenctPage uint `json:"currenct_page"`
	PerPage      uint `json:"per_page"`
	TotalData    uint `json:"total_data"`
	TotalPage    uint `json:"toal_page"`
}
type PaginateResponse struct {
	Meta *PaginateMeta `json:"meta"`
	Data interface{}   `json:"data"`
}

func Paginate(data interface{}, total uint, currentPage uint, limit uint) *PaginateResponse {
	totalPages := math.Ceil(float64(total) / float64(limit))

	meta := &PaginateMeta{
		CurrenctPage: currentPage,
		PerPage:      limit,
		TotalData:    total,
		TotalPage:    uint(totalPages),
	}

	return &PaginateResponse{
		Meta: meta,
		Data: data,
	}
}
