package utils

import (
	"fmt"
	"net/http"
)

type (
	httpError struct {
		code int
	}
)

func (e *httpError) Error() string {
	return fmt.Sprintf("HTTP %d", e.code)
}

func NewHttpError(code int) error {
	return &httpError{code}
}

func GetHttpErrorCode(err error) int {
	e, ok := err.(*httpError)
	if !ok {
		return http.StatusInternalServerError
	}
	return e.code
}

func ConcatFields(fields ...string) string {
	r := ""
	for _, field := range fields {
		if len(field) > 0 {
			if len(r) > 0 {
				r = fmt.Sprintf("%s,%s", r, field)
			} else {
				r = field
			}
		}
	}
	return r
}
