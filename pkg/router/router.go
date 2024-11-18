package router

import (
	"github.com/gin-gonic/gin"
	"zvaljean/proxy-subscribe-api/pkg/config"
	"zvaljean/proxy-subscribe-api/pkg/controller"
)

func InitRouter(g *gin.Engine) {

	userController := controller.NewUserController(config.Db.Db)
	userController.SetupRouter(g)

}
