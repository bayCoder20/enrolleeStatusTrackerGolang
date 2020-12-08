package main

import (
	"database/sql"

	_ "gopkg.in/go-sql-driver/mysql.v1"
)

var db *sql.DB

// GetDBconn to return DB connection
func GetDBconn() *sql.DB {

	db, err := sql.Open("mysql", "root:5filas@tcp(127.0.0.1:3306)/enrolleetrackergo")

	if err != nil {
		panic(err.Error())
	}
	return db
}
