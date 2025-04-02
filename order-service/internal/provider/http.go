package provider

import (
	"github.com/gin-gonic/gin"
	"github.com/hailsayan/sophocles/order-service/internal/config"
	"github.com/hailsayan/sophocles/order-service/internal/controller"
)

func BootstrapHttp(cfg *config.Config, router *gin.Engine) {
	appController := controller.NewAppController()
	appController.Route(router)

	BootstrapOrder(cfg, router)
}
