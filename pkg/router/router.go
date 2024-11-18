package router

import (
	"github.com/gin-gonic/gin"
	"valjean/proxy/subscribe/pkg/config"
	"valjean/proxy/subscribe/pkg/controller"
)

func InitRouter(g *gin.Engine) {

	userController := controller.NewUserController(config.Db.Db)
	userController.SetupRouter(g)

}
