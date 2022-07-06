package controller

import "github.com/gin-gonic/gin"

type AuthController interface {
	Auth(c *gin.Context)
	AuthCheck(c *gin.Context)
	Token(c *gin.Context)
	Certs(c *gin.Context)
	UserInfo(c *gin.Context)
}

func NewAuthController() AuthController {
	return &authController{}
}

type authController struct{}

func (ctrl *authController) Auth(c *gin.Context) {
	panic("NotImplementedError")
}

func (ctrl *authController) AuthCheck(c *gin.Context) {
	panic("NotImplementedError")
}

func (ctrl *authController) Token(c *gin.Context) {
	panic("NotImplementedError")
}

func (ctrl *authController) Certs(c *gin.Context) {
	panic("NotImplementedError")
}

func (ctrl *authController) UserInfo(c *gin.Context) {
	panic("NotImplementedError")
}
