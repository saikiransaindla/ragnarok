package controller

import (
	mysql "ragnarok/core/db/mysql"
	"ragnarok/core/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	user := models.User{Id:"abcd1", Name:"Sai", Email:"saikiran@gmail.com"}

	result := mysql.MysqlDB.Create(&user)

	ctx.JSON(200, gin.H{
		"message":        "PONG",
		"result": result,
		"param": ctx.Param("name"),
	})
}

func UpdateUser(ctx *gin.Context) {
	mysqlConnected := mysql.Ping()
	ctx.JSON(200, gin.H{
		"message":        "PONG",
		"mysqlConnected": mysqlConnected,
	})
}

func FetchUser(ctx *gin.Context) {
	user := models.User{Id:"abcd4", Name:"Sai", Email:"saikiran@gmail.com"}

	mysql.MysqlDB.Create(&user)

	ctx.JSON(200, gin.H{
		"message":"PONG",
		"result": user,
		"param": ctx.Param("name"),
	})
}

func FetchUsers(ctx *gin.Context) {
	mysqlConnected := mysql.Ping()
	ctx.JSON(200, gin.H{
		"message":        "PONG",
		"mysqlConnected": mysqlConnected,
	})
}