package storage

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	//"reflect"
)

//var username string
//var password string

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:123456789@tcp(localhost:3306)/fuelproject")
	var err = db.Ping()
	if err != nil {
		log.Fatal("Database connection failed")
	}
}

func GetDb() *sql.DB {
	return db
}