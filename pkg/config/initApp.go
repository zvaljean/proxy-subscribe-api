package config

import (
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"time"
	"valjean/proxy/subscribe/pkg/data"
	"valjean/proxy/subscribe/pkg/log"
)

var (
	Cnf  Config
	User *map[string]string
)

func InitCnf() {

	cnf, err := LoadConfig("configs")
	if err != nil {
		fmt.Printf(" error -> %v", err)
	}

	Cnf = cnf
}

func InitLog() {

	logPath := Cnf.Log.Path
	if len(logPath) == 0 {
		logPath = "app.log"
	}

	log.InitLog(logPath)
}

func InitEngine() *gin.Engine {

	gin.SetMode(gin.DebugMode)

	if Cnf.Log.Debug == false {
		gin.SetMode(gin.ReleaseMode)
	}

	server := gin.Default()

	server.ForwardedByClientIP = true
	_ = server.SetTrustedProxies([]string{"127.0.0.1"})

	server.Use(ginzap.Ginzap(log.ZapL(), time.RFC3339, true))
	server.Use(ginzap.RecoveryWithZap(log.ZapL(), true))

	return server
}

/**
初始化业务数据
*/

func InitBiz() {
	userPath := Cnf.Biz.UserPath
	if len(userPath) > 0 {
		User = data.ParseCsvForMap(userPath, 1, 2)
	}

}
