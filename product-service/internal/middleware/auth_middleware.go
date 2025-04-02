package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hailsayan/sophocles/pkg/httperror"
	"github.com/hailsayan/sophocles/pkg/utils/ginutils"
)

func AuthMiddleware(ctx *gin.Context) {
	if _, ok := ginutils.GetXUserID(ctx); !ok {
		ctx.Error(httperror.NewUnauthorizedError())
		ctx.Abort()
		return
	}
}
