package httperror

import (
	"errors"
	"net/http"

	"github.com/hailsayan/sophocles/pkg/constant"
)

func NewRequestDuplicateError() *ResponseError {
	msg := constant.RequestDuplicateErrorMessage
	err := errors.New(msg)

	return NewResponseError(err, http.StatusConflict, msg)
}
