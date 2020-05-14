package shopware

import (
	"net/http"
)

// ErrorResponse model
type ErrorResponse struct {
	Errors *[]ErrorDetail `json:"errors"`
}

// ErrorDetails model
type ErrorDetails struct {
	Errors []*ErrorDetail `json:"errors,omitempty"`
}

// ErrorDetail model
type ErrorDetail struct {
	Code   string      `json:"code,omitempty"`
	Status string      `json:"status,omitempty"`
	Title  interface{} `json:"title,omitempty"`
	Detail string      `json:"detail,omitempty"`
}

func (e ErrorResponse) Error() string {
	return "TODO"
}

// AccessTokenInvalidError for 401 errors
type AccessTokenInvalidError struct {
	APIError
}

func (e AccessTokenInvalidError) Error() string {
	return "TODO"
}

// NotFoundError for 404 errors
type NotFoundError struct {
	APIError
}

func (e NotFoundError) Error() string {
	return "the requested resource can not be found"
}

// APIError model
type APIError struct {
	req *http.Request
	res *http.Response
	err *ErrorResponse
}

// RateLimitExceededError for rate limit errors
type RateLimitExceededError struct {
	APIError
}

func (e RateLimitExceededError) Error() string {
	return "TODO"
}

// BadRequestError error model for bad request responses
type BadRequestError struct{}

// InvalidQueryError error model for invalid query responses
type InvalidQueryError struct{}

// AccessDeniedError error model for access denied responses
type AccessDeniedError struct{}

// ServerError error model for server error responses
type ServerError struct{}
