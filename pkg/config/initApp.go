package config

import (
	"flag"
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
	"time"
	"valjean/proxy/subscribe/pkg/data"
	"valjean/proxy/subscribe/pkg/log"
)

var (
	Cnf  Config
	User *map[string]string
	Db   *SqliteDb
)

func InitCnf() {
	// 定义命令行参数
	var (
		configPath = flag.String("path", "configs", "config path, eg: --path=./configs")
	)

	// 解析命令行参数
	flag.Parse()

	//if len(*configPath) == 0 {
	//	*configPath = "configs"
	//}

	cnf, err := LoadConfig(*configPath)
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

func InitDb() {
	//var gormLogger logger.Interface
	//
	//if Cnf.Log.Debug {
	//	gormLogger = logger.Default
	//} else {
	//	gormLogger = logger.Discard
	//}
	logger := zapgorm2.New(zap.L())

	cnf := &gorm.Config{
		//Logger: gormLogger,
		Logger: logger,
	}

	db, err := gorm.Open(sqlite.Open(Cnf.Server.DbPath), cnf)
	log.FatalCheck(err, "failed to connect sqliteCnf")

	//Db = db

	Db = NewSqliteDb(db)
	Db.InitDb()
}

/**
初始化业务数据
*/

func InitBiz() {
	userPath := Cnf.Biz.UserPath
	if len(userPath) > 0 {
		// user,conf,data
		User = data.ParseCsvForMap(userPath, 2, 4)
	}

}
