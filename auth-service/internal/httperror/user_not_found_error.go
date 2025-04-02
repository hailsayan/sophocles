package httperror

import (
	"errors"
	"net/http"

	"github.com/hailsayan/sophocles/auth-service/internal/constant"
	"github.com/hailsayan/sophocles/pkg/httperror"
)

func NewUserNotFoundError() *httperror.ResponseError {
	msg := constant.UserNotFoundErrorErrorMessage
	err := errors.New(msg)

	return httperror.NewResponseError(err, http.StatusBadRequest, msg)
}
