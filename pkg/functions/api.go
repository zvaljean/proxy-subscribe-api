package functions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"valjean/proxy/subscribe/pkg/log"
)

var wallet = map[string]int{}

// 用户获取proxy相关配置信息
func UserConf(ctx *gin.Context) {

	user := ctx.Param("user")
	conf := ctx.Param("conf")

	if len(user) == 0 || len(conf) == 0 {
		log.Error("user or conf param empty: user: %s, conf: %s", user, conf)
		ctx.String(http.StatusNotFound, "")
		return
	}

	log.Info(" user: %s, conf: %s", user, conf)

}
