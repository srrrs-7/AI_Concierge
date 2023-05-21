package db

import (
	"database/sql"
	"template/util/env"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDb(env *env.Env) (*sql.DB, error) {
	db, err := sql.Open(env.DB_DRIVER, env.DB_ADDR)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(60 * time.Second)

	return db, nil
}
