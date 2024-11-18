package config

import (
	"gorm.io/gorm"
	"time"
	"zvaljean/proxy-subscribe-api/pkg/entity"
	"zvaljean/proxy-subscribe-api/pkg/log"
)

type SqliteDb struct {
	Db *gorm.DB
}

func NewSqliteDb(db *gorm.DB) *SqliteDb {

	// var gormLogger logger.Interface
	//
	// if config.SqliteCnf.Debug {
	// 	gormLogger = logger.Default
	// } else {
	// 	gormLogger = logger.Discard
	// }
	//
	// cnf := &gorm.Config{
	// 	Logger: gormLogger,
	// }
	//
	// db, err := gorm.Open(sqlite.Open(config.SqliteCnf.Path), cnf)
	// if err != nil {
	// 	panic("failed to connect sqliteCnf")
	// }

	return &SqliteDb{Db: db}
}

func (data *SqliteDb) SetUpPool() {
	db, err := data.Db.DB()
	if err != nil {
		return
	}

	db.SetConnMaxIdleTime(time.Hour)
	db.SetConnMaxLifetime(24 * time.Hour)
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(200)

	session := &gorm.Session{PrepareStmt: true}
	data.Db.Session(session)

}

func (data *SqliteDb) Close() error {
	return nil
}

func (data *SqliteDb) InitDb() {
	err := data.Db.AutoMigrate(&entity.User{})
	log.FatalCheck(err, "Db.AutoMigrate")
	//err = data.Db.AutoMigrate(&entity.User{})
	//log.FatalCheck(err, "Db.AutoMigrate")
}
