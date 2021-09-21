package httperors

import (
	"net/http"
)

////////////errors ////////////////////////
type HttpError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
}

////////////success ////////////////////////
type HttpSuccess struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *HttpError {
	return &HttpError{
		Message: message,
		Code:    http.StatusBadRequest,
		Error:   "Ohh Well bad request buddy",
	}
}
func NewNotFoundError(message string) *HttpError {
	return &HttpError{
		Message: message,
		Code:    http.StatusBadRequest,
		Error:   "ooh boy the resource was Not Found!",
	}
}
func NewSuccessMessage(message string) *HttpSuccess {
	return &HttpSuccess{
		Message: message,
		Code:    http.StatusOK,
		Error:   "Delete success",
	}
}
func NewNoResultsMessage(message string) *HttpSuccess {
	return &HttpSuccess{
		Message: message,
		Code:    http.StatusOK,
		Error:   "No Results found",
	}
}
