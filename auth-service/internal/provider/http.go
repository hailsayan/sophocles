package provider

import (
	"github.com/gin-gonic/gin"
	"github.com/hailsayan/sophocles/auth-service/internal/config"
	"github.com/hailsayan/sophocles/auth-service/internal/controller"
)

func BootstrapHttp(cfg *config.Config, router *gin.Engine) {
	appController := controller.NewAppController()
	appController.Route(router)

	BootstrapUser(router)
}
