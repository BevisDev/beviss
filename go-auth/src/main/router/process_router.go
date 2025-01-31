package router

import "github.com/gin-gonic/gin"

func ProcessRouter(g *gin.RouterGroup) {
	p := g.Group("/process")
	{
		p.POST(":name")
	}
}
