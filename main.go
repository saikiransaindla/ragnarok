package main

import (
	mysql "ragnarok/core/db/mysql"
	router "ragnarok/core/router"
)

//Go application entrypoint
func main() {

	mysql.NewConnection()
	router.NewRouter().Run(":8080")

}
