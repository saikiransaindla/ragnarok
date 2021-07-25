package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var mysqlDB *sql.DB

func NewConnection() {
	config := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "ragnarok",
		AllowNativePasswords: true,
	}

	var err error

	mysqlDB, err = sql.Open("mysql", config.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingError := mysqlDB.Ping()

	if pingError != nil {
		log.Fatal(pingError)
	}

	fmt.Println("MySQL Connected")
}

func Ping() bool {
	pingError := mysqlDB.Ping()

	return pingError == nil
}
