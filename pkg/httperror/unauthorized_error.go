package httperror

import (
	"errors"
	"net/http"

	"github.com/hailsayan/sophocles/pkg/constant"
)

func NewUnauthorizedError() *ResponseError {
	msg := constant.UnauthorizedErrorMessage
	err := errors.New(msg)

	return NewResponseError(err, http.StatusUnauthorized, msg)
}
