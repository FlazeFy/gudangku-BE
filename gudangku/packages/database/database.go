package database

import (
	"database/sql"
	"gudangku/configs"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	conf := configs.GetConfig()

	// username:password@protocol(address)/dbname
	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME
	db, err = sql.Open("mysql", connectionString)

	if err != nil {
		panic("connection error!!")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN Invalid!!")
	}
}

func CreateCon() *sql.DB {
	return db
}
