package db

import (
	"auth/util/env"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDb(env *env.EnvParams[string]) (*gorm.DB, error) {
	dsn := env.DB_ADDR.Value
	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db, err := gormDb.DB()
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(60 * time.Second)

	return gormDb, nil
}

func CloseDb(gormDb *gorm.DB) {
	db, err := gormDb.DB()
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}
