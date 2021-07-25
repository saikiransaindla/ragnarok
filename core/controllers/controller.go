package controller

import (
	"time"

	mysql "ragnarok/core/db/mysql"

	"github.com/gin-gonic/gin"
)

//Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
	Name string
	Time string
}

func PingController(res *gin.Context) {
	mysqlConnected := mysql.Ping()
	res.JSON(200, gin.H{
		"message":        "PONG",
		"mysqlConnected": mysqlConnected,
	})
}

func WelcomePage(res *gin.Context) {
	welcome := Welcome{"Sai Kiran", time.Now().Format(time.Stamp)}

	res.HTML(200, "welcome-template.html", welcome)
}
