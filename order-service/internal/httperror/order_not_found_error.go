package httperror

import (
	"errors"
	"net/http"

	"github.com/hailsayan/sophocles/order-service/internal/constant"
	"github.com/hailsayan/sophocles/pkg/httperror"
)

func NewOrderNotFoundError() *httperror.ResponseError {
	msg := constant.OrderNotFoundErrorMessage
	err := errors.New(msg)

	return httperror.NewResponseError(err, http.StatusNotFound, msg)
}
