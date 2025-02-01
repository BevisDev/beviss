package router

import (
	"github.com/gin-gonic/gin"
	"goiam/src/main/di"
)

func AuthRouter(g *gin.RouterGroup) {
	c := di.NewAuthDI()
	g.GET("/signin", c.SignIn)
	g.GET("/signup", c.SignUp)
}
