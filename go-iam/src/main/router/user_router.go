package router

import "github.com/gin-gonic/gin"

func UserRouter(g *gin.RouterGroup) {
	g.Group("/user")
	{
		g.POST("/")
		g.POST("/role")
	}
}
