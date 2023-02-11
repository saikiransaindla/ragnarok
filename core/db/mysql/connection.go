package mysql

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

func NewConnection() {
	// config := mysql.Config{
	// 	User:                 "root",
	// 	Passwd:               "password",
	// 	Net:                  "tcp",
	// 	Addr:                 "127.0.0.1:3306",
	// 	DBName:               "ragnarok",
	// 	AllowNativePasswords: true,
	// }

	var err error

	const dsn = "root:password@tcp(127.0.0.1:3306)/ragnarok"

	MysqlDB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Fatal(err)
	}

	db, _ := MysqlDB.DB()
	pingError := db.Ping()

	if pingError != nil {
		log.Fatal(pingError)
	}

	// // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	// db.SetMaxIdleConns(10)

	// // SetMaxOpenConns sets the maximum number of open connections to the database.
	// db.SetMaxOpenConns(100)

	// // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// db.SetConnMaxLifetime(time.Hour)

	fmt.Println("MySQL Connected")
}

func Ping() bool {
	db, _ := MysqlDB.DB()
	pingError := db.Ping()

	return pingError == nil
}
