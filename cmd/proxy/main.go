package main

import (
	"fmt"
	"zvaljean/proxy-subscribe-api/pkg/config"
	"zvaljean/proxy-subscribe-api/pkg/log"
	"zvaljean/proxy-subscribe-api/pkg/router"
)

func main() {

	config.InitCnf()
	config.InitLog()
	//config.InitBiz()
	config.InitDb()
	engine := config.InitEngine()

	router.InitRouter(engine)

	err := engine.Run(fmt.Sprintf(":%d", config.Cnf.Server.Port))
	log.FatalCheck(err, "proxy boot error!")
}
