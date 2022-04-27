package router

import (
	controller "ragnarok/core/controllers"

	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/ping", controller.PingController)

	router.StaticFS("/static", http.Dir("static"))

	router.GET("/", controller.WelcomePage)

	return router
}
