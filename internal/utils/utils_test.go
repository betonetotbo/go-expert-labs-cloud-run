package utils

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_ConcatFields(t *testing.T) {
	assert.Equal(t, "1,2,3", ConcatFields("", "1", "2", "", "3"))
}

func Test_HttpError(t *testing.T) {
	e := NewHttpError(http.StatusBadGateway)
	assert.Equal(t, http.StatusBadGateway, GetHttpErrorCode(e))
}

func Test_HttpErrorUnknown(t *testing.T) {
	e := errors.New("bla")
	assert.Equal(t, http.StatusInternalServerError, GetHttpErrorCode(e))
}
