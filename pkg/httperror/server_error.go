package httperror

import (
	"errors"
	"net/http"

	"github.com/hailsayan/sophocles/pkg/constant"
)

func NewServerError() *ResponseError {
	msg := constant.InternalServerErrorMessage
	err := errors.New(msg)

	return NewResponseError(err, http.StatusInternalServerError, msg)
}
