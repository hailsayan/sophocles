package httperror

import (
	"errors"
	"net/http"

	"github.com/hailsayan/sophocles/auth-service/internal/constant"
	"github.com/hailsayan/sophocles/pkg/httperror"
)

func NewTokenAlreadyExistError() *httperror.ResponseError {
	msg := constant.TokenAlreadyExistErrorMessage
	err := errors.New(msg)

	return httperror.NewResponseError(err, http.StatusTooManyRequests, msg)
}
