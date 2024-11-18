package functions

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"zvaljean/proxy-subscribe-api/pkg/config"
	"zvaljean/proxy-subscribe-api/pkg/log"
	"zvaljean/proxy-subscribe-api/pkg/utils"
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

	flag := parseFlag(conf)

	log.Info("user: %s, value: %s, flag: %s", user, value, flag)

	/**
	*	0. base64
	*	1. base64 decode
	 */
	if flag == "1" {
		value = utils.Base64toStr(value)
	}

	ctx.String(http.StatusOK, value)

}

// 2410091
// 241009
func parseFlag(conf string) string {
	length := len(conf)
	if length < 7 {
		return "0"
	}
	flag := conf[len(conf)-1]
	return string(flag)
}
