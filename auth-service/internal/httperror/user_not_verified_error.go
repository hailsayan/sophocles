package httperror

import (
	"errors"
	"net/http"

	"github.com/hailsayan/sophocles/auth-service/internal/constant"
	"github.com/hailsayan/sophocles/pkg/httperror"
)

func NewUserNotVerifiedError() *httperror.ResponseError {
	msg := constant.UserNotVerifiedErrorMessage
	err := errors.New(msg)

	return httperror.NewResponseError(err, http.StatusUnauthorized, msg)
}
