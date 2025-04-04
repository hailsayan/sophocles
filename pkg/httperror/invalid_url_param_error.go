package httperror

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/hailsayan/sophocles/pkg/constant"
)

func NewInvalidURLParamError(param string) *ResponseError {
	msg := fmt.Sprintf(constant.InvalidURLParamErrorMessage, param)
	err := errors.New(msg)

	return NewResponseError(err, http.StatusBadRequest, msg)
}
