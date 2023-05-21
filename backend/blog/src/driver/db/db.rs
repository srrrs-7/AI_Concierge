package db

import (
	"log"
	"template/util/env"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDb(env *env.Env) *gorm.DB {
	tx, err := gorm.Open(mysql.Open(env.DbAddr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	d, err := tx.DB()
	if err != nil {
		log.Fatal(err)
	}

	d.SetMaxOpenConns(50)
	d.SetConnMaxLifetime(60 * time.Second)

	return tx
}

func CloseDb(tx *gorm.DB) {
	d, err := tx.DB()
	if err != nil {
		log.Fatal(err)
	}

	defer d.Close()
}
