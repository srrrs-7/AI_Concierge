package db

import (
	"auth/util/env"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDb(env *env.EnvParams[string]) (*sql.DB, error) {
	db, err := sql.Open(env.DB_DRIVER.Value, env.DB_ADDR.Value)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(60 * time.Second)

	return db, nil
}
