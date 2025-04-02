package httperror

import (
	"errors"
	"net/http"

	"github.com/hailsayan/sophocles/auth-service/internal/constant"
	"github.com/hailsayan/sophocles/pkg/httperror"
)

func NewUserAlreadyExistError() *httperror.ResponseError {
	msg := constant.UserAlreadyExistErrorMessage
	err := errors.New(msg)

	return httperror.NewResponseError(err, http.StatusBadRequest, msg)
}
