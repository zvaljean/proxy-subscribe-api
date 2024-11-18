package main

import (
	"fmt"
	"valjean/proxy/subscribe/pkg/config"
	"valjean/proxy/subscribe/pkg/log"
	"valjean/proxy/subscribe/pkg/router"
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
