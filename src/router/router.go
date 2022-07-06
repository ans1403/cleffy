package router

import (
	"cleffy/src/controller"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {

	router := gin.Default()

	ctrl := controller.NewAuthController()

	router.GET("/auth", ctrl.Auth)
	router.GET("/authcheck", ctrl.AuthCheck)
	router.GET("/token", ctrl.Token)
	router.GET("/certs", ctrl.Certs)
	router.GET("/userinfo", ctrl.UserInfo)

	return router
}
