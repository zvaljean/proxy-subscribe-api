package functions

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"valjean/proxy/subscribe/pkg/config"
	"valjean/proxy/subscribe/pkg/log"
)

// 用户获取proxy相关配置信息
func UserConf(ctx *gin.Context) {

	user := ctx.Param("user")
	conf := ctx.Param("conf")

	if len(user) == 0 || len(conf) == 0 {
		log.Error("user or conf param empty: user: %s, conf: %s", user, conf)
		ctx.String(http.StatusNotFound, "")
		return
	}

	// 根据用户名获取配置
	// user-conf,data
	key := fmt.Sprintf("%s-%s", user, conf)
	//key := user
	value, exist := (*config.User)[key]
	if !exist {
		log.Info("user: %s, not exist", user)
		ctx.String(http.StatusNotFound, "")
		return
	}

	log.Info("user: %s, value: %s", user, value)

	ctx.String(http.StatusOK, value)

}
