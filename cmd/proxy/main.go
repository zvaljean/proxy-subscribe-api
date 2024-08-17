package main

import (
	"fmt"
	"valjean/proxy/subscribe/pkg/config"
	"valjean/proxy/subscribe/pkg/functions"
	"valjean/proxy/subscribe/pkg/log"
)

func main() {

	config.InitCnf()
	config.InitLog()
	config.InitBiz()
	server := config.InitEngine()

	/**
	用户中心: 获取openId
	*/
	server.GET("/proxy/:user/:conf", functions.UserConf)

	err := server.Run(fmt.Sprintf(":%d", config.Cnf.Server.Port))
	log.FatalCheck(err, "proxy boot error!")
}
