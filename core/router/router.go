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

	router.POST("/v1/user", controller.CreateUser)

	router.PATCH("v1/user", controller.UpdateUser)

	router.GET("v1/user", controller.FetchUser)

	router.GET("v1/users", controller.FetchUsers)

	return router
}
