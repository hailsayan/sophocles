package httperror

import (
	"errors"
	"net/http"

	"github.com/hailsayan/sophocles/pkg/httperror"
	"github.com/hailsayan/sophocles/product-service/internal/constant"
)

func NewProductNotFoundError() *httperror.ResponseError {
	msg := constant.ProductNotFoundErrorMessage
	err := errors.New(msg)

	return httperror.NewResponseError(err, http.StatusNotFound, msg)
}
