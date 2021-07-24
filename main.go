package main

import (
	router "ragnarok/core/router"
)

//Go application entrypoint
func main() {

	router.NewRouter().Run(":8080")

}
