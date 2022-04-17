package api

import (
	"net/http"
)

type ApiResponse interface{}

type ApiResponseWithErr struct {
	Err string `json:"error,omitempty"`
}

type ApiResponseWithErrors struct {
	Errors []string `json:"errors,omitempty"`
}

func UserErr(err error) (int, ApiResponseWithErr) {
	return http.StatusBadRequest, ApiResponseWithErr{err.Error()}
}

func UserErrs(errs []error) (int, ApiResponseWithErrors) {
	var strs []string
	for _, e := range errs {
		strs = append(strs, e.Error())
	}
	return http.StatusBadRequest, ApiResponseWithErrors{strs}
}

func ServerErr(err error) (int, ApiResponseWithErr) {
	return http.StatusInternalServerError, ApiResponseWithErr{err.Error()}
}

func ServerErrs(errs []error) (int, ApiResponseWithErrors) {
	var strs []string
	for _, e := range errs {
		strs = append(strs, e.Error())
	}
	return http.StatusInternalServerError, ApiResponseWithErrors{strs}
}
