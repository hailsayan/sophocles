package httperror

import (
	"errors"
	"net/http"

	"github.com/hailsayan/sophocles/order-service/internal/constant"
	"github.com/hailsayan/sophocles/pkg/httperror"
)

func NewOrderAlreadyPaidError() *httperror.ResponseError {
	msg := constant.OrderAlreadyPaidErrorMessage
	err := errors.New(msg)

	return httperror.NewResponseError(err, http.StatusConflict, msg)
}
