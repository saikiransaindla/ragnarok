package mysql

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlDB *gorm.DB

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

	mysqlDB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Fatal(err)
	}

	db, _ := mysqlDB.DB()
	pingError := db.Ping()

	if pingError != nil {
		log.Fatal(pingError)
	}

	fmt.Println("MySQL Connected")
}

func Ping() bool {
	db, _ := mysqlDB.DB()
	pingError := db.Ping()

	return pingError == nil
}
