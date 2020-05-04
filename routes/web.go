package routes

import (
	"github.com/gin-gonic/gin"
	"Gimic/controller"
)

func SetRoute()  *gin.Engine{
	router := gin.Default()

	router.GET("/user/all", controller.GetAllUser)
	router.POST("/user/create", controller.CreateUser)

	router.GET("/", controller.Halo)

	return router
}